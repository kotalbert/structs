package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Constructor
// Input validation
func newUser(firstName string, lastName string, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstname, lastname, birthdate are expecting non empty string")
	}

	return &user{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}, nil
}

func (u user) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	fn := getUserData("Please enter your first name: ")
	ln := getUserData("Please enter your last name: ")
	bd := getUserData("Please enter your bd (MM/DD/YYYY): ")

	appUser, err := newUser(fn, ln, bd)

	if err != nil {
		log.Fatal(err)
	}

	appUser.outputUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, _ = fmt.Scanln(&value)
	return value
}
