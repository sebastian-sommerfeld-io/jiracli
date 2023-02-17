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

const (
	// FlagBaseUrl contains the name of a mandatory flag
	FlagBaseUrl string = "baseUrl"

	// FlagUser contains the name of a mandatory flag
	FlagUser string = "user"

	// FlagPass contains the name of a mandatory flag
	FlagPass string = "pass"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = NewCmdRoot()
	addDefaultFlags(rootCmd)

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

func addDefaultFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(FlagBaseUrl, "", "Base URL for a Jira instance (e.g. http://localhost:8080)")
	cmd.PersistentFlags().String(FlagUser, "", "Jira user used to consume the Rest API")
	cmd.Flags().String(FlagPass, "", "Password for the Jira user")
	cmd.MarkPersistentFlagRequired(FlagBaseUrl)
	cmd.MarkPersistentFlagRequired(FlagUser)
	cmd.MarkPersistentFlagRequired(FlagPass)
}

// Execute acts as the entrypoint for the command line interface.
func Execute() error {
	return rootCmd.Execute()
}

// GetBaseUrl returns the value for the `--baseUrl` flag as string
func GetBaseUrl(cmd *cobra.Command) string {
	return getFlagValue(cmd, FlagBaseUrl)
}

// GetUsername returns the value for the `--user` flag as string
func GetUsername(cmd *cobra.Command) string {
	return getFlagValue(cmd, FlagUser)
}

// GetPassword returns the value for the `--pass` flag as string
func GetPassword(cmd *cobra.Command) string {
	return getFlagValue(cmd, FlagPass)
}

func getFlagValue(cmd *cobra.Command, flag string) string {
	value, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
