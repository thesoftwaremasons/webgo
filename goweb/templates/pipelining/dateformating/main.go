package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}

var fm = template.FuncMap{
	"fdateDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}
