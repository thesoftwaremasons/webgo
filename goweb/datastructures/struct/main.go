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
	//sages := map[string]string{"India": "New Delhi", "Britian": "London", "Rwanda": "Kigali"}
	err := tmp.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal(err)
	}
}
