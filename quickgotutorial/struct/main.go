package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// surName   string
	// age       int
	// gender    string
	// city      string
	firstName, surName, gender, city string
	age                              int
}

//create a value reciever
func (p Person) greet() string {
	return p.firstName + " " + p.surName + "" + strconv.Itoa(p.age)
}
func (p *Person) increment() {
	p.age++
}
func (p *Person) checkFemale() {
	if p.gender == "f" {
		p.firstName = "ya"
	}

}
func main() {
	// Initiallise struct
	person1 := Person{firstName: "James", surName: "Lannister", age: 35, gender: "f", city: "Lagos"}
	person1.checkFemale()
	fmt.Print(person1.greet())
}
