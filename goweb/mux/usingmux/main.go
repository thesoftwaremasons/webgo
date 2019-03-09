package main

import (
	"io"
	"net/http"
)

type handdog int
type handcat int

func main() {
	var d handdog
	var c handcat
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)
}

func (d handdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Works for dogs")
}

func (c handcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Works for cat")
}
