---
title: "Getting Started"
tags: basic
---

## Getting Started

To download the go compiler just follow the following [instructions](https://golang.org/doc/)

### Environment

The go compiler is expecting a structure for the imports, to do this it relays in the `$GOPATH` env variable, this variable will tell the go compiler were to found the modules to import.
The `$GOPATH` will contain a directory with 3 folders: `src`, `pkg` and `bin`.

#### bin
The bin directory contains the binary files installed by the command
`go install`
#### pkg
The pkg directory contains the precompiled packages for future project compiles.
#### src
The src directory contains the source code of the projects. Since is usual to integrate github  in the projects, it is common that the structure of the src directory contains a github.com/user/project.
