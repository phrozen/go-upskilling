---
title: "Command Line Interfaces"
tags: next
index: 2
---
## Command Line Interfaces

### Flag package

This package has a lot utility methods that allows an easy integration with command line arguments. The following example will take two flags, one for a number and another for a string. The program will create a open the file with the name and will write numbers from 0 to the number.

 ``` go
package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
)

var nFlag = flag.Int("n", 0, "display message for n")
var fFlag = flag.String("f", "", "display message for f")

func main() {
    flag.Parse()
    if *nFlag == 0 {
        fmt.Println("Value of n cannot be 0")
        return
    }

    if *fFlag == "" {
        fmt.Println("Value of f cannot be empty")
        return
    }

    fd, err := os.OpenFile(*fFlag, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer fd.Close()
    data := ""
    for i := 0; i <= *nFlag; i++ {
        data += strconv.Itoa(i) + " "
    }

    fd.Write([]byte(data))
}
```

The code is run like this:

``` bash
$ go build hello.go
$ ./hello -n 10 -f something
```

And the file contains:
``` text
$ cat something
0 1 2 3 4 5 6 7 8 9 10
```

### Viper & Cobra

Cobra and Viper are third party repos that allow an easy configuration for applications.
Viper allows reading configuration files from different formats and translates and allows to set values, on the other hand Cobra is a powerful tool that helps with CLI applicartions. Some examples that use those tools are Hugo, kubernates and github CLI.


### Use case

Vegeta is one CLI tool that is used perform testing on http servers, is totally created in go.

https://github.com/tsenart/vegeta
https://golang.org/pkg/flag/
https://github.com/spf13/viper
https://github.com/spf13/cobra
