package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func userFoundAssertions(t *testing.T, user User, err error) {
	expectedUser := User{
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

func userNotFoundAssertions(t *testing.T, user User, err error) {
	expectedMessage := "no user found"

	assert.EqualErrorf(t, err, expectedMessage, "Error should be: %v, got: %v", expectedMessage, err)
	assert.Equal(t, User{}, user)
}

func Test_ShouldFindByUsername(t *testing.T) {
	search := "jim.panse"
	user, err := FindByUsername(search)
	userFoundAssertions(t, user, err)
}

func Test_ShouldNotFindByUsername(t *testing.T) {
	search := "non.existing"
	user, err := FindByUsername(search)
	userNotFoundAssertions(t, user, err)
}

func Test_ShouldFindByEmail(t *testing.T) {
	search := "jim.panse@localhost"
	user, err := FindByEmail(search)
	userFoundAssertions(t, user, err)
}

func Test_ShouldNotFindByEmail(t *testing.T) {
	search := "non.existing@localhost"
	user, err := FindByEmail(search)
	userNotFoundAssertions(t, user, err)
}
