package main

import (
	"io"
	"net/http"
)

type hotdog int

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "cool dog")
	case "/goat":
		io.WriteString(w, "nice goat")

	}
}
