package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/sebastian-sommerfeld-io/jiracli/services/jira/license"
	"github.com/spf13/cobra"
)

type licenseViewOptions struct {
	BaseURL string
	User    string
	Pass    string
}

// NewCmdLicenseView initializes the `jiracli license` command.
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

			printJiraSoftwareLicense(opts)
		},
	}

	AddMandatoryFlags(cmd)

	return cmd
}

func printJiraSoftwareLicense(opts *licenseViewOptions) {
	jiraLicense, err := license.ReadJiraLicense(opts.BaseURL, opts.User, opts.Pass)
	if err != nil {
		log.Fatal(err)
	}

	n, err := fmt.Fprintf(os.Stdout, jiraLicense.RawJson)
	if err != nil || n <= 0 {
		log.Fatal(err)
	}
}
