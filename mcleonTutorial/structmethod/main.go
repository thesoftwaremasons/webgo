package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	sex       string
}

func (p person) fullname() string {
	return p.firstname + "" + p.lastname
}

func main() {

	p1 := person{"Femi", "Samuel", "Male"}
	x := p1.fullname()
	fmt.Print(x)
}
