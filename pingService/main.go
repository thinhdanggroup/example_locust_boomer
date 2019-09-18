package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/ping", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "pong")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8181", nil)
}