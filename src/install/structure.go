package install

import (
	"os"
	"path/filepath"

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
		cli.Printf(cli.ColorWarn, "Setup skipped...\n")
		cli.Printf(cli.ColorWarn, "Folder structure already exists at %s\n", gamesPath)
		return nil
	}

	// Start by making sure base folder exist on home
	err = os.MkdirAll(gamesPath, 0755)
	if err != nil {
		return err
	}

	BIOSPath := filepath.Join(gamesPath, "BIOS")
	ROMsPath := filepath.Join(gamesPath, "ROMs")

	// Check if must install it on another location with symlink
	// If not, then just create the base games folder structure on home
	if !useSymlink {
		err = os.MkdirAll(BIOSPath, 0755)
		if err != nil {
			return err
		}

		err = os.MkdirAll(ROMsPath, 0755)
		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "Setup completed!\n")
		cli.Printf(cli.ColorSuccess, "Folder structure created at: %s\n", gamesPath)
		return nil
	}

	// Get storage path to perform install
	storagePath = filepath.Join(storagePath, "Games")
	storageBIOSPath := filepath.Join(storagePath, "BIOS")
	storageROMsPath := filepath.Join(storagePath, "ROMs")

	// Make sure base folders exist on storage path
	err = os.MkdirAll(storageBIOSPath, 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(storageROMsPath, 0755)
	if err != nil {
		return err
	}

	// Create symlinks on home from storage path folders
	err = os.Symlink(storageBIOSPath, BIOSPath)
	if err != nil {
		return err
	}

	err = os.Symlink(storageROMsPath, ROMsPath)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Setup completed!\n")
	cli.Printf(cli.ColorSuccess, "Folder structure created at: %s\n", storagePath+"/Games")
	cli.Printf(cli.ColorSuccess, "Symlinks available at: %s\n", gamesPath)

	return nil
}
