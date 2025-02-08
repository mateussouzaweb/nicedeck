package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// Command creates a new command struct
func Command(script string) *exec.Cmd {

	var cmd *exec.Cmd

	// Use PowerShell on Windows
	// Use Bash on MacOS and Linux
	if IsWindows() {
		cmd = exec.Command("powershell", script)
	} else if IsMacOS() {
		cmd = exec.Command("bash", "-c", script)
	} else {
		cmd = exec.Command("/bin/bash", "-c", script)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

// Run script and return captured error if there is any
func Run(script string) error {

	var stderr bytes.Buffer
	cmd := Command(script)
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", err, stderr.String())
	}

	return nil
}

// Start process with blocking channel
func Start(script string) error {

	// Start the command
	command := Command(script)
	err := command.Start()
	if err != nil {
		return err
	}

	// Waiting until it closes and report back to main channel
	finished := make(chan bool, 1)

	go func() {
		err = command.Wait()
		finished <- true
	}()

	<-finished
	return err
}
