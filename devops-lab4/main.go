package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    name := os.Getenv("APP_NAME")
    if name == "" {
        name = "Go MultiStage App"
    }
    fmt.Fprintf(w, "Hello from %s!\n", name)
}

func main() {
    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", handler)
    fmt.Printf("Server is running on port %s...\n", port)
    http.ListenAndServe(":"+port, nil)
}
