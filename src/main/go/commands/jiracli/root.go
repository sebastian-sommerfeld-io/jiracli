package jiracli

import (
	"log"

	"github.com/spf13/cobra"
)

// NewCmdRoot initializes the `jiracli` root command.
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jiracli",
		Short: "A command line interface to access Jira and automate recurring tasks",
		Long:  "The Jira CLI interacts with a Jira instance through the command line. The app wraps Jira Rest API calls into simple commands.",
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}

const FLAG_BASEURL string = "baseUrl"
const FLAG_USER string = "user"
const FLAG_PASS string = "pass"

var rootCmd *cobra.Command

func init() {
	rootCmd = NewCmdRoot()
	rootCmd.PersistentFlags().String(FLAG_BASEURL, "", "Base URL for a Jira instance (e.g. http://localhost:8080)")
	rootCmd.PersistentFlags().String(FLAG_USER, "", "Jira user used to consume the Rest API")
	rootCmd.PersistentFlags().String(FLAG_PASS, "", "Password for the Jira user")
	rootCmd.MarkPersistentFlagRequired(FLAG_BASEURL)
	rootCmd.MarkPersistentFlagRequired(FLAG_USER)
	rootCmd.MarkPersistentFlagRequired(FLAG_PASS)

	licenseCmd := NewCmdLicense()
	rootCmd.AddCommand(licenseCmd)

	licenseViewCmd := NewCmdLicenseView()
	licenseCmd.AddCommand(licenseViewCmd)

	userCmd := NewCmdUser()
	rootCmd.AddCommand(userCmd)

	userCountCmd := NewCmdUserCount()
	userCmd.AddCommand(userCountCmd)

	userViewCmd := NewCmdUserView()
	userCmd.AddCommand(userViewCmd)
}

// Execute acts as the entrypoint for the command line interface.
func Execute() error {
	return rootCmd.Execute()
}

// GetBaseUrl returns the value for the `--baseUrl` flag as string
func GetBaseUrl(cmd *cobra.Command) string {
	return getFlagValue(cmd, FLAG_BASEURL)
}

// GetUser returns the value for the `--user` flag as string
func GetUsername(cmd *cobra.Command) string {
	return getFlagValue(cmd, FLAG_USER)
}

// GetPassword returns the value for the `--pass` flag as string
func GetPassword(cmd *cobra.Command) string {
	return getFlagValue(cmd, FLAG_PASS)
}

func getFlagValue(cmd *cobra.Command, flag string) string {
	value, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
