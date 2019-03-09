package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("index.gohtml"))
}
func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatal(err)
	}
}
