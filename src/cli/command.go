package cli

import (
	"os"
	"os/exec"
	"strings"
)

// Command creates a new command struct
func Command(script string) *exec.Cmd {

	// Ensure a clean script when outside flatpak
	script = strings.Replace(script, "/usr/bin/flatpak-spawn --host", "", 1)
	cmd := exec.Command("/bin/bash", "-c", script)

	// Apply flatpak-spawn if application is running inside flatpak
	if os.Getenv("FLATPAK_ID") != "" {
		cmd = exec.Command("/usr/bin/flatpak-spawn", "--host", "/bin/bash", "-c", script)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
