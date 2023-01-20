package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findByAssertions(t *testing.T, user User, e error) {
	verificationUser := User{
		Id:        2,
		Firstname: "Jim",
		Lastname:  "Panse",
		Username:  "jim.panse",
		Email:     "jim.panse@localhost",
	}

	assert.Equal(t, nil, e)
	assert.Equal(t, verificationUser.Id, user.Id)
	assert.Equal(t, verificationUser.Firstname, user.Firstname)
	assert.Equal(t, verificationUser.Lastname, user.Lastname)
	assert.Equal(t, verificationUser.Username, user.Username)
	assert.Equal(t, verificationUser.Email, user.Email)
}

func Test_FindByUsername(t *testing.T) {
	in := "jim.panse"
	user, e := FindByUsername(in)
	findByAssertions(t, user, e)
}

func Test_FindByEmail(t *testing.T) {
	in := "jim.panse@localhost"
	user, e := FindByEmail(in)
	findByAssertions(t, user, e)
}
