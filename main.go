package main

import (
	"net/http"

	"example.com/go-test/biz"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(biz.GetHello()))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
