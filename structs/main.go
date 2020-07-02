package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Max",
		lastName:  "sdg",
		contactInfo: contactInfo{
			email:   "email",
			zipCode: 123,
		},
	}
	alex.updateName("Alex")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
