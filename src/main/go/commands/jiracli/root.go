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

// FlagBaseUrl contains the name of a mandatory flag
const FlagBaseUrl string = "baseUrl"

// FlagUser contains the name of a mandatory flag
const FlagUser string = "user"

// FlagPass contains the name of a mandatory flag
const FlagPass string = "pass"

var rootCmd *cobra.Command

func init() {
	rootCmd = NewCmdRoot()
	addDefaultFlags(rootCmd)

	licenseCmd := NewCmdLicense()
	addDefaultFlags(licenseCmd)
	rootCmd.AddCommand(licenseCmd)

	licenseViewCmd := NewCmdLicenseView()
	addDefaultFlags(licenseViewCmd)
	licenseCmd.AddCommand(licenseViewCmd)

	userCmd := NewCmdUser()
	addDefaultFlags(userCmd)
	rootCmd.AddCommand(userCmd)

	userCountCmd := NewCmdUserCount()
	addDefaultFlags(userCountCmd)
	userCmd.AddCommand(userCountCmd)

	userViewCmd := NewCmdUserView()
	addDefaultFlags(userViewCmd)
	userCmd.AddCommand(userViewCmd)

	versionCmd := NewCmdVersion()
	rootCmd.AddCommand(versionCmd)
}

func addDefaultFlags(cmd *cobra.Command) {
	cmd.Flags().String(FlagBaseUrl, "", "Base URL for a Jira instance (e.g. http://localhost:8080)")
	cmd.Flags().String(FlagUser, "", "Jira user used to consume the Rest API")
	cmd.Flags().String(FlagPass, "", "Password for the Jira user")
	cmd.MarkFlagRequired(FlagBaseUrl)
	cmd.MarkFlagRequired(FlagUser)
	cmd.MarkFlagRequired(FlagPass)
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
