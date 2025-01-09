package cmd

import (
	"strings"
	"github.com/waggle-sensor/wglctl/logic"
	"github.com/spf13/cobra"
)

// netCmd represents the net command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Use to control network.",
	Long: `net is used to control network in your waggle deployment.`,
}

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Use to control the network switch.",
	Long:  "switch is used to control the network switch in your waggle deployment.",
}

// switchPortalCmd represents the portal command
var switchPortalCmd = &cobra.Command{
	Use:   "portal",
	Short: "Use to control switch portal(s).",
	Long:  "portal is used to control the node's network switch portal(s).",
}

// switchPortalupCmd represents the portal up command
var switchPortalupCmd = &cobra.Command{
	Use:   "up <somenode> [port]",
	Short: "Use to access switch portal.",
	Long:  `up is used to access the node's network switch portal.
	
	Arguments:
	  <somenode>  The vsn of the node (e.g., "W030").
	  [port]      The local port to use for the tunnel (optional, default is 10000).`,
	Example: `portal up W030, portal up W030 8082`,
	Args:  cobra.MinimumNArgs(1), // Require at least 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		// Extract arguments
		node := strings.ToUpper(args[0])
		localPort := "10000" // Default port

		if len(args) >= 2 {
			localPort = args[1]
		}

		logic.StartPortal(node, localPort, "net.switch.portforwading", "switch", "https", "443")
	},
}

// switchPortaldownCmd represents the portal down command
var switchPortaldownCmd = &cobra.Command{
	Use:   "down <somenode/all>",
	Short: "Use to terminate switch portal.",
	Long:  `down is used to terminate the node's network switch portal or all active switch portals.
	
	Arguments:
	  <somenode/all>  The vsn of the node (e.g., "W030") or all.`,
	Example: `portal down W030`,
	Args:  cobra.ExactArgs(1), // Require exactly 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		// Extract arguments
		arg := strings.ToUpper(args[0])
		config := "net.switch.portforwading"

		if arg == "ALL" {
			logic.StopAll(config)
		} else {
			logic.StopTunnel(arg, config)
		}
	},
}

// switchPortaldownCmd represents the portal ls command
var switchPortalListCmd = &cobra.Command{
	Use:   "ls [somenode]",
	Short: "Use to list switch portal(s).",
	Long:  `ls is used to list active network switch portal(s).
	
	Arguments:
	  [somenode]  The vsn of the node (e.g., "W030"). optional, default is all.`,
	Example: `portal ls, portal ls W030`,
	Args:  cobra.MaximumNArgs(1), // Require no greater than 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		// Extract arguments
		node := "" // Default node

		if len(args) >= 1 {
			node = strings.ToUpper(args[0])
		}

		logic.ListTunnel(node, "net.switch.portforwading")
	},
}

func init() {
	// Add the net command to the root
	rootCmd.AddCommand(netCmd)

	// Add the switch command as a subcommand of net
	netCmd.AddCommand(switchCmd)

	// Add the portal command as a subcommand of switch
	switchCmd.AddCommand(switchPortalCmd)

	// Add the actions for the portal command
	switchPortalCmd.AddCommand(switchPortalupCmd)
	switchPortalCmd.AddCommand(switchPortaldownCmd)
	switchPortalCmd.AddCommand(switchPortalListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

