package main

import (
	"fmt"
	"kotalbert/structs/user"
	"log"
)

func main() {
	fn := getUserData("Please enter your first name: ")
	ln := getUserData("Please enter your last name: ")
	bd := getUserData("Please enter your bd (MM/DD/YYYY): ")

	appUser, err := user.NewUser(fn, ln, bd)

	if err != nil {
		log.Fatal(err)
	}

	appUser.OutputUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, _ = fmt.Scanln(&value)
	return value
}
