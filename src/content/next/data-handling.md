---
title: "Data handling"
tags: next
index: 1
---
## Data handling
### Files

The package os  include most of the operation for file management. The package offers the file struct that allow to have a pointer to a file or a directory, this struct allow to open an close files, change owner, read or write to the file.
It is highly recommended to add a defer with the close statement to make sure that the file is always close at the end.
os.File offers two open methods `File.Open(string)` and `File.OpenFile(string,int,FileMode)`,  File.Open(String) will open the file in read only mode.
The File structs have implemented already a Read and a Write method, those methods are declared in the package io as part of the interfaces `io.Writer` and `io.Reader`.
This interface allows also to read from console or other media.

 ```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fd, err := os.Open("/Users/csantana/example")

    if err != nil {
        fmt.Println(err)
        return
    }

    defer fd.Close()
    names, err := fd.Readdirnames(5)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Contents of directory " + fd.Name())

    for _, name := range names {
        childFd, err := os.Open(fd.Name() + "/" + name)
        buffer := make([]byte, 128)

        if err != nil {
            fmt.Println(err)
            return
        }

        defer childFd.Close()
        childFd.Read(buffer)
        fmt.Println("Text in file with name " + name)
        fmt.Println(string(buffer))
    }

}
```

This example will read up to 5 files in the directory and then will print the contents of the files

### JSON

The package "encoding/JSON" has a lot of methods that allow an easy conversion from JSON-struct and struc-JSON, this can be done with the Marshal and UnMarshal methods.
This important to mention that those methods respect the exported and non-exported variables, allowing to hide values when something is marshall to JSON

 ``` go
package main

import (
    "encoding/json"
    "fmt"
)

type personData struct {
    Name       string
    LastName   string
    Age        int
    Email      string
    creditCard string
}

func createPerson(name string, lastName string, age int, email string, cc string) *personData {
    return &personData{name, lastName, age, email, cc}
}

func main() {
    person := createPerson("Carlos", "Santana", 27, "csantana@luxoft.com", "XXXX XXXX XXXX XXXX")
    fmt.Println(person)
    b, err := json.Marshal(person)
    if err != nil {
        return
    }
    fmt.Println(string(b))
}
 ```

Output:

``` plain
&{Carlos Santana 27 csantana@luxoft.com XXXX XXXX XXXX XXXX}
{"Name":"Carlos","LastName":"Santana","Age":27,"Email":"csantana@luxoft.com"}
```
In this example we can see that struct person data was converted to JSON, in the nex example I will do the inverse, will transform the json readed from a file to the struct.

 ``` go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type personData struct {
    Name       string
    LastName   string
    Age        int
    Email      string
    creditCard string
}

func main() {
    var person personData
    fd, err := os.Open("jsonData")
    if err != nil {
        return
    }
    defer fd.Close()
    b := make([]byte, 114)
    fd.Read(b)
    fmt.Println("data on b :")
    fmt.Println(b)
    fmt.Println("String readeable of b")
    fmt.Println(string(b))
    err = json.Unmarshal(b, &person)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(person)
}
```

output:

``` plain
data on b :
[10 123 34 78 97 109 101 34 58 34 67 97 114 108 111 115 34 44 34 76 97 115 116 78 97 109 101 34 58 34 83 97 110 116 97 110 97 34 44 34 65 103 101 34 58 50 55 44 3
4 69 109 97 105 108 34 58 34 99 115 97 110 116 97 110 97 64 108 117 120 111 102 116 46 99 111 109 34 44 34 99 114 101 100 105 116 67 97 114 100 34 58 34 88 88 88
88 32 88 88 88 88 32 88 88 88 88 32 88 88 88 88 34 125 10]
String readeable of b

{"Name":"Carlos","LastName":"Santana","Age":27,"Email":"csantana@luxoft.com","creditCard":"XXXX XXXX XXXX XXXX"}

{Carlos Santana 27 csantana@luxoft.com }
```

Tags can be added to the code so the name of the variables is not the same show in the JSON, those tags can be added in the declaration of the structs

``` go
type personData struct {
    Name       string `json:"name"`
    LastName   string `json:"lastName"`
    Age        int    `json:"years"`
    Email      string `json:"email"`
    creditCard string
}
```
https://blog.golang.org/json
https://golang.org/pkg/os/
https://golang.org/pkg/io/
