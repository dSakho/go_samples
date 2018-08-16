package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type city struct {
	City  string
	State string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	nyc := city{
		City:  "New York",
		State: "NY",
	}

	err := tpl.Execute(os.Stdout, nyc)
	if err != nil {
		log.Fatalln(err)
	}
}
