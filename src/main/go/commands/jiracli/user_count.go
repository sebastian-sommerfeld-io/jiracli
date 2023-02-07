package jiracli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/sebastian-sommerfeld-io/jiracli/services/users"
)

var (
	userCountCmd = &cobra.Command{
		Use:   "count",
		Short: "Count all Jira users",
		Long:  "Count all users from a Jira instance.",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(users.CountUsers())
		},
	}
)

func init() {
	userCmd.AddCommand(userCountCmd)
}
