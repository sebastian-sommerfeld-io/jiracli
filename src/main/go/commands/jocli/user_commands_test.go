package jocli

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"sebastian-sommerfeld-io/jira-ops-cli/services/user"
)

func Test_viewUser(t *testing.T) {
	verificationUser := user.User{
		Id:        2,
		Firstname: "Jim",
		Lastname:  "Panse",
		Username:  "jim.panse",
		Email:     "jim.panse@localhost",
	}

	user := viewUser("jim.panse")
	assert.Equal(t, verificationUser.Id, user.Id)
	assert.Equal(t, verificationUser.Firstname, user.Firstname)
	assert.Equal(t, verificationUser.Lastname, user.Lastname)
	assert.Equal(t, verificationUser.Username, user.Username)
	assert.Equal(t, verificationUser.Email, user.Email)
}
