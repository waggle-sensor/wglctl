package cmd

import (
	"fmt"
	"strings"
	"github.com/waggle-sensor/wglctl/logic"
	"github.com/spf13/cobra"
)

// netCmd represents the net command
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
	Long:  "portal is used to access the node's Chirpstack network server portal.",
	Args:  cobra.MinimumNArgs(2), // Require at least 2 arguments
	Run: func(cmd *cobra.Command, args []string) {
		// Extract arguments
		node := strings.ToUpper(args[0])
		action := args[1]
		port := "8081" // Default port

		if len(args) >= 3 {
			port = args[2]
		}

		switch action {
		case "up":
			logic.StartTunnel(node, port)
		case "down":
			logic.StopTunnel(node)
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
