package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Ensure folder structure to install programs
func Structure() error {

	// Retrieve home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Check for the presence of games folder, if exist, then is ok
	info, err := os.Stat(filepath.Join(home, "Games"))
	if !os.IsNotExist(err) && info.IsDir() {
		return nil
	}

	// Create base games folder structure
	err = cli.Command(`
		mkdir -p $HOME/Games/BIOS
		mkdir -p $HOME/Games/ROMs
		mkdir -p $HOME/Games/Save
	`).Run()

	if err != nil {
		return err
	}

	// Check if must install it on microSD, if no, then is ok
	toMicroSD := cli.Read("INSTALL_TO_MICROSD", "Install to MicroSD? (Y/N)", "N")
	if toMicroSD == "N" {
		return nil
	}

	// Get path to microSD
	microSDPath := cli.Read("MICROSD_PATH", "What is the path of the MicroSD?", "/run/media/mmcblk0p1")
	microSDPath = strings.TrimRight(microSDPath, "/")

	// Make symlinks
	err = cli.Command(fmt.Sprintf(`
		# Remove folders in home to create symlink
		[ -d "$HOME/Games/BIOS" ] && rm -r $HOME/Games/BIOS
		[ -d "$HOME/Games/ROMs" ] && rm -r $HOME/Games/ROMs
		[ -d "$HOME/Games/Save" ] && rm -r $HOME/Games/Save

		# Make sure base folders exist on microSD
		mkdir -p %s/Games/BIOS
		mkdir -p %s/Games/ROMs
		mkdir -p %s/Games/Save

		# Create symlinks
		ln -s %s/Games/BIOS $HOME/Games/BIOS
		ln -s %s/Games/ROMs $HOME/Games/ROMs
		ln -s %s/Games/Save $HOME/Games/Save`,
		microSDPath,
		microSDPath,
		microSDPath,
		microSDPath,
		microSDPath,
		microSDPath,
	)).Run()

	if err != nil {
		return err
	}

	return nil
}
