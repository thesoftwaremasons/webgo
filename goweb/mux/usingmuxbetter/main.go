package main

import (
	"io"
	"net/http"
)

type handdog int
type handcat int

func main() {
	// var d handdog
	// var c handcat
	//mux := http.NewServeMux()
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8080", nil)
}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Works for dogs")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Works for cat")
}
