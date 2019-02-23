package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		log.Fatalln(err)
	}
	val, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(val))

}
