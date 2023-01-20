package jocli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"sebastian-sommerfeld-io/jira-ops-cli/services/user"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Work with Jira users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo ...")
	},
}

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo ...")
	},
}

var userCountCmd = &cobra.Command{
	Use:   "count",
	Short: "Count all users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo ...")
	},
}

var userViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display Name, Email and other information about a user. Pass username as string.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viewUser(args[0]))
	},
}

func init() {
	log.SetPrefix("user_commands - ")

	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userListCmd)
	userCmd.AddCommand(userCountCmd)
	userCmd.AddCommand(userViewCmd)
}

func viewUser(needle string) user.User {
	user, e := user.FindByUsername(needle)

	if e != nil {
		log.Fatal(e)
	}

	return user
}
