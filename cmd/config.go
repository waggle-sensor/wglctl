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

func init() {

	// Add the config command to the root
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(restartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

