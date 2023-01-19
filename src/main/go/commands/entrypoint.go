/*
 * Root-command. Entrypoint into the CLI app.
 */
package commands

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "cli",
	Short: "cli - a simple CLI to interact with a Jira instance",
	Long: "cli is a simple CLI to interact with a Jira instance. It encapsulates the Jira Rest API and provides easy access to frequently used tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("I am a root command ... now delegate actual todos to some business logic or print some help texts")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while running command '%s'", err)
		os.Exit(1)
	}
}
   
