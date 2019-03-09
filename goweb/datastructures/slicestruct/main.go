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

var tmp *template.Template

func init() {
	tmp = template.Must(tmp.ParseFiles("tpl.gohtml"))
}
func main() {
	var p = Person{Firstname: "Sammy", Surname: "James", Age: 30}
	var d = Person{Firstname: "mmy", Surname: "James", Age: 40}
	var a = Person{Firstname: "Samy", Surname: "James", Age: 30}
	//sages := map[string]string{"India": "New Delhi", "Britian": "London", "Rwanda": "Kigali"}
	all := []Person{p, d, a}
	err := tmp.Execute(os.Stdout, all)
	if err != nil {
		log.Fatal(err)
	}
}
