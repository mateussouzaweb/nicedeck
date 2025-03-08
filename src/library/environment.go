package library

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Init library by ensure desired environment paths
func Init(version string) error {

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

	// On Linux, add special shortcuts
	if cli.IsLinux() {
		cli.SetEnv("SHARE", fs.ExpandPath("$HOME/.local/share"), true)
		cli.SetEnv("VAR", fs.ExpandPath("$HOME/.var/app"), true)
	}

	// On Windows, add special shortcuts
	if cli.IsWindows() {
		cli.SetEnv("APPDATA", filepath.Dir(configDir), true)
		cli.SetEnv("DOCUMENTS", fs.ExpandPath("$HOME\\Documents"), true)
		cli.SetEnv("PROGRAMS", fs.ExpandPath("$HOMEDRIVE\\Program Files"), true)
		cli.SetEnv("PROGRAMS_X86", fs.ExpandPath("$HOMEDRIVE\\Program Files (x86)"), true)
		cli.SetEnv("START_MENU", fs.ExpandPath("$CONFIG\\Microsoft\\Windows\\Start Menu\\Programs"), true)
	}

	// Expose environment variables for internal usage
	cli.SetEnv("GAMES", fs.ExpandPath("$HOME/Games"), false)
	cli.SetEnv("APPLICATIONS", fs.ExpandPath("$GAMES/Applications"), false)
	cli.SetEnv("EMULATORS", fs.ExpandPath("$GAMES/Emulators"), false)
	cli.SetEnv("BIOS", fs.ExpandPath("$GAMES/BIOS"), false)
	cli.SetEnv("ROMS", fs.ExpandPath("$GAMES/ROMs"), false)
	cli.SetEnv("STATE", fs.ExpandPath("$GAMES/State"), false)

	// Print debug information
	cli.Printf(cli.ColorNotice, "NiceDeck\n")
	cli.Printf(cli.ColorNotice, "\n")
	cli.Printf(cli.ColorNotice, "- Version: %s\n", version)
	cli.Printf(cli.ColorNotice, "- OS: %s-%s\n", runtime.GOOS, runtime.GOARCH)
	cli.Printf(cli.ColorNotice, "- Home: %s\n", cli.GetEnv("HOME", ""))
	cli.Printf(cli.ColorNotice, "- Config: %s\n", cli.GetEnv("CONFIG", ""))
	cli.Printf(cli.ColorNotice, "- Cache: %s\n", cli.GetEnv("CACHE", ""))
	cli.Printf(cli.ColorNotice, "- Games: %s\n", cli.GetEnv("GAMES", ""))
	cli.Printf(cli.ColorNotice, "- Applications: %s\n", cli.GetEnv("APPLICATIONS", ""))
	cli.Printf(cli.ColorNotice, "- Emulators: %s\n", cli.GetEnv("EMULATORS", ""))
	cli.Printf(cli.ColorNotice, "- BIOS: %s\n", cli.GetEnv("BIOS", ""))
	cli.Printf(cli.ColorNotice, "- ROMs: %s\n", cli.GetEnv("ROMS", ""))
	cli.Printf(cli.ColorNotice, "- State: %s\n", cli.GetEnv("STATE", ""))

	if cli.IsLinux() {
		cli.Printf(cli.ColorNotice, "- Share: %s\n", cli.GetEnv("SHARE", ""))
		cli.Printf(cli.ColorNotice, "- Var: %s\n", cli.GetEnv("VAR", ""))
	}

	if cli.IsWindows() {
		cli.Printf(cli.ColorNotice, "- App Data: %s\n", cli.GetEnv("APPDATA", ""))
		cli.Printf(cli.ColorNotice, "- Documents: %s\n", cli.GetEnv("DOCUMENTS", ""))
		cli.Printf(cli.ColorNotice, "- Programs: %s\n", cli.GetEnv("PROGRAMS", ""))
		cli.Printf(cli.ColorNotice, "- Programs X86: %s\n", cli.GetEnv("PROGRAMS_X86", ""))
		cli.Printf(cli.ColorNotice, "- Start Menu: %s\n", cli.GetEnv("START_MENU", ""))
	}

	cli.Printf(cli.ColorNotice, "\n")

	return nil
}
