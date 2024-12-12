package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// netCmd represents the net command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Use to control network.",
	Long: `net is used to control network in your waggle deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net called")
	},
}

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Use to control the network switch.",
	Long:  "switch is used to control the network switch in your waggle deployment.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switch called")
	},
}

// switchPortalCmd represents the portal command
var switchPortalCmd = &cobra.Command{
	Use:   "portal",
	Short: "Use to access switch portal.",
	Long:  "portal is used to access the node's network switch portal.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("portal called")
	},
}

func init() {
	// Add the net command to the root
	rootCmd.AddCommand(netCmd)

	// Add the switch command as a subcommand of net
	netCmd.AddCommand(switchCmd)

	// Add the portal command as a subcommand of switch
	switchCmd.AddCommand(switchPortalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

