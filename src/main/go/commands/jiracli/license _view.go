package jiracli

import (
	"fmt"
	"log"

	"github.com/sebastian-sommerfeld-io/jiracli/services/license"
	"github.com/spf13/cobra"
)

type licenseViewOptions struct {
	BaseURL string
	User    string
	Pass    string
}

// NewCmdUserView initializes the `jiracli license` command.
func NewCmdLicenseView() *cobra.Command {
	opts := &licenseViewOptions{
		BaseURL: "http://localhost:8080",
		User:    "admin",
		Pass:    "admin",
	}

	cmd := &cobra.Command{
		Use:   "view",
		Short: "Show license information (Jira Software, not plugins)",
		Long:  "Show license information as JSON (Jira Software, not plugins) .",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			json, err := getJiraSoftwareLicense(opts)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(json)
		},
	}

	return cmd
}

func getJiraSoftwareLicense(opts *licenseViewOptions) (string, error) {
	return license.ReadJiraSoftwareLicense(opts.BaseURL, opts.User, opts.Pass)
}
