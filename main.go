package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {

	alex := person{firstName: "Alex", lastName: "DeLarge"}
	var billie person

	billie.firstName = "Billy"
	billie.lastName = "Boy"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	fmt.Printf("%+v", billie)
}
