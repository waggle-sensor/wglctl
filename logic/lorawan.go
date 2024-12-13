package logic

import (
	"fmt"
	"os/exec"
)

// getChirpStackIp gets the cluster ip for the ChirpStack server.
// 	node: the vsn of the node.
func GetChirpStackIp(node string) string {
	// Get the cluster IP using kubectl via SSH
    cmd := exec.Command("ssh", fmt.Sprintf("node-%s", node), "kubectl get svc wes-chirpstack-server -o=jsonpath=\"{.spec.clusterIP}\"")
    clusterIP, err := cmd.Output()

    if err != nil || len(clusterIP) == 0 {
        fmt.Printf("Error: Unable to get cluster IP for 'wes-chirpstack-server' on %s.\n", node)
        return ""
    }

	return string(clusterIP)
}