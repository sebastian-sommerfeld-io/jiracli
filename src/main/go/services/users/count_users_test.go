package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldCountUsers(t *testing.T) {
	// Given
	count := CountUsers()                 // When
	expectedCount := 7                    // Then
	assert.Equal(t, expectedCount, count) // Then
}
