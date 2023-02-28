package commands

import (
	"log"

	"github.com/sebastian-sommerfeld-io/jiracli/services"
	"github.com/spf13/cobra"
)

const (
	// FlagBaseUrl contains the name of a mandatory flag
	FlagBaseUrl string = "baseUrl"

	// FlagUser contains the name of a mandatory flag
	FlagUser string = "user"

	// FlagPass contains the name of a mandatory flag
	FlagPass string = "pass"
)

// NewCmdRoot initializes the `jiracli` root command.
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "jiracli",
		Version: appVersion(),
		Short:   "A command line interface to access Jira and automate recurring tasks",
		Long:    "The Jira CLI interacts with a Jira instance through the command line. The app wraps Jira Rest API calls into simple commands.",
		Args:    cobra.ExactArgs(0),
	}

	return cmd
}

var rootCmd *cobra.Command

func init() {
	rootCmd = NewCmdRoot()

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

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

// AddMandatoryFlags adds a set of flags to the given command. These flags are mandatory for all
// commands, that actually connect to a Jira instance.
func AddMandatoryFlags(cmd *cobra.Command) {
	cmd.Flags().String(FlagBaseUrl, "", "An URL to a Jira instance without any path information (e.g. http://localhost:8080)")
	cmd.Flags().String(FlagUser, "", "Name of the Jira user to consume the Rest API")
	cmd.Flags().String(FlagPass, "", "Password of the Jira user to consume the Rest API")
	markFlagRequired(cmd, FlagBaseUrl)
	markFlagRequired(cmd, FlagUser)
	markFlagRequired(cmd, FlagPass)
}

func markFlagRequired(cmd *cobra.Command, flagName string) {
	err := cmd.MarkFlagRequired(flagName)
	if err != nil {
		log.Fatal(err)
	}
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

func appVersion() string {
	info := services.JiracliAppInfo{}
	info, err := info.Read()

	if err != nil {
		log.Fatal(err)
	}

	return info.Version
}
