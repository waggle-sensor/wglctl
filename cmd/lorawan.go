package cmd

import (
	"strings"
	"github.com/waggle-sensor/wglctl/logic"
	"github.com/spf13/cobra"
)

// lorawanCmd represents the lorawan command
var lorawanCmd = &cobra.Command{
	Use:   "lorawan",
	Short: "Use to control lorawan.",
	Long: "lorawan is used to control the lorawan hardware/software in your waggle deployment.",
}

// lwPortalCmd represents the portal command
var lwPortalCmd = &cobra.Command{
	Use:   "portal",
	Short: "Use to control ChirpStack portal(s).",
	Long:  "portal is used to control the node's Chirpstack network server portal(s).",
}

// lwPortalupCmd represents the portal up command
var lwPortalupCmd = &cobra.Command{
	Use:   "up <somenode> [port]",
	Short: "Use to access ChirpStack portal.",
	Long:  `up is used to access the node's Chirpstack network server portal.
	
	Arguments:
	  <somenode>  The vsn of the node (e.g., "W030").
	  [port]      The local port to use for the tunnel (optional, default is 8081).`,
	Example: `portal up W030, portal up W030 8082`,
	Args:  cobra.MinimumNArgs(1), // Require at least 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		// Extract arguments
		node := strings.ToUpper(args[0])
		localPort := "8081" // Default port

		if len(args) >= 2 {
			localPort = args[1]
		}

		portalIp := logic.GetChirpStackIp(node)
		logic.StartPortal(node, localPort, "lora.portforwading", portalIp, "http", "8080")
	},
}

// lwPortaldownCmd represents the portal down command
var lwPortaldownCmd = &cobra.Command{
	Use:   "down <somenode/all>",
	Short: "Use to terminate Chirpstack portal.",
	Long:  `down is used to terminate the node's Chirpstack portal or all active Chirpstack portals.
	
	Arguments:
	  <somenode/all>  The vsn of the node (e.g., "W030") or all.`,
	Example: `portal down W030`,
	Args:  cobra.ExactArgs(1), // Require exactly 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		// Extract arguments
		arg := strings.ToUpper(args[0])
		config := "lora.portforwading"

		if arg == "ALL" {
			logic.StopAll(config)
		} else {
			logic.StopTunnel(arg, config)
		}
	},
}

// lwPortalListCmd represents the portal ls command
var lwPortalListCmd = &cobra.Command{
	Use:   "ls [somenode]",
	Short: "Use to list portal(s).",
	Long:  `ls is used to list active Chirpstack portal(s).
	
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

		logic.ListTunnel(node, "lora.portforwading")
	},
}

// setLorawanDashboardIDCmd represents the command to set the LoRaWAN Dashboard ID
var setLorawanDashboardIDCmd = &cobra.Command{
	Use:   "set-lorawan-dashboard-id <dashboard_id>",
	Short: "Set the LoRaWAN dashboard ID for Grafana",
	Long: `This command sets the LoRaWAN dashboard ID for Grafana. 
	It saves the ID in the configuration so all lorawan dashboard-related commands use it.`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument (the ID)
	Run: func(cmd *cobra.Command, args []string) {
		dashboardID := args[0]

		logic.SetLorawanDashboardID(dashboardID)
	},
}

// lwReportCmd represents the report command
var lwReportCmd = &cobra.Command{
	Use:   "dashboard <somenode>",
	Short: "Use to access a lorawan dashboard of the node.",
	Long:  `dashboard is used to access a lorawan Grafana dashboard of the node. If no node is provided, 
	the default will just open the Grafana dashboard with the default node selected.
	
	Arguments:
	  <somenode>  The vsn of the node (e.g., "W030"). optional, default is None.`,
	Example: `dashboard, dashboard W030`,
	Args:  cobra.MaximumNArgs(1), // Require no greater than 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		// Extract arguments
		node := "" // Default node

		if len(args) >= 1 {
			node = strings.ToUpper(args[0])
		}

		logic.OpenDashboard(node)
	},
}

func init() {
	// Add the lorawan command to the root
	rootCmd.AddCommand(lorawanCmd)

	// Add the set-lorawan-dashboard-id command as a subcommand of lorawan
	lorawanCmd.AddCommand(setLorawanDashboardIDCmd)

	// Add the report command as a subcommand of lorawan
	lorawanCmd.AddCommand(lwReportCmd)

	// Add the portal command as a subcommand of lorawan
	lorawanCmd.AddCommand(lwPortalCmd)

	// Add the actions for the portal command
	lwPortalCmd.AddCommand(lwPortalupCmd)
	lwPortalCmd.AddCommand(lwPortaldownCmd)
	lwPortalCmd.AddCommand(lwPortalListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
