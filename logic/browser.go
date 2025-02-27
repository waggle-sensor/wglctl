package logic

import (
	"runtime"
	"os/exec"
)

// openBrowser tries to open the URL in the default web browser.
func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default: // Linux and other UNIX-based OS
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Start()
}