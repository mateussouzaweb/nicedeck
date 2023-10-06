package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/install"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

func main() {

	// Create mapping for easy install
	installMap := map[string]func() error{
		"bottles":       install.Bottles,
		"cemu":          install.Cemu,
		"citra":         install.Citra,
		"dolphin":       install.Dolphin,
		"firefox":       install.Firefox,
		"flycast":       install.Flycast,
		"google-chrome": install.GoogleChrome,
		"heroic-games":  install.HeroicGamesLauncher,
		"jellyfin":      install.JellyfinMediaPlayer,
		"lutris":        install.Lutris,
		"melonds":       install.MelonDS,
		"mgba":          install.MGBA,
		"moonlight":     install.MoonlightGameStreaming,
		"pcsx2":         install.PCSX2,
		"ppsspp":        install.PPSSPP,
		"rpcs3":         install.RPCS3,
		"ryujinx":       install.Ryujinx,
		"xemu":          install.Xemu,
		"yuzu":          install.Yuzu,
	}

	args := os.Args[1:]
	subCommand := cli.Arg(args, "0", "")

	// Version command
	if subCommand == "version" {
		fmt.Println("Version 0.0.2")
		return
	}

	// Help command
	if subCommand == "help" {
		programs := make([]string, 0, len(installMap))
		for program := range installMap {
			programs = append(programs, program)
		}

		fmt.Println("")
		fmt.Println("nicedeck version                    (show version)")
		fmt.Println("nicedeck help                       (print this help)")
		fmt.Println("nicedeck setup                      (install all programs)")
		fmt.Println("nicedeck install --programs=KEY,... (install specific program or programs)")
		fmt.Println("nicedeck list-shortcuts             (list steam shortcuts with respective app id)")
		fmt.Println("")
		fmt.Println("Available programs to install: ", strings.Join(programs, ", "))
		fmt.Println("")

		return
	}

	// Retrieve userdata path
	path, err := steam.GetUserPath("/config")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set runtime configs
	config := &steam.Config{
		ArtworksPath:  path + "/grid",
		DebugFile:     path + "/niceconfig.json",
		ShortcutsFile: path + "/shortcuts.vdf",
	}

	save, err := steam.Use(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer save()

	// Setup command (to install all programs)
	if subCommand == "setup" {

		// Make sure structure installation is done
		err = install.Structure()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Install each program
		for _, command := range installMap {
			err := command()
			if err != nil {
				fmt.Println(err)
				break
			}
		}

		return
	}

	// Install command (for specific programs only)
	if subCommand == "install" {

		// Make sure structure installation is done
		err = install.Structure()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Install selected programs
		programs := cli.Arg(args, "--programs", "")
		programs = strings.ReplaceAll(programs, " ", "")

		for _, program := range strings.Split(programs, ",") {
			if command, ok := installMap[program]; ok {
				err := command()
				if err != nil {
					fmt.Println(err)
					break
				}
			} else {
				fmt.Println("Program not found to install:", program)
			}
		}

		return
	}

	// List shortcuts command
	if subCommand == "list-shortcuts" {
		for _, shortcut := range config.Shortcuts {
			fmt.Printf("%s => %v\n", shortcut.AppName, shortcut.AppID)
		}
		return
	}

	fmt.Println("Unknown command.")
}
