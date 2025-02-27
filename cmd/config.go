package cmd

import (
	"github.com/waggle-sensor/wglctl/logic"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Use to control wglctl config.",
	Long: "config is used to control the config file of wglctl.",
}

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:     "restart",
	Short:   "Restart the wglctl configuration file. Use this command carefully as it may cause wglctl to stop working.",
	Long:    `Restart is used to clear and reset the wglctl configuration file. Use this command carefully as it may cause wglctl to stop working.`,
	Example: "wglctl config restart",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		logic.RestartConfig()
	},
}

// setGrafanaURLCmd represents the command to set the Grafana Base URL
var setGrafanaURLCmd = &cobra.Command{
	Use:   "set-grafana-url <url>",
	Short: "Set the base URL for commands that use Grafana.",
	Long: `This command sets the base URL for Grafana. 
	It saves the URL in the configuration so all Grafana-related commands use it.`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument (the URL)
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0] // Get the user-provided URL

		logic.SetGrafanaURL(url)
	},
}

// printConfigCmd represents the command to print the current configuration
var printConfigCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the current wglctl configuration.",
	Long:  `This command prints all the current configuration settings used by wglctl.`,
	Example: `wglctl config print`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		logic.PrintConfig()
	},
}

func init() {

	// Add the config command to the root
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(restartCmd)
	configCmd.AddCommand(setGrafanaURLCmd)
	configCmd.AddCommand(printConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

