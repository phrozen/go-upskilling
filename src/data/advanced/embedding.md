---
title: "Embedding"
tag: advanced
---
## Embedding structs
The official documentation describe a type struct as:

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField).

So, in other words embedding is another way of declaring fields into a new struct, this can be done just by adding the desire struct in the declaration of the new struct. When a struct is embedded in another, the new struct will have access to all the fields and methods that the first one has, this meaning that if struct A, implemented interfaces IA, IB and IC, then struct A is embedded in struct B, struct B also implement interfaces IA, IB and IC.

Example:

```go
type person struct{
      name string
      lastName string
      age int
}

type employee struct{
      person
      salary float32
      company string
}

type hello interface{
      sayHello()
}

type invoiceInfo interface{
      getName() string
      getCompany() string
}

type formalHello interface{
      sayFormalHello()
}

func (p *person) sayhello(){
            fmt.Println("Hello my name is ", p.name)
}

func (e *eemployee) sayFormalHello(){
      fmt.Println("Good Morning everyone, my name is ",e.name)
}

func (p *person) getName() string{
      return p.name
}

func (e * employee) getCompany() string{
      return e.company
}
```
In this example we have 2 structs person and employee, and we can see that the employee struct have an embedded type of persona, this mean that all the instances of employee can access to the fields name, lastname and age.

Also both structs implement the interface  hello, since person has the method implemented of sayHello() and employee can access this method both suffice this interface. But employee also satisfice the interface formalHello and invoice
