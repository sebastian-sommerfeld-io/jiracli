package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldCountUsers(t *testing.T) {
	// Given
	count := CountUsers()
	expectedCount := 7
	assert.Equal(t, expectedCount, count)
}
