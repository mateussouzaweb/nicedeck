package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/docs"
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/install"
	"github.com/mateussouzaweb/nicedeck/src/library"
	"github.com/mateussouzaweb/nicedeck/src/roms"
)

// Version command
func printVersion() error {
	cli.Printf(cli.ColorDefault, "Version 0.0.12\n")
	return nil
}

// Help command
func printHelp() error {

	// Read help content
	content, err := docs.GetContent("HELP.md", true)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorDefault, content)
	return nil
}

// Setup command (to install all programs)
func runSetup() error {

	// Load library
	err := library.Load()
	if err != nil {
		return err
	}

	// Save config on finish
	defer func() {
		err := library.Save()
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
		}
	}()

	// Make sure has required structure
	err = install.Structure()
	if err != nil {
		return err
	}

	// Install each program
	for _, program := range install.GetPrograms() {
		err := install.Install(program.ID)
		if err != nil {
			return err
		}
	}

	cli.Printf(cli.ColorSuccess, "All programs installed!\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	return nil
}

// Install command (for specific programs only)
func runInstall() error {

	// Load library
	err := library.Load()
	if err != nil {
		return err
	}

	// Save config on finish
	defer func() {
		err := library.Save()
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
		}
	}()

	// Make sure has required structure
	err = install.Structure()
	if err != nil {
		return err
	}

	// Read advanced command arguments
	args := os.Args[1:]
	programs := cli.Arg(args, "1,--programs", "")
	programs = strings.ReplaceAll(programs, " ", "")

	// Install programs in the list
	for _, program := range strings.Split(programs, ",") {
		err := install.Install(program)
		if err != nil {
			return err
		}
	}

	cli.Printf(cli.ColorSuccess, "Process finished!\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	return nil
}

// ROMs command (to update library)
func runROMs() error {

	// Load library
	err := library.Load()
	if err != nil {
		return err
	}

	// Save config on finish
	defer func() {
		err := library.Save()
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
		}
	}()

	// Read advanced command arguments
	args := os.Args[1:]
	platforms := cli.Arg(args, "1,--platforms", "")
	preferences := cli.Arg(args, "--preferences", "")
	rebuild := cli.Flag(args, "--rebuild", false)

	// Process ROMs to add/update/remove
	options := roms.ToOptions(platforms, preferences, rebuild)
	err = roms.Process(options)
	if err != nil {
		return err
	}

	return nil
}

// List shortcuts command
func listShortcuts() error {

	// Load library
	err := library.Load()
	if err != nil {
		return err
	}

	// List detected shortcuts
	shortcuts := library.GetShortcuts()
	for _, shortcut := range shortcuts {
		cli.Printf(cli.ColorDefault, "[%v]: %s\n", shortcut.AppID, shortcut.AppName)
	}

	return nil
}

// Main command
func main() {

	args := os.Args[1:]
	subCommand := cli.Arg(args, "0", "")

	var err error

	switch subCommand {
	case "version":
		err = printVersion()
	case "help":
		err = printHelp()
	case "setup":
		err = runSetup()
	case "install":
		err = runInstall()
	case "roms":
		err = runROMs()
	case "shortcuts":
		err = listShortcuts()
	case "":
		err = fmt.Errorf("command is required")
	default:
		err = fmt.Errorf("unknown command: %s", subCommand)
	}

	if err != nil {
		cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
		return
	}

}
