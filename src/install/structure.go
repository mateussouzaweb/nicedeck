package install

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Ensure folder structure to install programs
func Structure(installOnMicroSD bool, microSDPath string) error {

	// Check for the presence of games folder in home
	// If exist, then is ok and we can skip
	gamesPath := os.ExpandEnv("$HOME/Games")
	exist, err := fs.DirectoryExist(gamesPath)
	if err != nil {
		return err
	} else if exist {
		return nil
	}

	// Check if must install it on microSD
	// If not, then just create the base games folder structure on home
	if !installOnMicroSD {
		err := cli.Command(`
			mkdir -p $HOME/Games/BIOS
			mkdir -p $HOME/Games/ROMs
		`).Run()

		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "Folder structure created at: %s\n", gamesPath)
		return nil
	}

	// Get microSD path to perform install
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

	cli.Printf(cli.ColorSuccess, "Folder structure created at: %s\n", microSDPath+"/Games")
	cli.Printf(cli.ColorSuccess, "Symlinks available at: %s\n", gamesPath)

	return nil
}
