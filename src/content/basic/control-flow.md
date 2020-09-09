---
title: "Control Flow"
tags: basic
index: 6
---

## Control Flow

### if/else

`if/else` syntax is the same as java or c, but without the () and the {} are required.

``` go
if a == b {
    // ...
} else if a == c{
    // ...
} else
    // ...
}
```

Also it is possible to create a variable to have a initialization statement at the start of the
`if`:

``` go
if val, ok := myMap[324]; ok{
    //...
}
```

In the example abode the if will check if the map contains the key.

### for
For cycle is almost the same as C/Java, but the () are gone and the {} is required, also the for in Go can be used as a while.

#### Simple `for` loop
``` go
for i := 0; i < 10; i++ {
    // ...
}
```

It can be used like a while, with only the condition

``` go
for i < 10 {
    i++
}
```

Or it can be used to have an infinity loop

``` go
 for {
    // ...
}
```

### Switch
Switches are very similar to C/Java but instead of using the “break” keyword the code only execute one case and will not continue with the following cases unless the keyword “fallthrough” is used. Some other rules for the switch are that the cases can contain multiple values, () are removed from the syntax and {} are mandatory.

Example

``` go
switch integer {
    case 1:
        // code1
    case 2, 3, 5, 7:
        // code2
        fallthrough
    case 4:
        // code3
    default:
        // code4
}
```

In this example if condition1 is meet only code1 is run, but if condition2 or condition3 is meet the code2 and code3 is executed, but for condition4 just code3 is executed, if none condition is meet code4 is executed.
