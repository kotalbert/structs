package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// NewUser Constructor
// Input validation
func NewUser(firstName string, lastName string, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstname, lastname, birthdate are expecting non empty string")
	}

	return &User{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}, nil
}

func (u User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
