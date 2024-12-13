package cmd

import (
	"fmt"
	"strings"
	"github.com/waggle-sensor/wglctl/logic"
	"github.com/spf13/cobra"
)

// lorawanCmd represents the lorawan command
var lorawanCmd = &cobra.Command{
	Use:   "lorawan",
	Short: "Use to control lorawan.",
	Long: "lorawan is used to control the lorawan hardware/software in your waggle deployment.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lorawan called")
	},
}

// lwPortalCmd represents the portal command
var lwPortalCmd = &cobra.Command{
	Use:   "portal <somenode> <up|down> [port]",
	Short: "Use to access ChirpStack portal.",
	Long:  `portal is used to access the node's Chirpstack network server portal.
	
	Arguments:
	  <somenode>  The vsn of the node (e.g., "W030").
	  <up|down>   The action to perform (either "up" to start the tunnel or "down" to stop it).
	  [port]      The local port to use for the tunnel (optional, default is 8081).`,
	Example: `portal W030 up 8082, portal W030 down`,
	ValidArgs: []string{"up","down"},
	Args:  cobra.MinimumNArgs(2), // Require at least 2 arguments
	Run: func(cmd *cobra.Command, args []string) {
		// Extract arguments
		node := strings.ToUpper(args[0])
		action := args[1]
		localPort := "8081" // Default port

		if len(args) >= 3 {
			localPort = args[2]
		}

		switch action {
		case "up":
			portalIp := logic.GetChirpStackIp(node)
			logic.StartPortal(node, localPort, "lora.portforwading", portalIp, "8080")
		case "down":
			logic.StopTunnel(node, "lora.portforwading")
		default:
			fmt.Println("Invalid action:", action)
			fmt.Println("Usage: portal <somenode> <up|down> [port]")
		}
	},
}

func init() {
	// Add the net command to the root
	rootCmd.AddCommand(lorawanCmd)

	lorawanCmd.AddCommand(lwPortalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
