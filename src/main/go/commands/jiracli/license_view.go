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
		User: "admin",
		Pass: "admin",
	}

	cmd := &cobra.Command{
		Use:   "view",
		Short: "Show license information (Jira Software, not plugins)",
		Long:  "Show license information as JSON (Jira Software, not plugins) .",
		Args:  cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			opts.BaseURL = GetBaseUrl(cmd)
			opts.User = GetUsername(cmd)
			opts.Pass = GetPassword(cmd)

			jiraLicense, err := getJiraSoftwareLicense(opts)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(jiraLicense.RawJson)
		},
	}

	return cmd
}

func getJiraSoftwareLicense(opts *licenseViewOptions) (license.JiraLicense, error) {
	return license.ReadJiraLicense(opts.BaseURL, opts.User, opts.Pass)
}
