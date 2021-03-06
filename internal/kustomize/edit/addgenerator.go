// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package edit

import (
	"errors"
	"log"

	"github.com/carbonrelay/konjure/internal/kustomize/edit/kustinternal"
	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/api/filesys"
)

type addGeneratorOptions struct {
	generatorFilePaths []string
}

// newCmdAddGenerator adds the name of a file containing a generator to the kustomization file.
func newCmdAddGenerator(fSys filesys.FileSystem) *cobra.Command {
	var o addGeneratorOptions

	cmd := &cobra.Command{
		Use:   "generator",
		Short: "Add the name of a file containing a generator to the kustomization file.",
		Example: `
		add generator {filepath}`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := o.Validate(args)
			if err != nil {
				return err
			}
			err = o.Complete(cmd, args)
			if err != nil {
				return err
			}
			return o.RunAddGenerator(fSys)
		},
	}
	return cmd
}

// Validate validates addGenerator command.
func (o *addGeneratorOptions) Validate(args []string) error {
	if len(args) == 0 {
		return errors.New("must specify a generator file")
	}
	o.generatorFilePaths = args
	return nil
}

// Complete completes addGenerator command.
func (o *addGeneratorOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

// RunAddGenerator runs addGenerator command (do real work).
func (o *addGeneratorOptions) RunAddGenerator(fSys filesys.FileSystem) error {
	generators, err := kustinternal.GlobPatterns(fSys, o.generatorFilePaths)
	if err != nil {
		return err
	}
	if len(generators) == 0 {
		return nil
	}

	mf, err := kustinternal.NewKustomizationFile(fSys)
	if err != nil {
		return err
	}

	m, err := mf.Read()
	if err != nil {
		return err
	}

	for _, generator := range generators {
		if kustinternal.StringInSlice(generator, m.Generators) {
			log.Printf("generator %s is already in kustomization file", generator)
			continue
		}
		m.Generators = append(m.Generators, generator)
	}

	return mf.Write(m)
}
