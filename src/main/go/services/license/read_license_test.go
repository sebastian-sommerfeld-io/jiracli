package license

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldGetLicense(t *testing.T) {
	baseUrl := "http://localhost:8080"
	user := "admin"
	pass := "admin"

	result, err := ReadJiraLicense(baseUrl, user, pass)

	assert.Nil(t, err)
	assert.True(t, json.Valid([]byte(result.RawJson)), "JSON should be valid")
}
