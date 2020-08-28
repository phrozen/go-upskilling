---
title: "Errors"
tags: intermediate
---
## Errors
The package errors contains the interface error and the method New  that return a error instance.
```go
type error interface {
Error() string
}
```
By convention, errors are the last return value and have type error, if the function has no error is common to return a **nil** otherwise errors.New return a new error with the given message.
