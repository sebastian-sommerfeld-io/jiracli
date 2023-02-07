package jiracli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"

	"github.com/sebastian-sommerfeld-io/jiracli/services/users"
)

var (
	userViewCmd = &cobra.Command{
		Use:   "view",
		Short: "Show Username, Email and other information about a user.",
		Long:  "Show Username, Email and other information about a user. Find user by passing a username.",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viewUser(args[0]))
		},
	}
)

func init() {
	log.SetPrefix("commands/user/view - ")

	userCmd.AddCommand(userViewCmd)
}

func viewUser(needle string) users.User {
	user, err := users.FindByUsername(needle)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
