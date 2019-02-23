package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	sex       string
}
type doubleZero struct {
	person
	isActive bool
}

func (p person) greetings() string {
	return p.firstname
}
func (z doubleZero) greetings() person {
	return z.person
}

func main() {
	firstval := doubleZero{
		person: person{
			firstname: "james",
			lastname:  "lannister",
			sex:       "Male",
		},
		isActive: true,
	}
	secondval := doubleZero{
		person: person{
			firstname: "james",
			lastname:  "Bond",
			sex:       "Male",
		},
		isActive: true,
	}
	greet := firstval.greetings()
	greets := secondval.greetings()
	fmt.Println(greet)
	fmt.Println(greets)
}
