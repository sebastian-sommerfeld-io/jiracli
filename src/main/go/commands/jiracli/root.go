package jiracli

import (
	"github.com/spf13/cobra"
)

// NewCmdRoot initializes the `jiracli` root command.
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jiracli",
		Short: "A command line interface to access Jira and automate recurring tasks",
		Long:  "The Jira CLI interacts with a Jira instance through the command line. The app wraps Jira Rest API calls into simple commands.",
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}

var rootCmd *cobra.Command

func init() {
	rootCmd = NewCmdRoot()

	userCmd := NewCmdUser()
	rootCmd.AddCommand(userCmd)

	userCountCmd := NewCmdUserCount()
	userCmd.AddCommand(userCountCmd)

	userViewCmd := NewCmdUserView()
	userCmd.AddCommand(userViewCmd)
}

// Execute acts as the entrypoint for the command line interface.
func Execute() error {
	return rootCmd.Execute()
}
