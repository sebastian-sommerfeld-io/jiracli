package jiracli

import (
	"github.com/spf13/cobra"
)

// NewCmdUser initializes the `user` command.
func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Interact with a Jira instance and work with Jira users",
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}
