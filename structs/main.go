package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int16
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Carey",
		contactInfo: contactInfo{
			email:   "jim.carey@gmail.com",
			zipCode: 12345,
		},
	}

	jim.updateName("James")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
