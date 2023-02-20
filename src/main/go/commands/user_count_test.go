package commands

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateCmdUserCount(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "count",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdUserCount()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.True(t, got.Runnable(), "Command should be runnable")
}
