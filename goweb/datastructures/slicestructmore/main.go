package main

import (
	"html/template"
	"log"
	"os"
)

type Person struct {
	Firstname string
	Surname   string
	Age       int
}
type Animal struct {
	Name string
}

var tmp *template.Template

func init() {
	tmp = template.Must(tmp.ParseFiles("tpl.gohtml"))
}
func main() {
	var p = Person{Firstname: "Sammy", Surname: "James", Age: 30}
	var d = Person{Firstname: "mmy", Surname: "James", Age: 40}
	var a = Person{Firstname: "Samy", Surname: "James", Age: 30}
	var an = Animal{Name: "Cow"}
	var ani = Animal{Name: "Cat"}
	//sages := map[string]string{"India": "New Delhi", "Britian": "London", "Rwanda": "Kigali"}
	allPersons := []Person{p, d, a}
	allAnimals := []Animal{an, ani}
	anon := struct {
		Person []Person
		Animal []Animal
	}{
		allPersons,
		allAnimals,
	}

	err := tmp.Execute(os.Stdout, anon)
	if err != nil {
		log.Fatal(err)
	}
}
