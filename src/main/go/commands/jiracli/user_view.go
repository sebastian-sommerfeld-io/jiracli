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
		Short: "Show Username, Email and other information about a single user",
		Long:  "Show Username, Email and other information about a single user. Find user by passing a username.",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viewUser(args[0]))
		},
	}
)

func init() {
	userCmd.AddCommand(userViewCmd)
}

func viewUser(search string) users.User {
	user, err := users.FindByUsername(search)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
