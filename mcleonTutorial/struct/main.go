package main

import "fmt"

type num int

type person struct {
	firstname string
	lastname  string
	sex       string
}

func main() {
	var noa num
	p1 := person{"Femi", "Samuel", "Male"}

	fmt.Printf("%T", noa)
	fmt.Print(p1.firstname)
}
