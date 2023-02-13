package jiracli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"

	"github.com/sebastian-sommerfeld-io/jiracli/services/users"
)

type userViewOptions struct {
	ByEmail    bool
	ByUsername bool
	Search     string
}

// NewCmdUserView initializes the `user view` command and its flags.
func NewCmdUserView() *cobra.Command {
	opts := &userViewOptions{}

	cmd := &cobra.Command{
		Use:   "view",
		Short: "Show Username, Email and other information about a single user",
		Long:  "Show Username, Email and other information about a single user. Find user by passing a username.",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			opts.Search = args[0]
			user, err := viewUser(opts)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(user)
		},
	}

	cmd.Flags().BoolVarP(&opts.ByEmail, "email", "e", false, "Find user by email")
	cmd.Flags().BoolVarP(&opts.ByUsername, "username", "u", false, "Find user by username (default)")
	cmd.MarkFlagsMutuallyExclusive("email", "username")

	return cmd
}

func viewUser(opts *userViewOptions) (users.User, error) {
	if opts.ByEmail {
		// todo validate email input
		return users.FindByEmail(opts.Search)
	} else if opts.ByUsername {
		return users.FindByUsername(opts.Search)
	} else {
		return users.FindByUsername(opts.Search)
	}
}
