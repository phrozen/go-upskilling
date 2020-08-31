---
title: Writing code
tags: basic
index: 3
---

## Writing Code

First let's check that the installation is working, for this we will create a new folder inside the src directory in the $GOPATH

``` bash
cd $GOPATH
mkdir example
```

inside the directory we will create the file hello.go with this content:

``` go
package main func main()
{
    fmt.Printf("hello, world\n")
}
```

Run it like a script!

``` bash
$ go run hello.go
```

or compile it!

``` bash
$ go build
```

**_Note_**: go build utilises the working directory as a target for compilation, you can also pass the path as a param.

And run it

``` bash
$ ./hello
```

### Packages

A package is a library in Go, If a package name is **main** go will search for the code for the main function and compile the project as an executable; otherwise the package is compile as a library

A project can contain multiple packages, each will be represented as a folder in the src directory, packages can contain more than a file, the scope of the package are all files in the folder.


**Importing** and **exporting** Go has 2 access settings for packages, **exported** and **unexported**. All names of methods, functions, variables, constants and interfaces that start with a capital letter are exported and all that not are unexported.

**Unexported**( something like private) access mean that it can be access from other files in the same package, but cannot be access outside the package scope.

**Exported**(Something like public) access mean that it can be access from other packages when the package is imported. All exported variables, interfaces, functions, methods and constants have to be commented.

**Importing** packages is done with $GOPATH + the directory of your local package or the external package

```go
import (
    "fmt"
    "github.com/user/hello/example"
    "github.com/google/go-cmp/cmp"
)
```

During the import of the package one can overwrite the name of the package.

```go
import sql "mysql"
```

In the example adobe we are importing the mysql package and renaming it "sql" if in a moment in time me want to change from mysql to sqlserver or another sql driver, we just need to change the imported packaged and all the code can remain the same.

Another form of importing packages using the underscore symbol, this will import the package only for his side-efects (inititalition funcitons)

More info

<https://golang.org/doc/code.html>

<https://golang.org/doc/effective_go.html#package-names>

More about go tool can be found with the go help command
