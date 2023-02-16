package jiracli

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateCmdRoot(t *testing.T) {
	expectedCmd := &cobra.Command{
		Use:  "jiracli",
		Args: cobra.ExactArgs(0),
	}

	got := NewCmdRoot()
	assert.NotNil(t, got)
	assert.Equal(t, expectedCmd.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.False(t, got.Runnable(), "Command should NOT be runnable")
}
