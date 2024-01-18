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
	alex.updateName("Aleksej")
	alex.print()
}

func (p *person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
	// equivalent syntax
	//(*p).firstName = newFirstName
}
