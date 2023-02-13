package jiracli

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"

	"github.com/sebastian-sommerfeld-io/jiracli/services/users"
)

type ViewOptions struct {
	ByEmail    bool
	ByUsername bool
	Search     string
}

// NewCmdUserView initializes the `user view` command and its flags.
func NewCmdUserView() *cobra.Command {
	opts := &ViewOptions{}

	cmd := &cobra.Command{
		Use:   "view",
		Short: "Show Username, Email and other information about a single user",
		Long:  "Show Username, Email and other information about a single user. Find user by passing a username.",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			opts.Search = args[0]
			fmt.Println(viewUser(opts))
		},
	}

	cmd.Flags().BoolVarP(&opts.ByEmail, "email", "e", false, "Find user by email")
	cmd.Flags().BoolVarP(&opts.ByUsername, "username", "u", false, "Find user by username (default)")
	cmd.MarkFlagsMutuallyExclusive("email", "username")

	return cmd
}

func viewUser(opts *ViewOptions) users.User {
	if opts.ByEmail {
		fmt.Println("----> e")
	}

	user, err := users.FindByUsername(opts.Search)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
