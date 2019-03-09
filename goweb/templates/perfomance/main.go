package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("perfomance/*"))
}
func main() {
	err := tmp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("touch:", nil)
	}

}
