package main

import (
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/install"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

func main() {

	// Create mapping for easy install
	installMap := map[string]func() error{
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

	args := os.Args[1:]
	subCommand := cli.Arg(args, "0", "")

	// Version command
	if subCommand == "version" {
		cli.Printf(cli.ColorDefault, "Version 0.0.5\n")
		return
	}

	// Help command
	if subCommand == "help" {
		programs := make([]string, 0, len(installMap))
		for program := range installMap {
			programs = append(programs, program)
		}

		cli.Printf(cli.ColorDefault, "\n"+
			"nicedeck version                    (show version)\n"+
			"nicedeck help                       (print this help)\n"+
			"nicedeck setup                      (install all programs)\n"+
			"nicedeck install --programs=KEY,... (install specific program or programs)\n"+
			"nicedeck list-shortcuts             (list steam shortcuts with respective app id)\n"+
			"\n"+
			"Available programs to install: %s\n"+
			"\n",
			strings.Join(programs, ", "),
		)
		return
	}

	// Retrieve user config path
	userConfig, err := steam.GetPath("userdata/*/config")
	if err != nil {
		cli.Printf(cli.ColorFatal, "%s\n", err.Error())
		return
	}

	// Retrieve controller templates path
	controllerTemplates, err := steam.GetPath("controller_base/templates")
	if err != nil {
		cli.Printf(cli.ColorFatal, "%s\n", err.Error())
		return
	}

	// Set runtime configs
	config := &steam.Config{
		ArtworksPath:   userConfig + "/grid",
		DebugFile:      userConfig + "/niceconfig.json",
		ShortcutsFile:  userConfig + "/shortcuts.vdf",
		ControllerFile: controllerTemplates + "/controller_neptune_nicedeck.vdf",
	}

	saveConfig, err := steam.Use(config)
	if err != nil {
		cli.Printf(cli.ColorFatal, "%s\n", err.Error())
		return
	}

	// Setup command (to install all programs)
	if subCommand == "setup" {

		// Save config on finish
		defer func() {
			err := saveConfig()
			if err != nil {
				cli.Printf(cli.ColorFatal, "%s\n", err.Error())
			}
		}()

		// Make sure structure installation is done
		err = install.Structure()
		if err != nil {
			cli.Printf(cli.ColorFatal, "%s\n", err.Error())
			return
		}

		// Install each program
		for _, command := range installMap {
			err := command()
			if err != nil {
				cli.Printf(cli.ColorFatal, "%s\n", err.Error())
				break
			}
		}

		cli.Printf(cli.ColorSuccess, "All programs installed!\n")
		cli.Printf(cli.ColorNotice, "Please restart the device to changes take effect.\n")

		return
	}

	// Install command (for specific programs only)
	if subCommand == "install" {

		// Save config on finish
		defer func() {
			err := saveConfig()
			if err != nil {
				cli.Printf(cli.ColorFatal, "%s\n", err.Error())
			}
		}()

		// Make sure structure installation is done
		err = install.Structure()
		if err != nil {
			cli.Printf(cli.ColorFatal, "%s\n", err.Error())
			return
		}

		// Install selected programs
		programs := cli.Arg(args, "--programs", "")
		programs = strings.ReplaceAll(programs, " ", "")

		for _, program := range strings.Split(programs, ",") {
			if command, ok := installMap[program]; ok {
				err := command()
				if err != nil {
					cli.Printf(cli.ColorFatal, "%s\n", err.Error())
					break
				}
			} else {
				cli.Printf(cli.ColorWarn, "Program not found to install: %s\n", program)
			}
		}

		cli.Printf(cli.ColorSuccess, "Programs installed!\n")
		cli.Printf(cli.ColorNotice, "Please restart the device to changes take effect.\n")

		return
	}

	// List shortcuts command
	if subCommand == "list-shortcuts" {
		for _, shortcut := range config.Shortcuts {
			cli.Printf(cli.ColorDefault, "%s => %v\n", shortcut.AppName, shortcut.AppID)
		}
		return
	}

	cli.Printf(cli.ColorFatal, "Unknown command: %s\n", subCommand)
}
