# Provider Harbor

`provider-harbor` is a [Crossplane](https://crossplane.io/) provider
harbor that is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the Harbor
API.

Based on this Terraform provider:  
https://registry.terraform.io/providers/goharbor/harbor/latest/docs

Partly based on code by GlobalLogic UK&I:  
https://github.com/thisisnttheway/provider-harbor

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/thisisnttheway/provider-harbor/issues).
