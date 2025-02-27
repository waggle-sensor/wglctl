package cmd

import (
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wglctl",
	Short: "A cli for controlling waggle.",
	Long: `wglctl is CLI for controlling Waggle: An Edge 
	Computing Platform for Artificial Intelligence and Sensing.
	Any user from an Admin to a data consumer can use this CLI
	to better manage their waggle deployment.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Global config flag for specifying a custom config file
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/wglctl/config.yaml)")

	// Grafana Base URL flag
	// This allows users to override the variables using a CLI flag
	rootCmd.PersistentFlags().String("grafana-base-url", "", "Base URL for Grafana dashboards")
	viper.BindPFlag("GRAFANA_BASE_URL", rootCmd.PersistentFlags().Lookup("grafana-base-url"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Define the configuration directory.
		configDir := filepath.Join(home, ".config", "wglctl")

		// Ensure the configuration directory exists.
		err = os.MkdirAll(configDir, 0755)
		cobra.CheckErr(err)

		// Search config in ~/.config/wglctl directory with name "config" (without extension).
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

		// set the config file path
		cfgFile = filepath.Join(configDir, "config.yaml")
	}

	// Read environment variables automatically
	viper.AutomaticEnv()

	// Try to read the config file, but don't fail if it doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("No existing config file found. Creating a new one: %s\n", cfgFile)
	}

	// Set default values only if they are missing
	setDefaultConfig("GRAFANA_BASE_URL", "http://localhost:3000")
	setDefaultConfig("GRAFANA_LW_DASHBOARD_ID", "1")

	// Save the config if any defaults were added
	saveConfig()
}

// setDefaultConfig sets a default value only if the key is missing
func setDefaultConfig(key string, value string) {
	if !viper.IsSet(key) {
		viper.Set(key, value)
	}
}

// saveConfig writes the current viper config to the config file
func saveConfig() {
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Error updating configuration file: %v\n", err)
	}
}
