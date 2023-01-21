package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Incoming message ...")
    fmt.Fprintf(w, "Hello there, I looove %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("started ....")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
