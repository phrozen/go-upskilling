---
  title: "Interfaces"
  tags: intermediate
---
## Interfaces
Interfaces help specifying the behaviour of the objects. For an object to implement an interface it should follow one simple rule, it must suffice all methods declared in the interface.

An object can implement multiple interfaces.

Declaring is done with the keyword interface as follow:
```go
type Store interface{
    Get(string) (string, error)
    Set(string, string) error
    Delete(string) error
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```
In this example, if a struct wants to implement the **interface** **Writer**, it should implement the method **Write**
