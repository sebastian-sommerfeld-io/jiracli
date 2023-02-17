package commands

import (
	"testing"

	"github.com/sebastian-sommerfeld-io/jiracli/services/jira/users"
	"github.com/stretchr/testify/assert"
)

func Test_EmailShouldBeValid(t *testing.T) {
	email := "jim.panse@localhost"
	valid, err := validateEmail(email)
	assert.Nil(t, err)
	assert.True(t, valid)
}

func Test_EmailShouldBeInvalid(t *testing.T) {
	email := "jim.panse"
	valid, err := validateEmail(email)
	assert.False(t, valid)
	assert.NotNil(t, err)
}

func userFoundAssertions(t *testing.T, user users.User, err error) {
	expectedUser := users.User{
		Id:        2,
		Firstname: "Jim",
		Lastname:  "Panse",
		Username:  "jim.panse",
		Email:     "jim.panse@localhost",
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedUser.Id, user.Id)
	assert.Equal(t, expectedUser.Firstname, user.Firstname)
	assert.Equal(t, expectedUser.Lastname, user.Lastname)
	assert.Equal(t, expectedUser.Username, user.Username)
	assert.Equal(t, expectedUser.Email, user.Email)
}

func userNotFoundAssertions(t *testing.T, user users.User, err error) {
	assert.NotNil(t, err)
	assert.Equal(t, users.User{}, user)
}

func Test_ShouldGetUserByEmail(t *testing.T) {
	opts := &userViewOptions{
		ByEmail: true,
		Search:  "jim.panse@localhost",
	}
	user, err := getUser(opts)
	userFoundAssertions(t, user, err)
}

func Test_ShouldNotGetUserByEmail(t *testing.T) {
	opts := &userViewOptions{
		ByEmail: true,
		Search:  "non.existing@localhost",
	}
	user, err := getUser(opts)
	userNotFoundAssertions(t, user, err)
}

func Test_ShouldGetErrorDueToInvalidEmail(t *testing.T) {
	opts := &userViewOptions{
		ByEmail: true,
		Search:  "invalid-email",
	}
	user, err := getUser(opts)
	userNotFoundAssertions(t, user, err)
}

func Test_ShouldGetUserByUsername(t *testing.T) {
	opts := &userViewOptions{
		ByUsername: true,
		Search:     "jim.panse",
	}
	user, err := getUser(opts)
	userFoundAssertions(t, user, err)
}

func Test_ShouldNotGetUserByUsername(t *testing.T) {
	opts := &userViewOptions{
		ByUsername: true,
		Search:     "non.existing",
	}
	user, err := getUser(opts)
	userNotFoundAssertions(t, user, err)
}
