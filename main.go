package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, world2!"))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
