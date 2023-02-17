package commands

import (
	"fmt"
	"log"
	"net/mail"

	"github.com/spf13/cobra"

	"github.com/sebastian-sommerfeld-io/jiracli/services/jira/users"
)

type userViewOptions struct {
	ByEmail    bool
	ByUsername bool
	Search     string
}

const (
	// FlagByEmail contains the name of the flag
	FlagByEmail string = "by-email"

	// FlagByUsername contains the name of the flag
	FlagByUsername string = "by-username"
)

// NewCmdUserView initializes the `jiracli user view` command and its flags.
func NewCmdUserView() *cobra.Command {
	opts := &userViewOptions{}

	cmd := &cobra.Command{
		Use:   "view",
		Short: "Show Username, Email and other information about a single user",
		Long:  "Show Username, Email and other information about a single user. Find user by passing a username.",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			opts.Search = args[0]
			user, err := getUser(opts)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(user)
		},
	}

	AddMandatoryFlags(cmd)

	cmd.Flags().BoolVarP(&opts.ByEmail, FlagByEmail, "", false, "Find user by email")
	cmd.Flags().BoolVarP(&opts.ByUsername, FlagByUsername, "", false, "Find user by username (default)")
	cmd.MarkFlagsMutuallyExclusive(FlagByEmail, FlagByUsername)

	return cmd
}

func getUser(opts *userViewOptions) (users.User, error) {
	if opts.ByEmail {
		valid, err := validateEmail(opts.Search)
		if valid {
			return users.FindByEmail(opts.Search)
		} else {
			user := users.User{}
			return user, err
		}
	} else if opts.ByUsername {
		return users.FindByUsername(opts.Search)
	} else {
		return users.FindByUsername(opts.Search)
	}
}

func validateEmail(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	return err == nil, err
}
