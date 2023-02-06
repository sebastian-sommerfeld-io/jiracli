package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var AssertionUser = User{}

func init() {
	AssertionUser = User{
		Id:        2,
		Firstname: "Jim",
		Lastname:  "Panse",
		Username:  "jim.panse",
		Email:     "jim.panse@localhost",
	}
}

func userAssertions(t *testing.T, user User, err error) {
	assert.Equal(t, nil, err)
	assert.Equal(t, AssertionUser.Id, user.Id)
	assert.Equal(t, AssertionUser.Firstname, user.Firstname)
	assert.Equal(t, AssertionUser.Lastname, user.Lastname)
	assert.Equal(t, AssertionUser.Username, user.Username)
	assert.Equal(t, AssertionUser.Email, user.Email)
}

func Test_ShouldFindByUsername(t *testing.T) {
	search := "jim.panse"               // Given
	user, err := FindByUsername(search) // When
	userAssertions(t, user, err)        // Then
}

func Test_ShouldFindByEmail(t *testing.T) {
	search := "jim.panse@localhost"  // Given
	user, err := FindByEmail(search) // When
	userAssertions(t, user, err)     // Then
}
