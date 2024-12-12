package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// netCmd represents the net command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "A brief description of the net command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net called")
	},
}

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "A brief description of the switch command",
	Long:  "A longer description of the switch command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switch called")
	},
}

func init() {
	// Add the net command to the root
	rootCmd.AddCommand(netCmd)

	// Add the switch command as a subcommand of net
	netCmd.AddCommand(switchCmd)

	// Add the portal command as a subcommand of switch
	switchCmd.AddCommand(portalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

