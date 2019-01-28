package main

import (
	"fmt"
	"net/http"
)

//request response
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Hello About</p>")
}

//func index(w http.Res)
func main() {
	fmt.Println("ya")
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":3000", nil)

}
