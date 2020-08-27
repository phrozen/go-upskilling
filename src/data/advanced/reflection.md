---
title: "Reflection"
tags: advanced
---
## Reflection

The package reflect can be imported to work with reflections, it allow to know the type, the name and set values.
Since go is strongly typed reflection in go have 3 rules:

#### Reflection goes from interface value to reflection object.
The reflection package can get information from a interface such as type, name and value of a variable. This can be done with methods on the package, for example
```go
func TypeOf(i interface{}) Type
```
This functions receive a empty interface, since this can hold any value, so when a argument is passed in the function first is stored in a empty interface.
#### Reflection goes from reflection object to interface value.
Reflection goes both ways and it can be transform an interface to a value, this can be done with the following method
```go
func (v Value) Interface() interface{}
```
#### To modify a reflection object, the value must be settable.

Settability is the ability to modify the original values of the object, this is related to the addressability of the object and that if we have a copy of the value or the value himself.

One can berify the settability of an object with the method `v.CanSet()`

NOTE: All methods in reflection package assumes that the programmer knows that is doing and many functions and methods will panics when used incorrect.

More information: https://blog.golang.org/laws-of-reflection
