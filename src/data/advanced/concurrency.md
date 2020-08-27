---
title: "Concurrency"
tag: advanced
---
## Concurrency
### Go routines
Go Routines are lightweight threads that run in the same address space since this mean that the access to memory have to be synchronized.

To run a new routine it just need to use the keyword go
```go
go function()
 ```

### Channels
Using channels is the easy way to communicate between routines since they allow routines to syncronize without explicit locks or conditions by blocking the channel until the sender or receiver is ready.
The channels type is chan and they have to be initialized before can be used.
```go
c := make(chan int)
```
They also have a reserved operator for channels <-
```go
import "fmt"

func factorial(n int,c chan int){
    fact := 1:
    for i:=1; i<=n ; i++{
        fact = fact*i
    }
    c<-fact
}

func main(){
    c := make(chan, int)
    go factorial(5,c)
    fmt.Println(<-c)
}
```
Information is transmitted in or out the channel depending on the position of the channel and the arrow operator

### Buffers
Channels can have a buffered size that block the channel for send when is full or for receiving when is empty.
```go
c := make(chan int, 5)
```
Writing to a buffered channel will block if it is full.

### Range and close
A sender can send a signal to indicate that the channel will not be receiving any more updates. This can be done by the function close(c).
```go
package main

import (
 "fmt"
 "math/rand"
 "runtime"
 "sync"
 "time"
)

func init() {
 fmt.Printf("Using %d workers...\n", runtime.NumCPU())
 rand.Seed(time.Now().UnixNano())
}

func process(name int, data chan int, wg *sync.WaitGroup) {
 for i := range data {
 elapsed := (time.Duration(rand.Intn(5)) + 1) * time.Second
 time.Sleep(elapsed)
 fmt.Printf("Worker %d) Processed data: %d in %d[ms]\n", name, i, elapsed/1000000)
 }
 wg.Done()
}

func main() {
 wg := &sync.WaitGroup{}

 data := make(chan int)

 for i := 0; i < runtime.NumCPU(); i++ {
   wg.Add(1)
   go process(i+1, data, wg)
 }

 for i := 0; i < 64; i++ {
  data <- i + 1
 }
 close(data)
 wg.wait()
}
```
In this example we create a wait group, this will help to keep track of the routines, this strcut has two methods, `add(int)` and `wait()` we add 1 for every routine called then we waited for the routines to finish.

To check if a channel has been closet it can be done by adding a variable when assigning a value, similar to checking for a key existing in a map.
```go
 v, ok := <-c
```
ok is false if there are no more values to receive and the channel is closed.

### Select
The select keyword can be used to wait until a case can be executed
```go
package main

import (
    "fmt"
    "sync"
)

func selector(values chan int, runes chan rune, wg *sync.WaitGroup) {
    sum := 0
    myString := ""
    for {
        select {
        case value := <-values:
            sum += value
            fmt.Printf("Sum value: %d\n", sum)
            if sum == 20 {
                defer wg.Done()
                return
            }
        case letter := <-runes:
            myString += string(letter)
            fmt.Printf("String value: %s\n", myString)
        }
    }
}

func main() {
    values := make(chan int)
    runes := make(chan rune)
    letters := []rune("abcdfghijklm")
    wg := &sync.WaitGroup{}
    wg.Add(1)
    go selector(values, runes, wg)
    values <- 5
    values <- 5
    values <- 5
    for i := 0; i < len(letters); i++ {
        runes <- letters[i]
    }
    values <- 5
    wg.Wait()
}
```

### Locks
Go standard library offers a way to provided a easy way to avoid conflicts in routines. this can me done with sync.Mutex that to methods: **Lock()**, **Unlock()**, they allow only one routine to have access to the block of code. This can be used to create atomic operations and avoid race conditions.

Since this locks **_block the resource for reading and writing_**, the use of a **RWMutex** is recommended, since they allow to block the channel for reading or writing or for both. This can also be found in sync.RWMutex and have the methods **Lock()**, **Unlock()** for writing and **RLock()** and **RUnlock()** for reading.
