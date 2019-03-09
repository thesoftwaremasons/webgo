package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Println(err)
	}
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Print(err)
	}
	defer newFile.Close()

	err = tpl.Execute(newFile, nil)
	if err != nil {
		log.Println(err)
	}

}
