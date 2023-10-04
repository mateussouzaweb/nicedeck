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

	// Working directory should be home always
	// This will not affect the shell working directory
	err := cli.EnsureHome()
	if err != nil {
		fmt.Println(err)
		return
	}

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
		"moonlight":     install.Moonlight,
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
		fmt.Println("Version 0.0.1")
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
		fmt.Println("")
		fmt.Println("Available programs to install: ", strings.Join(programs, ", "))

		return
	}

	// Retrieve userdata path
	path, err := steam.GetUserPath("/config")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set runtime configs
	save, err := steam.Use(&steam.Config{
		ArtworksPath:      path + "/grid",
		ShortcutsFilePath: path + "/shortcuts.vdf",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	defer save()

	// Make sure structure installation is done
	err = install.Structure()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Install command (for specific programs only)
	if subCommand == "install" {

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

	// Setup command (to install all programs)
	if subCommand == "setup" {

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

	fmt.Println("Unknown command.")
}
