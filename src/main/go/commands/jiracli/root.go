package jiracli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "jiracli",
		Short: "A command line interface to access Jira and automate recurring tasks.",
		Long:  "The Jira CLI interacts with a Jira instance through the command line. The app wraps Jira Rest API calls into simple commands.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
