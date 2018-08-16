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

type airport struct {
	Name      string
	Terminals uint8
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

	jfk := airport{
		Name:      "John F. Kennedy Internal Airport",
		Terminals: 8,
	}

	lax := airport{
		Name:      "Los Angeles International Airport",
		Terminals: 8,
	}

	mia := airport{
		Name:      "Miami International Airport",
		Terminals: 3,
	}

	cities := []city{newyork, chi, sanfran, dallas, seattle}
	airports := []airport{jfk, lax, mia}

	data := struct {
		Cities   []city
		Airports []airport
	}{
		Cities:   cities,
		Airports: airports,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
