---
title: "Anonymous Functions and Clousures"
tags: "advanced"
index: 5
---
## Anonymous functions and clousures

### Anonymous Functions
An anonymous functions is a function that does not have a name, they are created dynamically, like the variables, they can be described as follow:

 ``` go
dynamic := func() {

}
```

This allow to have flexibility at the momment of writting code

 ``` go
import "fmt"


func printHello() {
    fmt.println("Hello")
}

func printBye() {
    fmt.println("Bye")
}

func main() {
    //We set the value of myVar to an anonymous function
    myVar := func() {
        println("stuff")
    }
    myVar()

    //We assing the funciton to the variable, changing the anonymous function
    // for printBye
    myVar = printBye
    myVar()

    myVar = printHello
    myVar()

    //We can overwrite the non anonymous function for an anonymous one
    myVar = func() {
        fmt.println("something")
    }
    myVar()
}
```

In the example abode we have a variable that is pointing to a anonymous function that can be executed has a function, but the content of the function is changed with each assign, it even can be assign to non-anonymous functions

### Clousures
Clousures are an application of the use anonymous functions that references a variable that is declared outside of the anonymous function, they have a firm like this
``` go
func name() func() int
```
What this will return is a anonymous function that is declared inside the function that we called, closures have also to proporty that they keep referencing the variables that weren't passed as parameters.
Example
 ``` go
 //This function will return a new anonymous function each time is called
func NewCounter() func() int{
    n := 0
    //The return function will increase the counter and return the value,
    return func() int{
        n += 1
        return n
    }
}
func main(){
    counter := NewCounter()
    fmt.println(counter())
    fmt.println(counter())
}
```
output
``` text
1
2
```
The function abode will create an anonymous function that will keep referring to the variable n and will increate the count each time is called it will increment the counter inside the function. Each newCounter that is created it will have his own n variable. This allow to create data isolation.

Closure function can also refer variable variables.
