---
title: "Variables"
tags: basic
---
## Variables

Golang support the following built-in types:

For strings:
  - rune  
  - string  

For Boolean values:  
  - bool

For numeric types:
  - int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr
  - byte (alias of uint8)
  - rune (alias of int32)
  - float32, float64.
  - complex64, complex128

A variable can be declared in the following ways:
```go
var name string
var name1, name2 string = "value", "value2"
```
A short way to declare a variable is using the **:=** operator, that is used to declare a variable with an implicit type based on the right side of the operation.
```go
name := "1"
name1, name2, name3 := "value", "value2", "value3"
i := 1
i,j,k := 0, 1, 2

```

### Strings and Runes

Strings by default are ***UTF8***, with this using a char is not possible, since the size of a char is 1 byte, and a UTF8 symbol can use up to 4 bytes (32 bits), to solve this problem runes are used to avoid having problems with unicode sets.  
Example  
```go
package main
import (
    "fmt"
)

func printByChar(s string) {
    chars := []byte(s)
    fmt.Printf("This slice has a len of %d\n", len(chars))
    for i, c := range chars {
        fmt.Printf("Char %d value is %d\n", i, c)
    }
}

func printByRune(s string) {
    runes := []rune(s)
    fmt.Printf("This slice has a len of %d\n", len(runes))
    for i, r := range runes {
        fmt.Printf("Rune %d value is %c\n", i, r)
    }
}

func main() {
    hello := "Hello world 👋🌎"
    hello2 := "Hello workd!!!"
    fmt.Printf("String 1 has a len of %d\n", len(hello))
    fmt.Printf("String 2 has a len of %d\n", len(hello2))

    printByChar(hello)
    printByChar(hello2)
    printByRune(hello)
    printByRune(hello2)

}
```
Output:
```plain
String 1 has a len of 20
String 2 has a len of 14
This slice has a len of 20
Char 0 value is 72
Char 1 value is 101
Char 2 value is 108
Char 3 value is 108
Char 4 value is 111
Char 5 value is 32
Char 6 value is 119
Char 7 value is 111
Char 8 value is 114
Char 9 value is 108
Char 10 value is 100
Char 11 value is 32
Char 12 value is 240
Char 13 value is 159
Char 14 value is 145
Char 15 value is 139
Char 16 value is 240
Char 17 value is 159
Char 18 value is 140
Char 19 value is 142
This slice has a len of 14
Char 0 value is 72
Char 1 value is 101
Char 2 value is 108
Char 3 value is 108
Char 4 value is 111
Char 5 value is 32
Char 6 value is 119
Char 7 value is 111
Char 8 value is 114
Char 9 value is 107
Char 10 value is 100
Char 11 value is 33
Char 12 value is 33
Char 13 value is 33
This slice has a len of 14
Rune 0 value is H
Rune 1 value is e
Rune 2 value is l
Rune 3 value is l
Rune 4 value is o
Rune 5 value is  
Rune 6 value is w
Rune 7 value is o
Rune 8 value is r
Rune 9 value is l
Rune 10 value is d
Rune 11 value is  
Rune 12 value is 👋
Rune 13 value is 🌎
This slice has a len of 14
Rune 0 value is H
Rune 1 value is e
Rune 2 value is l
Rune 3 value is l
Rune 4 value is o
Rune 5 value is  
Rune 6 value is w
Rune 7 value is o
Rune 8 value is r
Rune 9 value is k
Rune 10 value is d
Rune 11 value is !
Rune 12 value is !
Rune 13 value is !
```
In this example we can see that the two string has the same number of characters both have different length, since one has unicode characters, but when a rune is used to read it, both have the same size.

### Alias
Alias are a way of renaming a type of variable, for example:
A byte and rune are alias and int32 respective, this mean that they are not different types of data but they have an alternative spelling.

This is something similar of int and char in C, you can do operation with them since they have the same size, but is not the same since Go will not let you operation between then because they have the size, but since they are the same type but with different name.

### Constants
Constants works like another language and have to follow this rules
They have to use they keyword cons  to be declared
Can only be can be ***character, string, boolean, or numeric values***.
Cannot be declared with the := shortcut

```go
const name = "value"
```


### Zero values
- **0** for numeric types
- **false** for the boolean
- **""** (empty string) for strings
- **nil** for pointer types
