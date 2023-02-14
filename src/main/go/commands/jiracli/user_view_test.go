package jiracli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EmailShouldBeValid(t *testing.T) {
	email := "admin@localhost"         // Given
	valid, err := validateEmail(email) // When
	assert.Nil(t, err)                 // Then
	assert.True(t, valid)              // Then
}

func Test_EmailShouldBeInvalid(t *testing.T) {
	email := "admin.admin"             // Given
	valid, err := validateEmail(email) // When
	assert.False(t, valid)             // Then
	assert.NotNil(t, err)              // Then
}
