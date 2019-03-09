package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	tmp, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	defer tmp.Close()
	bs, err := ioutil.ReadAll(tmp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(bs))
}
