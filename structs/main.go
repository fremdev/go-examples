package main

import (
	"fmt"
)

type contacts struct {
	zipCode int
	email   string
}

type person struct {
	firstName string
	lastName  string
	contacts
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contacts: contacts{
			zipCode: 45300,
			email:   "alex@vikings.com",
		},
	}

	alex.updateName("Alexander")
	alex.print()

	var bob person
	bob.firstName = "Bob"
	bob.lastName = "King"
	bob.contacts = contacts{email: "bob@gmail.com", zipCode: 90210}
	bob.print()
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
