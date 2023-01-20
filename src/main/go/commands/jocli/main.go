package jocli

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "jocli",
	Short: "jocli - a simple CLI to interact with a Jira instance",
	Long: "jocli is a simple CLI to interact with a Jira instance. It encapsulates the Jira Rest API and provides easy access to frequently used tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print help texts")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while running command '%s'", err)
		os.Exit(1)
	}
}
