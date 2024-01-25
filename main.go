package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func newUser(firstName string, lastName string, birthDate string) user {
	appUser := user{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}
	return appUser
}

func (u user) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	fn := getUserData("Please enter your first name: ")
	ln := getUserData("Please enter your last name: ")
	bd := getUserData("Please enter your bd (MM/DD/YYYY): ")

	appUser := newUser(fn, ln, bd)

	appUser.outputUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, _ = fmt.Scan(&value)
	return value
}
