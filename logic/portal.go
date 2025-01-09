package logic

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/spf13/viper"
)

// StartPortal starts the port forwarding for a given node's portalIp:portalPort.
// 	node: The vsn of the node.
//	localPort: Localhost port to forward to.
//  configObject: Object to use in the config file for persistant storage of active portals.
//  portalIp: The ip of the portal in the Node.
//  portalPort: The port of the portal in the Node.
func StartPortal(node, localPort, configObject, portalIp , protocol, portalPort string) {
    fmt.Printf("Starting tunnel to %s for portal access on port %s...\n", node, localPort)

	// Retrieve tunnel information from the config file
	tunnels := viper.GetStringMap(configObject)
	tunnelInfo, exists := tunnels[strings.ToLower(node)]
	if exists {
		// Extract the port configuration
		tunnel, ok := tunnelInfo.(map[string]interface{})
		if !ok {
			fmt.Printf("Invalid port forwading configuration for node %s.\n", node)
			fmt.Printf("Fix the config file.\n")
			return
		}
		port, _ := tunnel["localport"].(string)
		fmt.Printf("A tunnel is already active for node %s on port %s.\n", node, port)
		fmt.Printf("Visit portal at %s://localhost:%s\n", protocol, port)
		return
	}

    // Check if the port is already being used by an existing SSH proxy
    cmd := exec.Command("lsof", "-i", fmt.Sprintf(":%s", localPort))
    lsofOutput, err := cmd.Output()
	if err == nil && len(string(lsofOutput)) > 0 {
        fmt.Printf("Error: Port %s is already in use.\n", localPort)
		fmt.Printf("Processes using Port %s:\n", localPort)
		fmt.Printf("%s", string(lsofOutput))
        return
	} 

    // Create the SSH proxy
	sshArg := fmt.Sprintf("%s:%s:%s", localPort, strings.TrimSpace(string(portalIp)), portalPort)
    proxyCmd := exec.Command("ssh", "-N", "-L", sshArg, fmt.Sprintf("node-%s", node))
    if err := proxyCmd.Start(); err != nil {
        fmt.Printf("Error: Failed to establish tunnel to %s: %v\n", node, err)
        return
    }

    // Save tunnel information in configuration
    tunnels[node] = map[string]string{
        "localport": localPort,
        "svcip":   strings.TrimSpace(string(portalIp)),
		"svcport": portalPort,
    }
    viper.Set(configObject, tunnels)

    // Write back to the config file
    if err := viper.WriteConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            // If the config file doesn't exist, create it
            if err := viper.SafeWriteConfig(); err != nil {
                fmt.Printf("Error creating configuration file: %v\n", err)
            }
        } else {
            fmt.Printf("Error writing configuration file: %v\n", err)
        }
    }

    fmt.Printf("Tunnel to %s established.\n", node)
    fmt.Printf("Visit portal at %s://localhost:%s\n", protocol, localPort)
}

// StopAll stops port forwarding for all nodes.
// configObject: Object to use in the config file for finding active portals.
func StopAll(configObject string) {
	fmt.Println("Stopping tunnels for all nodes...")

	// Retrieve tunnel information from the config file
	tunnels := viper.GetStringMap(configObject)

	// Check if there are any active tunnels
	if len(tunnels) == 0 {
		fmt.Println("No active tunnels found.")
		return
	}

	// Iterate through each node and stop its tunnel
	for node, tunnelInfo := range tunnels {
		stopProcess(node, tunnelInfo)
		// Remove the node from the configuration
		delete(tunnels, strings.ToLower(node))
	}

	// Update the configuration file
	viper.Set(configObject, tunnels)
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Error updating configuration file: %v\n", err)
	} 

	fmt.Println("All active tunnels have been stopped.")
}

