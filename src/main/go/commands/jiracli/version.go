package jiracli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCmdVersion initializes the `jiracli version` command.
func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show jiracli version.",
		Long:  "Show jiracli version (not the version of any Jira instance).",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Some Version as JSON")
		},
	}

	return cmd
}
