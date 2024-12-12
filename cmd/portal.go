package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
// portalCmd represents the portal command
var portalCmd = &cobra.Command{
	Use:   "portal",
	Short: "A brief description of the portal command",
	Long:  "A longer description of the portal command.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("portal called")
	},
}