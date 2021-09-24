package main

import (
    "go_framework/framework"
    "net/http"
)

func main() {
    core := framework.NewCore()
    registerRouter(core)
    server := &http.Server{
        Handler: core,
        Addr:    ":8888",
    }
    server.ListenAndServe()
}
