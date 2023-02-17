package commands

import (
	"fmt"

	"github.com/sebastian-sommerfeld-io/jiracli/services/jira/users"
	"github.com/spf13/cobra"
)

// NewCmdUserCount initializes the `jiracli user count` command and its flags.
func NewCmdUserCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "count",
		Short: "Count all Jira users",
		Long:  "Count all users from a Jira instance.",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(users.CountUsers())
		},
	}

	AddMandatoryFlags(cmd)

	return cmd
}
