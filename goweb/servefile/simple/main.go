package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", gone)
	http.ListenAndServe(":8080", nil)
}

func gone(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="https://commons.wikimedia.org/wiki/File:Golde33443.jpg">`)
}
