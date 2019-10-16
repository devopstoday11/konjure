# Konjure Examples

## Example Usage

* [Helm Generator](helm.md)
* [Jsonnet Generator](jsonnet.md)

## Kustomize vs. CLI

These examples all use Konjure as a Kustomize executable plugin. While this is preferred way to use Konjure, all of the generators and transformers can also be used without Kustomize through direct invocation of the `konjure` binary.

For example, the Helm Generator example can be run directly from the CLI using:

```sh
konjure helm --name elasticsearch --version 1.31.1 --set data.replicas=3 stable/elasticsearch
```