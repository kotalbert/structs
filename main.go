package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	alex := person{firstName: "Alex",
		lastName: "DeLarge",
		contact:  contactInfo{"alex@foobar.cor", 94000},
	}

	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
