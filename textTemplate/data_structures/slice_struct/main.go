package main

import (
	"html/template"
	"log"
	"os"
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

	newyork := city{
		City:  "New York",
		State: "NY",
	}

	chi := city{
		City:  "Chicago",
		State: "IL",
	}

	sanfran := city{
		City:  "San Francisco",
		State: "CA",
	}

	dallas := city{
		City:  "Dallas",
		State: "TX",
	}

	seattle := city{
		City:  "Seattle",
		State: "WA",
	}

	cities := []city{newyork, chi, sanfran, dallas, seattle}

	err := tpl.Execute(os.Stdout, cities)
	if err != nil {
		log.Fatalln(err)
	}
}
