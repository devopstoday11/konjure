/*
Copyright 2019 GramLabs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package berglas

import (
	"github.com/carbonrelay/konjure/plugin/berglas/generator"
	"github.com/carbonrelay/konjure/plugin/berglas/transformer"
	"github.com/spf13/cobra"
)

var (
	// NewBerglasGeneratorExecPlugin creates a new command for running Berglas as an executable plugin
	NewBerglasGeneratorExecPlugin = generator.NewBerglasGeneratorExecPlugin
	// NewBerglasTransformerExecPlugin creates a new command for running Berglas as an executable plugin
	NewBerglasTransformerExecPlugin = transformer.NewBerglasTransformerExecPlugin
)

// NewBerglasCommand creates a new command for running Berglas from the CLI
func NewBerglasCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "berglas",
		Short:      "Generate secrets using Berglas",
		Deprecated: "Berglas support will be removed in the next release",
	}
	cmd.AddCommand(generator.NewBerglasGeneratorCommand())
	cmd.AddCommand(transformer.NewBerglasTransformerCommand())
	return cmd
}
