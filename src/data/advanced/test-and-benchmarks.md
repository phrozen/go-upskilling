---
title: "Test and benchmarks"
tags: advanced
---
### Test and benchmarks

### Test
The go tool provide his own automated testing for go source code, it is run by the command
```bash
go test
```
And will search for the files ended with _test.go and for the test cases all routines that have the following format
func TestXxxx(*testing.T)

The methods ErrorF, Fatal, Error or any other related method can be used to signal a non success test, all of those methods are declared in the package testing

### Benchmarks

Benchmarks are a type of test that is run b.N times until it can be long enough to be properly  timed. The variable b.N is auto adjusted until it have the correct size.
Benchmarks follow the following expression:
func BenchmarkXxx(*testing.T)

benchmarks are rune with the go tool:
```bash
go test -bench
```
