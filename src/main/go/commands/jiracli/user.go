package jiracli

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	userCmd = &cobra.Command{
		Use:   "user",
		Short: "Work with Jira users.",
		Args:  cobra.ExactArgs(1),
	}
)

func init() {
	log.SetPrefix("commands/user - ")

	rootCmd.AddCommand(userCmd)
}
