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

func (u user) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	fn := getUserData("Please enter your first name: ")
	ln := getUserData("Please enter your last name: ")
	bd := getUserData("Please enter your bd (MM/DD/YYYY): ")

	appUser := user{
		fn,
		ln,
		bd,
		time.Now(),
	}

	appUser.outputUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, _ = fmt.Scan(&value)
	return value
}
