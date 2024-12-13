package cmd

import (
	"fmt"
	"github.com/waggle-sensor/wglctl/logic"
	"github.com/spf13/cobra"
)

// manifestCmd represents the manifest command
var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "Use to control node's manifest.",
	Long: "manifest is used to control the nodes' manifest in your waggle deployment. A manifest is the metadata associated to a node.",
}

// vsnListCmd represents the vsnlist command
var vsnListCmd = &cobra.Command{ //TODO: add filtering
	Use:   "vsnlist",
	Short: "Use to list all node's vsn.",
	Long:  `vsnlist is used to retrieve a list of node's vsn, a unique ID a node is identified by.`,
	Example: `portal W030 up 8082, portal W030 down`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		vsns, err := logic.FetchVSNs("https://auth.sagecontinuum.org/manifests/")
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
		fmt.Printf("%s\n", vsns)
	},
}

func init() {
	// Add the net command to the root
	rootCmd.AddCommand(manifestCmd)

	manifestCmd.AddCommand(vsnListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}