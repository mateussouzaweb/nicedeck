package install

import (
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Ensure folder structure to install programs
func Structure(useSymlink bool, storagePath string) error {

	// Check for the presence of games folder in home
	// If exist, then is ok and we can skip
	gamesPath := os.ExpandEnv("$HOME/Games")
	exist, err := fs.DirectoryExist(gamesPath)
	if err != nil {
		return err
	} else if exist {
		return nil
	}

	// Check if must install it on another location with symlink
	// If not, then just create the base games folder structure on home
	if !useSymlink {
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

	// Get storage path to perform install
	storagePath = strings.TrimRight(storagePath, "/")

	// Make symlinks
	err = cli.Command(fmt.Sprintf(`
		# Make sure base folders exist on storage path
		mkdir -p %s/Games/BIOS
		mkdir -p %s/Games/ROMs

		# Make sure base folder exist on home
		mkdir -p $HOME/Games

		# Create symlinks on home from storage path folders
		ln -s %s/Games/BIOS $HOME/Games/BIOS
		ln -s %s/Games/ROMs $HOME/Games/ROMs`,
		storagePath,
		storagePath,
		storagePath,
		storagePath,
	)).Run()

	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Folder structure created at: %s\n", storagePath+"/Games")
	cli.Printf(cli.ColorSuccess, "Symlinks available at: %s\n", gamesPath)

	return nil
}
