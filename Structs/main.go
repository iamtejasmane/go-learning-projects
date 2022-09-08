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
	jim := person{
		firstName: "Jim",
		lastName:  "Moriarty",
		contactInfo: contactInfo{
			email:   "jim@didyoumissme.com",
			zipCode: 443,
		},
	}
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) updateName(firstName string) {
	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
