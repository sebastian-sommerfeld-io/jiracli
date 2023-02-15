package jiracli

import (
	"github.com/spf13/cobra"
)

// NewCmdUserView initializes the `jiracli license` command.
func NewCmdLicense() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "license",
		Short: "Interact with a Jira instance and get license information",
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}