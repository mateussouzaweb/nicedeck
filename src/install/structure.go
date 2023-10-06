package install

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Ensure folder structure to install programs
func Structure() error {

	// Check for the presence of games folder in home
	// If exist, then is ok and we can skip
	info, err := os.Stat(os.ExpandEnv("$HOME/Games"))
	if !os.IsNotExist(err) && info.IsDir() {
		return nil
	}

	// Check if must install it on microSD
	// If not, then just create the base games folder structure on home
	installOnMicroSD := cli.Read("INSTALL_ON_MICROSD", "Install on MicroSD? Y/N", "N")
	if strings.ToUpper(installOnMicroSD) != "Y" {
		return cli.Command(`
			mkdir -p $HOME/Games/BIOS
			mkdir -p $HOME/Games/ROMs
		`).Run()
	}

	// Get microSD path to install on it
	microSDPath := cli.Read("MICROSD_PATH", "What is the path of the MicroSD?", "/run/media/mmcblk0p1")
	microSDPath = strings.TrimRight(microSDPath, "/")

	// Make symlinks
	err = cli.Command(fmt.Sprintf(`
		# Make sure base folders exist on microSD
		mkdir -p %s/Games/BIOS
		mkdir -p %s/Games/ROMs

		# Make sure base folder exist on home
		mkdir -p $HOME/Games

		# Create symlinks on home from microSD folders
		ln -s %s/Games/BIOS $HOME/Games/BIOS
		ln -s %s/Games/ROMs $HOME/Games/ROMs`,
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
