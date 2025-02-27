package logic

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func OpenDashboard(vsn string) {
	// Get Grafana URL and lorawan dashboard ID from the configuration
	grafanaBaseURL := viper.GetString("GRAFANA_BASE_URL")
	dashboard_id := viper.GetString("GRAFANA_LW_DASHBOARD_ID")

	// Construct final Grafana URL with the node variable set
	url := grafanaBaseURL
	if vsn != "" {
		url += "/d/" + dashboard_id + "?var-vsn=" + vsn
	}

	fmt.Printf("Opening Grafana dashboard for node: %s\n", vsn)

	// Open the URL in the default web browser
	err := openBrowser(url)
	if err != nil {
		fmt.Println("Failed to open browser:", err)
	}
}

// SetGrafanaURL sets and stores the Grafana base URL in the config file
func SetGrafanaURL(url string) {
	// Validate URL format 
	if len(url) < 10 || !(strings.HasPrefix(url, "http")) {
		fmt.Println("Invalid URL. Please provide a valid Grafana URL, e.g., http://your-grafana-url")
		os.Exit(1)
	}

	// Remove any proceeding '/' from the URL
	url = strings.TrimRight(url, "/")

	// Store the cleaned URL in Viper
	viper.Set("GRAFANA_BASE_URL", url)

	// Save the configuration persistently
	err := viper.WriteConfig()
	if err != nil {
		fmt.Printf("Error saving config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Grafana Base URL updated successfully:", url)
}

// SetLorawanDashboardID sets the ID of the LoRaWAN dashboard in the configuration.
func SetLorawanDashboardID(dashboardID string) {
    // Store the ID in Viper
    viper.Set("GRAFANA_LW_DASHBOARD_ID", dashboardID)

    // Save the configuration persistently
    err := viper.WriteConfig()
    if err != nil {
        fmt.Printf("Error saving config: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("LoRaWAN Dashboard ID updated successfully:", dashboardID)
}