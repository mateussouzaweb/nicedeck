package library

import (
	"os"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Init library by ensure desired environment paths
func Init() error {

	// Retrieve relevant user directories
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	// Set base environment variables when necessary
	cli.SetEnv("HOME", homeDir, false)
	cli.SetEnv("CONFIG", configDir, false)
	cli.SetEnv("CACHE", cacheDir, false)

	// On Windows, add a few shortcuts
	if cli.IsWindows() {
		cli.SetEnv("APPDATA", fs.ExpandPath("$CONFIG"), false)
		cli.SetEnv("PROGRAMS", fs.ExpandPath("$HOMEDRIVE\\Program Files"), false)
		cli.SetEnv("PROGRAMS_X86", fs.ExpandPath("$HOMEDRIVE\\Program Files (x86)"), false)
	}

	// Expose environment variables for internal usage
	cli.SetEnv("GAMES", fs.ExpandPath("$HOME/Games"), false)
	cli.SetEnv("APPLICATIONS", fs.ExpandPath("$GAMES/Applications"), false)
	cli.SetEnv("EMULATORS", fs.ExpandPath("$GAMES/Emulators"), false)
	cli.SetEnv("BIOS", fs.ExpandPath("$GAMES/BIOS"), false)
	cli.SetEnv("ROMS", fs.ExpandPath("$GAMES/ROMs"), false)
	cli.SetEnv("STATE", fs.ExpandPath("$GAMES/State"), false)

	return nil
}
