---
title: "Pointers"
tags: intermediate
index: 1
---

## Pointers
A pointer is a special variable used to store a memory value, this mean that instead of the variable containing a value, it contains the value of the direction on memory of the variable.

This allows having some of the most important functions, allowing the variables be passed as reference in the functions instead of been passed as values.

The pointer has 2 reserves operators `*` and `&`, `&` can be used before the variable name and it will return the variable memory location. * can be used before the type of store value in the declaring and it will mean that is a pointer, and also is used to deference pointer values. When differencing a pointer, we will access to the value stored instead of the memory location of the value.

Something that we always keep in mind is that in Go does not support pointer arithmetic like C/C++.
``` go
package main
import "fmt"

func passAsValue(variable int){
    fmt.Println("Memory dir inside value func ", &variable)
    variable = 10
    fmt.Println("Value inside value func ", variable)
}



func passAsReference(variable *int){
    fmt.Println("Memory dir insde reference func ",variable)
    *variable = 20
    fmt.Println("Value inside reference func ", *variable)
}

func main() {
    variable := 1;
    fmt.Println("Memory dir inside main func " , &variable)
    passAsValue(variable)
    fmt.Println("Value after pass as a value " ,variable)
    passAsReference(&variable)
    fmt.Println("Value after pass as a reference", variable)

}
```

Output:

``` plain
Memory dir inside main func  0xc000018080
Memory dir inside value func  0xc000018088
Value inside value func  10
Value after pass as a value  1
Memory dir inside reference func  0xc000018080
Value inside reference func  20
Value after pass as a reference 20
```
