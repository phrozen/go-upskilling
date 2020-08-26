
## Arrays
Arrays are of specific types and length, their length is **fixed** it cannot grow or reduced, arrays are values in Go, this meaning that the array variable is all the array, instead of the pointer at the first element of the array.  
Every array length type is the same, it will make [3]T a different type from [4]T.  
To declare an Array the following syntax is used:
```go
var name int[3]
      arr := [3]int{1,2,3}˚
      arr := [...]int{1,2,3,4,5,6}
```
Using "..." in the declaration of the array will instruct the compiler to count the variables at moment of compilation.
## Slices
Slices in the hand are a more versatile type of data, a slice consist in the following:  
- A pointer to an array
- Length of the segment (elements in the array)
- Capacity of elements stored.  

To create a slice the function syntax is used:
```go
make([]T, len, cap)
```
It will return an slice of type T ( []T ). The length and capacity of the slice can query with the funcitons `len(s)` and `cap(s)`.  

Slices can dynamical grow to adjust to the desired length of the slice. This is done internally by the append function, when the new array index is bigger than the capacity of the slice, it will create and copy elements to the new array.
```go
func main() {
     s := make([]int,0,3)
     for i := 0; i < 7; i++ {
         s = append(s, i)
    fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
     }
}
```
Output
```plain
cap 3, len 1, 0xc0000b8000
cap 3, len 2, 0xc0000b8000
cap 3, len 3, 0xc0000b8000
cap 6, len 4, 0xc0000ac030
cap 6, len 5, 0xc0000ac030
cap 6, len 6, 0xc0000ac030
```

We can see in the example above the function append will duplicate the capacity of the

The slices have 3 built-in functions  
-`len()`: It will return the actual length of the array in the slice  
-`cap()`: It will return the max capacity with of the array in the slice  
-`append()`:  this will append the element to the array, if the array capacity will be smaller than the new size of the array, this function will create a new array with the double of the capacity and copy the array to it so it can be appended the new element

## Map
Go implementation of the Map uses a hash table, it have to follow this rule: the keytype of the hashmap has to be comparable type.  
The map has the following built interfaces:  
- `len()`: return the number of elements in a map
- `make(map[string]int)`: Function to create new maps, in this case it will create a new map with the keytype string and the value will be a integer
- `delete(m,"key")`: Function to delete a entry in a map.

Maps can be checked if a key exists using a two initialization variables
```go
myMap :=make(map[int]string)
...
val, ok := myMap[324]
```

The variable val will be assigned the value of the content of the hashmap if it contain the key, and the variable ok will be assigned true or false depending if the map content the key
