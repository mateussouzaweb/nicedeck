package library

import (
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Setup library structure to install programs
func Setup(useSymlink bool, storagePath string) error {

	// Check for the presence of root folder
	// If exist, then is ok and we can skip
	rootPath := fs.ExpandPath("$GAMES")
	exist, err := fs.DirectoryExist(rootPath)
	if err != nil {
		return err
	} else if exist {
		cli.Printf(cli.ColorWarn, "Setup skipped...\n")
		cli.Printf(cli.ColorWarn, "Folder structure already exists at: %s\n", rootPath)
		return nil
	}

	// Start by making sure root folder exist
	err = os.MkdirAll(rootPath, 0755)
	if err != nil {
		return err
	}

	// Maps folders and optional symbolic links
	symlinkPath := filepath.Join(storagePath, "Games")
	mappedPaths := map[string]string{
		fs.ExpandPath("$APPLICATIONS"): filepath.Join(symlinkPath, "APPLICATIONS"),
		fs.ExpandPath("$BIOS"):         filepath.Join(symlinkPath, "BIOS"),
		fs.ExpandPath("$ROMS"):         filepath.Join(symlinkPath, "ROMs"),
		fs.ExpandPath("$STATE"):        filepath.Join(symlinkPath, "STATE"),
	}

	// Create folders or links based on user choice
	for source, link := range mappedPaths {
		if !useSymlink {
			err = os.MkdirAll(source, 0755)
			if err != nil {
				return err
			}
		} else {
			err = os.MkdirAll(link, 0755)
			if err != nil {
				return err
			}
			err = os.Symlink(link, source)
			if err != nil {
				return err
			}
		}
	}

	// Check if must install it on another location with symlink
	// If not, then just create the base games folder structure on home
	if !useSymlink {
		cli.Printf(cli.ColorSuccess, "Setup completed!\n")
		cli.Printf(cli.ColorSuccess, "Folder structure created at: %s\n", rootPath)
		return nil
	}

	cli.Printf(cli.ColorSuccess, "Setup completed!\n")
	cli.Printf(cli.ColorSuccess, "Folder structure created at: %s\n", symlinkPath)
	cli.Printf(cli.ColorSuccess, "Symlinks available at: %s\n", rootPath)
	return nil
}
