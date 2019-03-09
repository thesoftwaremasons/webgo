package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func main() {
	var d hotdog
//step 1 ListenAndServer()	
	http.ListenAndServe(":8080", d)
}

//step 1 ServeHTTP(w http.ResponseWriter, r *http.Request)
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code in this func")
}
