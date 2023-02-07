package users

import (
	"errors"
)

// FindByUsername reads a single user from a Jira instance through the Jira API.
func FindByUsername(username string) (User, error) {
	for _, user := range dummyUserList {
		if user.Username == username {
			return user, nil
		}
	}
	user := User{}
	return user, errors.New("no user found")
}

// FindByEmail reads a single user from a Jira instance through the Jira API.
func FindByEmail(email string) (User, error) {
	for _, user := range dummyUserList {
		if user.Email == email {
			return user, nil
		}
	}
	user := User{}
	return user, errors.New("no user found")
}
