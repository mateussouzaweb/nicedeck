package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// Command creates a new command struct
func Command(script string) *exec.Cmd {

	var command *exec.Cmd

	// Use PowerShell on Windows
	// Use Bash on MacOS and Linux
	if IsWindows() {
		command = exec.Command("powershell", script)
	} else if IsMacOS() {
		command = exec.Command("bash", "-c", script)
	} else {
		command = exec.Command("/bin/bash", "-c", script)
	}

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command
}

// Run script and return captured error if there is any
func Run(command *exec.Cmd) error {

	var stderr bytes.Buffer
	command.Stderr = &stderr

	err := command.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", err, stderr.String())
	}

	return nil
}

// Start process with blocking channel
func Start(command *exec.Cmd) error {

	// Start the command
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
