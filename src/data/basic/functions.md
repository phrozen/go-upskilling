---
title: "Functions"
tags: basic
---
## Funtions
Functions in Go  start with the keyword “func” followed by the name of function, input values and return values

Example
```go
func example()
func example(x int) int
func example(a, _ int, z float32) bool
func example(a, b int, z float32) (bool)
func example(prefix string, values ...int)
func example(a, b int, z float64, opt ...interface{}) (success bool)
func example(int, int, float64) (float64, *[]int)
```
Returning values is done with the keyword “return”, it can return multiple values if the function expects multiple return values

## Defer
Defer are instructions that are always executed at the end of the function, they can be stacked. When something is deferred it is not executed in that moment, it enters a stack and is keep it there until the function ends or a early return is find, then there are executed in last-in first-out order.
```go
package main

import "fmt"

func main() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
```
The example above will print the following:
```plain
counting
done
9
8
7
6
5
4
3
2
1
0
```
