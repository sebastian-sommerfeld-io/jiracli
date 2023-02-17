package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldReadAppInfo(t *testing.T) {
	info := JiracliAppInfo{}
	got, err := info.Read()

	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.NotNil(t, got.Version)
}
