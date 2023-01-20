package user

import (
	"errors"
)

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
