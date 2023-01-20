package jocli

import (
    "fmt"
    "sebastian-sommerfeld-io/jira-ops-cli/services/dummy"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
    Use:   "project",
    Aliases: []string{"p"},
    Short:  "Work with Jira projects",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(dummy.Dummy(args[0]))
    },
}

var listCmd = &cobra.Command{
    Use:   "list",
    Aliases: []string{"p"},
    Short:  "List all Jira projects",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(dummy.Dummy(args[0]))
    },
}

var countCmd = &cobra.Command{
    Use:   "list",
    Aliases: []string{"p"},
    Short:  "Count all Jira projects",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(dummy.Dummy(args[0]))
    },
}

func init() {
    rootCmd.AddCommand(projectCmd)
    projectCmd.AddCommand(listCmd)
    projectCmd.AddCommand(countCmd)
}
