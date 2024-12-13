package logic

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/spf13/viper"
)

// StartTunnel starts the port forwarding for a given node's chirpstack server.
// 	node: the vsn of the node.
//	port: localhost port to forward to.
func StartTunnel(node, port string) {
    fmt.Printf("Starting tunnel to %s for ChirpStack access on port %s...\n", node, port)

	// Retrieve tunnel information from the config file
	tunnels := viper.GetStringMap("lora.port_forwading")
	tunnelInfo, exists := tunnels[strings.ToLower(node)]
	if exists {
		// Extract the port configuration
		tunnel, ok := tunnelInfo.(map[string]interface{})
		if !ok {
			fmt.Printf("Invalid port forwading configuration for node %s.\n", node)
			fmt.Printf("Fix the config file.\n")
			return
		}
		port, _ := tunnel["port"].(string)
		fmt.Printf("A ChirpStack tunnel is already active for node %s on port %s.\n", node, port)
		fmt.Printf("Visit portal at http://localhost:%s\n", port)
		return
	}

    // Check if the port is already being used by an existing SSH proxy
    cmd := exec.Command("lsof", "-i", fmt.Sprintf(":%s", port))
    lsofOutput, err := cmd.Output()
	if err == nil && len(string(lsofOutput)) > 0 {
        fmt.Printf("Error: Port %s is already in use.\n", port)
		fmt.Printf("Processes using %s.\n", port)
		fmt.Printf("%s", string(lsofOutput))
        return
	} 

    // Get the cluster IP using kubectl via SSH
    cmd = exec.Command("ssh", fmt.Sprintf("node-%s", node), "kubectl get svc wes-chirpstack-server -o=jsonpath=\"{.spec.clusterIP}\"")
    clusterIP, err := cmd.Output()

    if err != nil || len(clusterIP) == 0 {
        fmt.Printf("Error: Unable to get cluster IP for 'wes-chirpstack-server' on %s.\n", node)
        return
    }

    sshArg := fmt.Sprintf("%s:%s:8080", port, strings.TrimSpace(string(clusterIP)))

    // Create the SSH proxy
    proxyCmd := exec.Command("ssh", "-N", "-L", sshArg, fmt.Sprintf("node-%s", node))
    if err := proxyCmd.Start(); err != nil {
        fmt.Printf("Error: Failed to establish tunnel to %s: %v\n", node, err)
        return
    }

    // Save tunnel information in configuration
    tunnels[node] = map[string]string{
        "port": port,
        "ip":   strings.TrimSpace(string(clusterIP)),
    }
    viper.Set("lora.port_forwading", tunnels)

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

    fmt.Printf("ChirpStack tunnel to %s established.\n", node)
    fmt.Printf("Visit portal at http://localhost:%s\n", port)
}

// StopTunnel stops the port forwarding for a given node's chirpstack server.
// 	node: the vsn of the node.
func StopTunnel(node string) {
	fmt.Printf("Stopping ChirpStack tunnel for %s...\n", node)

	// Retrieve tunnel information from the config file
	tunnels := viper.GetStringMap("lora.port_forwading")
	tunnelInfo, exists := tunnels[strings.ToLower(node)]
	if !exists {
		fmt.Printf("No active ChirpStack tunnel found for node %s.\n", node)
		return
	}

    // Extract the port and IP from the configuration
    tunnel, ok := tunnelInfo.(map[string]interface{})
    if !ok {
        fmt.Printf("Invalid port forwading configuration for node %s.\n", node)
		fmt.Printf("Fix the config file.\n")
        return
    }
    port, portExists := tunnel["port"].(string)
    ip, ipExists := tunnel["ip"].(string)
    if !portExists || !ipExists {
        fmt.Printf("Missing port or IP in configuration for node %s.\n", node)
        return
    }
	if !exists {
		fmt.Printf("No active ChirpStack tunnel found for node %s.\n", node)
		return
	}

	// Construct the pattern to match the exact command
	pattern := fmt.Sprintf("ssh -N -L %s:%s:8080 node-%s", port, ip, node)

	// Use pgrep to find the process matching the pattern
	pgrepCmd := exec.Command("pgrep", "-f", pattern)
	pids, err := pgrepCmd.Output()

	if err != nil || len(pids) == 0 {
		fmt.Printf("No active SSH proxy found for node %s.\n", node)
		return
	}

	// Kill the process
	pidList := strings.Fields(string(pids))
	for _, pid := range pidList {
		killCmd := exec.Command("kill", pid)
		if err := killCmd.Run(); err != nil {
			fmt.Printf("Failed to stop ChirpStack tunnel for %s (PID %s): %v\n", node, pid, err)
		}
	}
	fmt.Printf("Stopped\n")

	// Remove the tunnel from the configuration file
	delete(tunnels, strings.ToLower(node))
	viper.Set("lora.port_forwading", tunnels)
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Error updating configuration file: %v\n", err)
	}
}