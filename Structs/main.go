package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var tejas person
	tejas.firstName = "Tejas"
	tejas.lastName = "Mane"
	fmt.Println(tejas)
	fmt.Printf("%+v", tejas)
}
