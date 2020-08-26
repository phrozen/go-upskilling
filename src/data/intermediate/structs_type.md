---
  title: "Structs and type"
  tags: intermeidate
---
## Type

The type keyword is used to create a new type, it can be used to create structs, interfaces or new alias for a existing type
 ```go
type name struct{...}
type name interface{...}
type myNumber int32
 ```
## Structs

Structs are a collection of fields with declared types. Structs can be initializaded by a JSON like syntax.
 ```go
package main

import "fmt"

type person struct{
      name string
      age int
      address string
}

func main() {
      p := person{
            name :"my name",
            age : 28,
            address : "my address"
      }

      fmt.Printf("name: %s\nage: %d\naddress %s\n",p.name,p.age,p.address)
      p.name = "change of name"
      p.address = "change of address"
      fmt.Printf("name: %s\nage: %d\naddress %s\n",p.name,p.age,p.address)
}
```
Output:
```plain
name: my name
age: 28
address my address
name: change of name
age: 28
address change of address
 ```
