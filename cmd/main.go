package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/gui"
	"github.com/mateussouzaweb/nicedeck/src/server"
)

var version = "0.0.26"

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

	// Define needed variables for Windows
	if cli.IsWindows() {
		cli.SetEnv("HOME", fs.ExpandPath("${HOMEDRIVE}${HOMEPATH}"), false)
	}

	// Expose environment variables for internal usage
	cli.SetEnv("APPLICATIONS", fs.ExpandPath("$HOME/Applications"), false)
	cli.SetEnv("GAMES", fs.ExpandPath("$HOME/Games"), false)
	cli.SetEnv("BIOS", fs.ExpandPath("$GAMES/BIOS"), false)
	cli.SetEnv("ROMS", fs.ExpandPath("$GAMES/ROMs"), false)
	cli.SetEnv("STATE", fs.ExpandPath("$GAMES/STATE"), false)

	// Init server
	go func() {
		err := server.Init(version, developmentMode, listenAddress, ready, done)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}
	}()

	// Open UI with server address
	go func() {
		<-ready
		err := gui.Open(displayMode, targetURL)
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
			exitCode = 1
			done <- true
		}
	}()

	<-done
}
