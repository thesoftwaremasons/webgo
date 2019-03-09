package main

import (
	"html/template"
	"log"
	"os"
)

var tmp *template.Template

func init() {
	tmp = template.Must(tmp.ParseFiles("tpl.gohtml"))
}
func main() {
	sages := map[string]string{"India": "New Delhi", "Britian": "London", "Rwanda": "Kigali"}
	err := tmp.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatal(err)
	}
}
