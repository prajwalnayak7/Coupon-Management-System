package main

import (
    "fmt"
    "net/http"
)

func main() {
	fmt.Println("Server Starting at 8686")
    http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8686", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}