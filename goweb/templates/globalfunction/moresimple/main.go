package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	slice := []string{"Femi", "Samuel", "Ayobami", "Olusegun", "Bamidele"}

	data := struct {
		Words []string
		Lname string
	}{
		slice,
		"Ogundipe",
	}  
	err := tmp.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

}
