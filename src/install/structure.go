package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Ensure folder structure to install programs
func Structure() error {

	// Retrieve home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Check for the presence of games folder, if exist, then is ok
	if fs.ExistDirectory(filepath.Join(home, "Games")) {
		return nil
	}

	// Create base games folder structure
	err = cli.Command(fmt.Sprintf(`
		mkdir -p %s/Games/BIOS
		mkdir -p %s/Games/ROMs
		mkdir -p %s/Games/Save
	`, home, home, home)).Run()

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
		return err
	}

	return nil
}
