package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/install"
)

func main() {

	args := os.Args[1:]

	// Mapping for easy install
	installMap := map[string]func() error{
		"xemu":          install.Xemu,
		"citra":         install.Citra,
		"melonds":       install.MelonDS,
		"mgba":          install.MGBA,
		"dolphin":       install.Dolphin,
		"yuzu":          install.Yuzu,
		"ryujinx":       install.Ryujinx,
		"cemu":          install.Cemu,
		"flycast":       install.Flycast,
		"pcsx2":         install.PCSX2,
		"rpcs3":         install.RPCS3,
		"ppsspp":        install.PPSSPP,
		"google-chrome": install.GoogleChrome,
		"firefox":       install.Firefox,
		"moonlight":     install.Moonlight,
		"heroic-games":  install.HeroicGamesLauncher,
		"lutris":        install.Lutris,
		"bottles":       install.Bottles,
		"jellyfin":      install.JellyfinMediaPlayer,
	}

	// Retrieve home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check for the presence of games folder
	if !fs.DirectoryExist(filepath.Join(home, "Games")) {

		// Create base games folder
		err := cli.Command(fmt.Sprintf(`
			mkdir -p %s/Games/BIOS
			mkdir -p %s/Games/ROMs
			mkdir -p %s/Games/Save
		`, home, home, home)).Run()

		if err != nil {
			fmt.Println(err)
			return
		}

		// Check if must install it on microSD
		toMicroSD := cli.Read("INSTALL_TO_MICROSD", "Install to MicroSD? (Y/N)", "N")
		if toMicroSD == "Y" {

			microSDPath := cli.Read("MICROSD_PATH", "What is the path of the MicroSD?", "/run/media/MicroSD")
			microSDPath = strings.TrimRight(microSDPath, "/")

			err := cli.Command(fmt.Sprintf(`
				# Remove folders in home to create symlink
				[ -d "%s/Games/BIOS" ] && rm -r %s/Games/BIOS
				[ -d "%s/Games/ROMs" ] && rm -r %s/Games/ROMs
				[ -d "%s/Games/Save" ] && rm -r %s/Games/Save

				# Make sure base folder exist on microSD
				mkdir -p %s/Games

				# Create symlinks
				ln -s %s/Games/BIOS %s/Games/BIOS
				ln -s %s/Games/ROMs %s/Games/ROMs
				ln -s %s/Games/Save %s/Games/Save`,
				home, home,
				home, home,
				home, home,
				microSDPath,
				microSDPath, home,
				microSDPath, home,
				microSDPath, home,
			)).Run()

			if err != nil {
				fmt.Println(err)
				return
			}
		}

	}

	// Move working directory to home folder
	err = os.Chdir(home)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Install specific programs only
	install := cli.Arg(args, "--install", "")
	if install != "" {
		for _, program := range strings.Split(install, ",") {
			if command, ok := installMap[program]; ok {
				err := command()
				if err != nil {
					fmt.Println(err)
					break
				}
			} else {
				fmt.Println("Program not found: ", program)
			}
		}
	}

	// Default action, install all programs
	for _, command := range installMap {
		err := command()
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}
