package commands

import (
	"github.com/spf13/cobra"
)

// NewCmdUser initializes the `jiracli user` command.
func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Interact with a Jira instance and get Jira user information",
		Long:  "Interact with a Jira instance and count users or access information about a Jira user",
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}
