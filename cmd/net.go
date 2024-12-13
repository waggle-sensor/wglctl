package cmd

import (
	"fmt"
	"strings"
	"github.com/waggle-sensor/wglctl/logic"
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
	Use:   "portal <somenode> <up|down> [port]",
	Short: "Use to access switch portal.",
	Long:  `portal is used to access the node's network switch portal.
	
	Arguments:
	  <somenode>  The vsn of the node (e.g., "W030").
	  <up|down>   The action to perform (either "up" to start the tunnel or "down" to stop it).
	  [port]      The local port to use for the tunnel (optional, default is 10000).`,
	Example: `portal W030 up 8082, portal W030 down`,
	ValidArgs: []string{"up","down"},
	Args:  cobra.MinimumNArgs(2), // Require at least 2 arguments
	Run: func(cmd *cobra.Command, args []string) {
		// Extract arguments
		node := strings.ToUpper(args[0])
		action := args[1]
		localPort := "10000" // Default port

		if len(args) >= 3 {
			localPort = args[2]
		}

		switch action {
		case "up":
			logic.StartPortal(node, localPort, "net.switch.portforwading", "switch", "443")
		case "down":
			logic.StopTunnel(node, "net.switch.portforwading")
		default:
			fmt.Println("Invalid action:", action)
			fmt.Println("Usage: portal <somenode> <up|down> [port]")
		}
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

