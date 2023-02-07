package jiracli

import (
	"github.com/spf13/cobra"
)

var (
	userCmd = &cobra.Command{
		Use:   "user",
		Short: "Interact with a Jira instance and work with Jira users",
		Args:  cobra.ExactArgs(0),
	}
)

func init() {
	rootCmd.AddCommand(userCmd)
}
