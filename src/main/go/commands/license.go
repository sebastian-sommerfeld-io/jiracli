package commands

import (
	"github.com/spf13/cobra"
)

// NewCmdLicense initializes the `jiracli license` command.
func NewCmdLicense() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "license",
		Short: "Interact with a Jira instance and get license information",
		Long:  "Interact with a Jira instance and get license information in JSON format",
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}