// StopTunnel stops the port forwarding for a specific node.
// node: The vsn of the node.
// configObject: Object to use in the config file for finding active portals.
func StopTunnel(node, configObject string) {
	fmt.Printf("Stopping tunnel for node %s...\n", node)

	// Retrieve tunnel information from the config file
	tunnels := viper.GetStringMap(configObject)
	tunnelInfo, exists := tunnels[strings.ToLower(node)]
	if !exists {
		fmt.Printf("No active tunnels found for node %s.\n", node)
		return
	}

	// Stop the process for the specified node
	stopProcess(node, tunnelInfo)

	// Remove the node from the configuration file
	delete(tunnels, strings.ToLower(node))
	viper.Set(configObject, tunnels)
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Error updating configuration file: %v\n", err)
	} else {
		fmt.Println("Configuration file updated.")
	}
}

// ListTunnel lists the current portals that are up.
//  node: The vsn of the node (optional).
//  configObject: Object to use in the config file for finding active portals.
func ListTunnel(node, configObject string) {
    // Retrieve tunnel information from the config file
    tunnels := viper.GetStringMap(configObject)

    if node == "" {
        // If no node is specified, list all tunnels
        if len(tunnels) == 0 {
            fmt.Println("No active tunnels found.")
            return
        }
        fmt.Println("Active Portals:")
        for node, tunnelInfo := range tunnels {
			formatPortal(node,tunnelInfo)
        }
    } else {
        // If a node is specified, list its tunnel information
        tunnelInfo, exists := tunnels[strings.ToLower(node)]
        if !exists {
            fmt.Printf("No active tunnels found for node %s.\n", node)
            return
        }
		fmt.Println("Active Portals:")
		formatPortal(node,tunnelInfo)
    }
}

// formatPortal formats active portal information for terminal output.
//  node: The vsn of the node.
//  tunnelInfo: tunnel information retrieved from viper.
func formatPortal(node string, tunnelInfo interface{}) {
	fmt.Printf("- Node: %s\n", strings.ToUpper(node))
	if tunnelMap, ok := tunnelInfo.(map[string]interface{}); ok {
		for key, value := range tunnelMap {
			fmt.Printf("   - %s:%v\n", key, value)
		}
	} else {
		fmt.Printf("   - Info: %v\n", tunnelInfo)
	}
}

// stopProcess stops the process associated with a tunnel.
//  node: The vsn of the node.
//  tunnelInfo: tunnel information retrieved from viper.
func stopProcess(node string, tunnelInfo interface{}) {
	//all caps the node
	node = strings.ToUpper(node)
	
	// Extract the port and IP from the configuration
	tunnel, ok := tunnelInfo.(map[string]interface{})
	if !ok {
		fmt.Printf("Invalid port forwarding configuration for node %s.\n", node)
		fmt.Printf("Fix the config file %s.\n", viper.ConfigFileUsed())
		return
	}

	localPort, localPortExists := tunnel["localport"].(string)
	ip, ipExists := tunnel["svcip"].(string)
	svcPort, svcPortExists := tunnel["svcport"].(string)

	if !localPortExists || !ipExists || !svcPortExists {
		fmt.Printf("Missing configuration for node %s.\n", node)
		return
	}

	// Construct the pattern to match the exact command
	pattern := fmt.Sprintf("ssh -N -L %s:%s:%s node-%s", localPort, ip, svcPort, node)

	// Use pgrep to find the process matching the pattern
	pgrepCmd := exec.Command("pgrep", "-f", pattern)
	pids, err := pgrepCmd.Output()

	if err != nil || len(pids) == 0 {
		fmt.Printf("No active SSH proxy found for node %s.\n", node)
		return
	}

	// Kill the processes associated with the tunnel
	pidList := strings.Fields(string(pids))
	for _, pid := range pidList {
		killCmd := exec.Command("kill", pid)
		if err := killCmd.Run(); err != nil {
			fmt.Printf("Failed to stop tunnel for node %s (PID %s): %v\n", node, pid, err)
		}
	}
}