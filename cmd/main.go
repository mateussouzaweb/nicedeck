package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/gui"
	"github.com/mateussouzaweb/nicedeck/src/nicedeck"
	"github.com/mateussouzaweb/nicedeck/src/server"
)

var version = "Version 0.0.18"
var address = "127.0.0.1:14935"

// Main command
func main() {

	// Exit with proper code
	exitCode := 0
	defer os.Exit(exitCode)

	// Graceful shutdown support
	exit := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	ready := make(chan bool, 1)

	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// Capture exit (CTRL-C)
	go func() {
		<-exit
		done <- true
	}()

	// Perform desktop install
	go func() {
		err := nicedeck.DesktopInstall()
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}
	}()

	// Run the program server
	go func() {
		err := server.Setup(version)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}

		err = server.Start(address, ready)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}

		done <- true
	}()

	// Open UI with server address
	// We should wait for the serve goes up first
	go func() {
		<-ready
		err := gui.OpenWithBrowser("http://"+address, 1280, 720)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}

		done <- true
	}()

	<-done
}
