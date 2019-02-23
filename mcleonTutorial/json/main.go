package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string
	Lastname  string
	age       int
}

func main() {
	p := person{"Sam", "Ogundipe", 30}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))

}
