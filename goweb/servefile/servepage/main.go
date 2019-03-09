package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir('.')))
}

// func dog(res http.ResponseWriter, req *http.Request) {

// }
