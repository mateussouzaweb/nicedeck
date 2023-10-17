package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/install"
	"github.com/mateussouzaweb/nicedeck/src/roms"
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
		"roms $PLATFORM,...    (parse ROMs folder to add, update or remove ROMs on Steam Library)\n"+
		"shortcuts             (list shortcuts added to the Steam Library)\n"+
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
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
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
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
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

	cli.Printf(cli.ColorSuccess, "Process finished!\n")
	cli.Printf(cli.ColorNotice, "Please restart the device to changes take effect.\n")

	return nil
}

// ROMs command (to update Steam Library)
func runROMs() error {

	// Load Steam library
	err := steam.Load()
	if err != nil {
		return err
	}

	// Save config on finish
	defer func() {
		err := steam.Save()
		if err != nil {
			cli.Printf(cli.ColorFatal, "Error: %s\n", err.Error())
		}
	}()

	// Read advanced command arguments
	args := os.Args[1:]
	platforms := cli.Arg(args, "1", "")

	// Process ROMs to add/update/remove
	err = roms.Process(platforms, false)
	if err != nil {
		return err
	}

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
	shortcuts := steam.GetShortcuts()

	if len(shortcuts) > 0 {
		cli.Printf(cli.ColorNotice, "%s => %s\n", "NAME", "APP_ID")
	}
	for _, shortcut := range shortcuts {
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
