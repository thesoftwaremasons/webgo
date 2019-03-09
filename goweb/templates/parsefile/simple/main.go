package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	temp, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Println("Reason:", err)
	}
	err = temp.Execute(os.Stdout, nil)
	if err != nil {
		log.Println("Reason:", err)
	}
}
