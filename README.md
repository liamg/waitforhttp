# waitforhttp
[![Travis](https://img.shields.io/travis/liamg/waitforhttp.svg?style=flat-square)](https://travis-ci.org/liamg/waitforhttp)

Waits for an HTTP server to be serving before returning.

Example:
```go
package main

import (
    "fmt"
	"net/http"
    "time"

    "github.com/liamg/waitforhttp"
)

func main() {
    
    server := &http.Server{
        Addr: ":8080",
    }
    
    go func() {
        if err := waitforhttp.Wait(server, time.Second*10); err != nil {
            panic(err)
        }   
        fmt.Println("Hooray! Server is listening now!")
    }()
    
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        panic(fmt.Sprintf("Failed to start server: %s", err))
    }
}
```
