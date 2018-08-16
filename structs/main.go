package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int32
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Daouda",
		lastName:  "Sakho",
		contactInfo: contactInfo{
			email:   "d.sakho89@gmail.com",
			zipCode: 99999,
		},
	}

	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
