package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Go!")
}

func main() {
    http.HandleFunc("/", handler)
    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
