package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/install"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

// Create mapping for easy install
var installMap = map[string]func() error{
	"bottles":          install.Bottles,
	"cemu":             install.Cemu,
	"citra":            install.Citra,
	"dolphin":          install.Dolphin,
	"emulationstation": install.EmulationStationDE,
	"firefox":          install.Firefox,
	"flycast":          install.Flycast,
	"google-chrome":    install.GoogleChrome,
	"heroic-games":     install.HeroicGamesLauncher,
	"jellyfin":         install.JellyfinMediaPlayer,
	"lutris":           install.Lutris,
	"melonds":          install.MelonDS,
	"mgba":             install.MGBA,
	"moonlight":        install.MoonlightGameStreaming,
	"pcsx2":            install.PCSX2,
	"ppsspp":           install.PPSSPP,
	"rpcs3":            install.RPCS3,
	"ryujinx":          install.Ryujinx,
	"xemu":             install.Xemu,
	"yuzu":             install.Yuzu,
}

// Version command
func printVersion() error {
	cli.Printf(cli.ColorDefault, "Version 0.0.9\n")
	return nil
}

// Help command
func printHelp() error {

	programs := make([]string, 0, len(installMap))
	for program := range installMap {
		programs = append(programs, program)
	}

	cli.Printf(cli.ColorDefault, "\n"+
		"NiceDeck usage help:\n"+
		"\n"+
		"version               (show version)\n"+
		"help                  (print this help)\n"+
		"setup                 (install all programs)\n"+
		"install $PROGRAM,...  (install specific program or programs)\n"+
		"shortcuts             (list Steam shortcuts with respective app id)\n"+
		"\n"+
		"Available programs to install: %s\n"+
		"\n",
		strings.Join(programs, ", "),
	)

	return nil
}

// Setup command (to install all programs)
func runSetup() error {

	// Load Steam library
	err := steam.Load()
	if err != nil {
		return err
	}

	// Save config on finish
	defer func() {
		err := steam.Save()
		if err != nil {
			cli.Printf(cli.ColorFatal, "%s\n", err.Error())
		}
	}()

	// Make sure has required structure
	err = install.Structure()
	if err != nil {
		return err
	}

	// Install each program
	for _, command := range installMap {
		err := command()
		if err != nil {
			return err
		}
	}

	cli.Printf(cli.ColorSuccess, "All programs installed!\n")
	cli.Printf(cli.ColorNotice, "Please restart the device to changes take effect.\n")

	return nil
}

// Install command (for specific programs only)
func runInstall() error {

	// Load Steam library
	err := steam.Load()
	if err != nil {
		return err
	}

	// Save config on finish
	defer func() {
		err := steam.Save()
		if err != nil {
			cli.Printf(cli.ColorFatal, "%s\n", err.Error())
		}
	}()

	// Make sure has required structure
	err = install.Structure()
	if err != nil {
		return err
	}

	// Install selected programs
	args := os.Args[1:]
	programs := cli.Arg(args, "1", "")
	programs = strings.ReplaceAll(programs, " ", "")

	for _, program := range strings.Split(programs, ",") {
		if command, ok := installMap[program]; ok {
			err := command()
			if err != nil {
				return err
			}
		} else {
			cli.Printf(cli.ColorWarn, "Program not found to install: %s\n", program)
		}
	}

	cli.Printf(cli.ColorSuccess, "Programs installed!\n")
	cli.Printf(cli.ColorNotice, "Please restart the device to changes take effect.\n")

	return nil
}

// List shortcuts command
func listShortcuts() error {

	// Load Steam library
	err := steam.Load()
	if err != nil {
		return err
	}

	// List detected shortcuts
	for _, shortcut := range steam.GetShortcuts() {
		cli.Printf(cli.ColorDefault, "%s => %v\n", shortcut.AppName, shortcut.AppID)
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
	case "shortcuts":
		err = listShortcuts()
	case "":
		err = fmt.Errorf("command is required")
	default:
		err = fmt.Errorf("unknown command: %s", subCommand)
	}

	if err != nil {
		cli.Printf(cli.ColorFatal, "%s\n", err.Error())
		return
	}

}
