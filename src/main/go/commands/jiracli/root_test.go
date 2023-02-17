package jiracli

import (
	"bytes"
	"io"
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

func Test_ShouldPrintHelpText(t *testing.T) {
	got := NewCmdRoot()
	outputStream := bytes.NewBufferString("")
	got.SetOut(outputStream)
	got.SetErr(outputStream)

	assert.NotNil(t, got)

	got.Execute()
	out, err := io.ReadAll(outputStream)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, string(out))
	assert.Contains(t, string(out), got.Long)
}
