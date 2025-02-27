package logic

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func RestartConfig() {
	fmt.Println("Restarting the wglctl configuration...")

	// Ensure the config file path is valid
	configFile := viper.ConfigFileUsed()
	if configFile == "" {
		fmt.Println("No configuration file found to reset.")
		return
	}

	// Reset Viper's in-memory configuration
	viper.Reset()

	// Write an empty map to the configuration file
	err := os.WriteFile(configFile, []byte{}, 0644)
	if err != nil {
		fmt.Printf("Failed to reset configuration: %v\n", err)
		return
	}

	fmt.Println("Configuration reset successfully.")
}

func PrintConfig() {
	fmt.Println("Current wglctl Configuration:")
	for _, key := range viper.AllKeys() {
		fmt.Printf("%s: %v\n", key, viper.Get(key))
	}
}
