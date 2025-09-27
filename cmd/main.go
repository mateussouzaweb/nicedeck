package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/command"
)

var version = "0.2.0"

// Main command
func main() {

	// Exit with proper code
	exitCode := 0
	defer os.Exit(exitCode)

	// Graceful init and shutdown support
	exit := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// Capture exit (CTRL-C)
	go func() {
		<-exit
		done <- true
	}()

	// Run command
	go func() {
		err := command.Run(version, os.Args[1:], done)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
		}

		done <- true
	}()

	<-done
}
