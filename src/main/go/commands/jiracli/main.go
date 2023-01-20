package jiracli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "jiracli",
	Short: "Interact with a Jira instance through the command line",
	Long:  "Jira Ops CLI interacts with a Jira instance through the command line. The app allows some automation of recurring tasks.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use --help flag for CLI docs")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while running command '%s'", err)
		os.Exit(1)
	}
}
