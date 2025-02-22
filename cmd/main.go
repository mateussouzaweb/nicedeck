package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/server"
)

var version = "0.1.0"

// Main command
func main() {

	// Exit with proper code
	exitCode := 0
	defer os.Exit(exitCode)

	// Graceful init and shutdown support
	exit := make(chan os.Signal, 1)
	ready := make(chan bool, 1)
	done := make(chan bool, 1)

	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// Capture exit (CTRL-C)
	go func() {
		<-exit
		done <- true
	}()

	// Retrieve program options
	displayMode := cli.Arg(os.Args[1:], "--gui", "")
	developmentMode := cli.Flag(os.Args[1:], "--dev", false)
	listenAddress := cli.Arg(os.Args[1:], "--address", "127.0.0.1:14935")
	targetURL := "http://" + listenAddress

	// Init server
	go func() {
		err := server.Init(version, developmentMode, listenAddress, ready, done)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}
	}()

	// Open UI with target URL
	go func() {
		<-ready

		// Headless mode
		if displayMode == "headless" {
			cli.Printf(cli.ColorWarn, "Running in headless mode...\n")
			cli.Printf(cli.ColorWarn, "Please open the following link in the navigator to use the app: %s\n", targetURL)
			return
		}

		// Browser mode
		if err := cli.Open(targetURL); err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}
	}()

	<-done
}
