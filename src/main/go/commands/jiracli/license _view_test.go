package jiracli

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldGetLicense(t *testing.T) {
	opts := &licenseViewOptions{
		BaseURL: "http://localhost:8080",
		User:    "admin",
		Pass:    "admin",
	}

	jsonResult, err := getJiraSoftwareLicense(opts)

	assert.Nil(t, err)
	assert.True(t, json.Valid([]byte(jsonResult)), "JSON should be valid")
}

func Test_ShouldGetUnauthorizedError(t *testing.T) {
	opts := &licenseViewOptions{
		BaseURL: "http://localhost:8080",
		User:    "cloud",
		Pass:    "cloud",
	}
	expectedMessage := "must have permission to access this resource"

	jsonResult, err := getJiraSoftwareLicense(opts)

	assert.NotNil(t, err)
	assert.EqualErrorf(t, err, expectedMessage, "Error should be: %v, got: %v", expectedMessage, err)
	assert.Equal(t, "", jsonResult)
}

func Test_ShouldGetConnectionRefusedError(t *testing.T) {
	opts := &licenseViewOptions{
		BaseURL: "http://localhost:8181",
		User:    "admin",
		Pass:    "admin",
	}
	expectedMessage := "dial tcp 127.0.0.1:8181: connect: connection refused"

	jsonResult, err := getJiraSoftwareLicense(opts)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), expectedMessage)
	assert.Equal(t, "", jsonResult)
}
