package library

import (
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Setup library structure to install programs
func Setup(useSymlink bool, storagePath string) error {

	gamesPath := fs.ExpandPath("$GAMES")
	BIOSPath := fs.ExpandPath("$BIOS")
	ROMsPath := fs.ExpandPath("$ROMS")
	statePath := fs.ExpandPath("$STATE")

	// Check for the presence of games folder
	// If exist, then is ok and we can skip
	exist, err := fs.DirectoryExist(gamesPath)
	if err != nil {
		return err
	} else if exist {
		cli.Printf(cli.ColorWarn, "Setup skipped...\n")
		cli.Printf(cli.ColorWarn, "Folder structure already exists at %s\n", gamesPath)
		return nil
	}

	// Start by making sure games folder exist
	err = os.MkdirAll(gamesPath, 0755)
	if err != nil {
		return err
	}

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

		err = os.MkdirAll(statePath, 0755)
		if err != nil {
			return err
		}

		cli.Printf(cli.ColorSuccess, "Setup completed!\n")
		cli.Printf(cli.ColorSuccess, "Folder structure created at: %s\n", gamesPath)
		return nil
	}

	// Get storage path to perform install
	// This mode will use symbolic links
	storageGamesPath := filepath.Join(storagePath, "Games")
	storageBIOSPath := filepath.Join(storagePath, "Games", "BIOS")
	storageROMsPath := filepath.Join(storagePath, "Games", "ROMs")
	storageStatePath := filepath.Join(storagePath, "Games", "STATE")

	// Make sure base folders exist on storage path
	err = os.MkdirAll(storageGamesPath, 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(storageBIOSPath, 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(storageROMsPath, 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(storageStatePath, 0755)
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

	err = os.Symlink(storageStatePath, statePath)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Setup completed!\n")
	cli.Printf(cli.ColorSuccess, "Folder structure created at: %s\n", storageGamesPath)
	cli.Printf(cli.ColorSuccess, "Symlinks available at: %s\n", gamesPath)

	return nil
}
