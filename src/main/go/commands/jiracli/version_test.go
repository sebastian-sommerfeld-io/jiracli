package jiracli

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateCmdVersion(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "version",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdVersion()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.True(t, got.Runnable(), "Command should be runnable")
}
