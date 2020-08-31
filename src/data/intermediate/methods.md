---
title: "Methods"
tags: intermediate
index: 5
---
## Methods

For an object in go to implement an interface it needs to suffice all the methods declared in the body of the interface, for an object to do this it just need to declare a function pointing to the object.

Example:

``` go
//Using as reference the method declared in the Writer interface
type console struct{
   ...
}

func (m *console) Write (n int, err error){
    //body....
    return 0, nil

}
//Now the struct console is implementing the interface Writer since it suffice the interface implementing all methods declared
```
