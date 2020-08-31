---
title: "Web server"
tag: next
index: 3
---

## Web server

The package net/http allows to handle http requests, it allows to easy create and edit endpoint for a web server, but some the of the methods that offer are a basic and will require to write more code than using a web server framework.

 ``` go
package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
)

// This is a handler that takes Requests and writes Responses
func index(res http.ResponseWriter, req *http.Request) {
    // Log just a bit of the request to Stdout
    log.Println(req.Method, req.URL, req.Proto)
    // Write a message to the ResponseWriter
    fmt.Fprintln(res, "Welcome to the Go server example.")
}
func main() {
    log.Println("Golang Server Example")
    // Set environment variable PORT to any available port (ie. 3000)
    if os.Getenv("PORT") == "" {
        log.Fatalln("Set the $PORT environment variable to run the server (ie. export PORT=3000)")
    }
    // Serve the root "/" path with the 'index' handler
    http.HandleFunc("/", index)
    // Log some more...
    log.Println("running server on port:", os.Getenv("PORT"))
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
```

Using this package will allow us to use all the HTTP methods, but it will need to a switch for each each the methods and one handle for each of the paths.

There are some frameworks that help simplification this issue:
- **Caddy server**
https://caddyserver.com/v2
- **Gorilla Toolkit**
https://github.com/gorilla/
- **Echo**
https://echo.labstack.com/
- **Fiber**
https://github.com/gofiber/fiber
- **Chi**
https://github.com/go-chi/chi
