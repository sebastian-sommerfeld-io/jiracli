package user

import (
	"errors"
)

type User struct {
	Id        int
	Firstname string
	Lastname  string
	Username  string
	Email     string
}

var users = []User{}

func init() {
	users = []User{
		User{
			Id:        1,
			Firstname: "Admin",
			Lastname:  "Admin",
			Username:  "admin.admin",
			Email:     "admin@localhost",
		},
		User{
			Id:        2,
			Firstname: "Jim",
			Lastname:  "Panse",
			Username:  "jim.panse",
			Email:     "jim.panse@localhost",
		},
		User{
			Id:        3,
			Firstname: "Baby",
			Lastname:  "Kong",
			Username:  "baby.kong",
			Email:     "baby.kong@localhost",
		},
	}
}

func FindByUsername(username string) (User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	user := User{}
	return user, errors.New("No user found")
}

func FindByEmail(email string) (User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	user := User{}
	return user, errors.New("No user found")
}
