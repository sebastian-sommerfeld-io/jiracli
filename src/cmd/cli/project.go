/*
 * Sub-command. Handle everything related to projects.
 */
package cli

import (
    "sebastian-sommerfeld-io/jira-ops-cli/services/dummy"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
    Use:   "project",
    Aliases: []string{"p"},
    Short:  "Handle Jira projects",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        dummy.Echo(args[0])
    },
}

func init() {
    rootCmd.AddCommand(projectCmd)
}
