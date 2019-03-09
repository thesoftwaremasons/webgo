package main

import (
	"html/template"
	"os"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "tpl.gohtml", 32)
	if err != nil {
		panic(err)
	}
}
