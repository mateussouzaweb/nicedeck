package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/gui"
	"github.com/mateussouzaweb/nicedeck/src/server"
)

var version = "0.0.25"

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

	// Retrieve program options
	displayMode := cli.Arg(os.Args[1:], "--gui", "")
	developmentMode := cli.Flag(os.Args[1:], "--dev", false)
	listenAddress := cli.Arg(os.Args[1:], "--address", "127.0.0.1:14935")

	// Expose environment variables for internal usage
	os.Setenv("APPLICATIONS", os.ExpandEnv("$HOME/Applications"))
	os.Setenv("GAMES", os.ExpandEnv("$HOME/Games"))
	os.Setenv("BIOS", os.ExpandEnv("$HOME/Games/BIOS"))
	os.Setenv("ROMS", os.ExpandEnv("$HOME/Games/ROMs"))
	os.Setenv("STATE", os.ExpandEnv("$HOME/Games/STATE"))

	// Run the program server
	go func() {
		err := server.Setup(version, developmentMode)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}

		err = server.Start(listenAddress, ready)
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
		time.Sleep(100 * time.Millisecond)

		err := gui.Open(displayMode, "http://"+listenAddress, version, developmentMode)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}

		done <- true
	}()

	<-done
}
