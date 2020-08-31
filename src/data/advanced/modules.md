---
title: "Modules"
tags: advanced
index: 6
---
## Modules
The use of modules is the new dependency management system added in version 1.11+, they work as a collection of packages that have a go.mod file in the root directory, this file contain all the path of packages, also contains all the external package requirements that need to be download, Go tool will check the local code and import all of this packages and add the latest version to the go.mod file if it have been not added.

The way to initialise a module is with he following command:
``` go
go mod init
```
