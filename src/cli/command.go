package cli

import (
	"os"
	"os/exec"
)

// Command creates a new command struct
func Command(script string) *exec.Cmd {

	cmd := exec.Command("/bin/bash", "-c", script)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
