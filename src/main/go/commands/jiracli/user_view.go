package jiracli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"

	"github.com/sebastian-sommerfeld-io/jiracli/services/user"
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
	log.SetPrefix("user_commands - ")

	userCmd.AddCommand(userViewCmd)
}

func viewUser(needle string) user.User {
	user, e := user.FindByUsername(needle)

	if e != nil {
		log.Fatal(e)
	}

	return user
}
