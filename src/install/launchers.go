package install

import (
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/emulationstation"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Installer for Bottles
func Bottles() *Program {
	return &Program{
		ID:           "bottles",
		Name:         "Bottles",
		Description:  "Run Windows in a Bottle",
		Category:     "Gaming",
		Tags:         []string{"Gaming", "Utilities"},
		FlatpakAppID: "com.usebottles.bottles",
		IconURL:      "https://cdn2.steamgriddb.com/icon/449ef87e4d3fa1f1f268196b185627dd.ico",
		LogoURL:      "https://cdn2.steamgriddb.com/logo/92491efa7cda6552f740334c9e601855.png",
		CoverURL:     "https://cdn2.steamgriddb.com/grid/8845e5d69c0f8a1d4b30334afb030214.png",
		BannerURL:    "https://cdn2.steamgriddb.com/grid/123a00ca793f7db5b771574116bc061f.png",
		HeroURL:      "https://cdn2.steamgriddb.com/hero/84bdc10b5cc3b036ce04a562b0e54d61.png",
	}
}

// Installer for EmulationStation Desktop Edition
func EmulationStationDE() *Program {

	installer := func(shortcut *shortcuts.Shortcut) error {

		// Get latest available version
		latest, err := emulationstation.GetLatestRelease()
		if err != nil {
			return err
		}

		// Download application
		executable := os.ExpandEnv("$HOME/Applications/EmulationStation-DE.AppImage")
		err = fs.DownloadFile(latest, executable, true)
		if err != nil {
			return err
		}

		// Make sure is executable
		err = os.Chmod(executable, 0775)
		if err != nil {
			return err
		}

		// Write settings
		err = emulationstation.WriteSettings()
		if err != nil {
			return err
		}

		// Fill shortcut information
		desktopShortcut := os.ExpandEnv("$HOME/.local/share/applications/emulationstation-de.desktop")
		shortcut.Exe = executable
		shortcut.StartDir = filepath.Dir(executable)
		shortcut.ShortcutPath = desktopShortcut
		shortcut.LaunchOptions = ""

		return nil
	}

	return &Program{
		ID:              "emulationstation",
		Name:            "EmulationStation DE",
		Description:     "Frontend for browsing and launching emulated games",
		Category:        "Gaming",
		Tags:            []string{"Gaming", "Emulator", "Launcher"},
		RequiredFolders: []string{"$HOME/Applications"},
		IconURL:         "https://cdn2.steamgriddb.com/icon/c0829dc52beb665d3e2fd05e36f97f35.ico",
		LogoURL:         "https://cdn2.steamgriddb.com/logo/c3bb9214431dec7ca7d1ebcfeca73236.png",
		CoverURL:        "https://cdn2.steamgriddb.com/grid/21bd6ea21e43de6dc80e2bc8917f4ba3.png",
		BannerURL:       "https://cdn2.steamgriddb.com/grid/67a900732336f1ce9d0c0496352fa9ab.png",
		HeroURL:         "https://cdn2.steamgriddb.com/hero/9323f21f2098b7288267c785458548b2.png",
		Installer:       installer,
	}
}

// Installer for Heroic Games Launcher
func HeroicGamesLauncher() *Program {
	return &Program{
		ID:           "heroic-games",
		Name:         "Heroic Games Launcher",
		Description:  "Launcher for Epic Games, GOG and Prime Gaming",
		Category:     "Gaming",
		Tags:         []string{"Gaming", "Utilities"},
		FlatpakAppID: "com.heroicgameslauncher.hgl",
		IconURL:      "https://cdn2.steamgriddb.com/icon/ae852ba7ae75fa4c5c7d186a61fcce92.png",
		LogoURL:      "https://cdn2.steamgriddb.com/logo/6eebc030d78d41b6cbcf9067aeda9198.png",
		CoverURL:     "https://cdn2.steamgriddb.com/grid/2b1c6cedeaf9571589e3dc9d51ba20e5.png",
		BannerURL:    "https://cdn2.steamgriddb.com/grid/94e8e64cdefe77dcc168855c54f14acd.png",
		HeroURL:      "https://cdn2.steamgriddb.com/hero/bee5ca2551bf346f067a3ac16057bc40.png",
	}
}

// Installer for Lutris
func Lutris() *Program {
	return &Program{
		ID:           "lutris",
		Name:         "Lutris",
		Description:  "Play all your games on Linux",
		Category:     "Gaming",
		Tags:         []string{"Gaming", "Utilities"},
		FlatpakAppID: "net.lutris.Lutris",
		IconURL:      "https://cdn2.steamgriddb.com/icon/8d060abe1e38ab179742bd3af495f407.png",
		LogoURL:      "https://cdn2.steamgriddb.com/logo/bbd451c375fb5b293a9b1f082bf8d024.png",
		CoverURL:     "https://cdn2.steamgriddb.com/grid/3b0d861c2cf5ed4d7b139ee277c8a04a.png",
		BannerURL:    "https://cdn2.steamgriddb.com/grid/3c5bf5a314017c84acae32394125cf26.png",
		HeroURL:      "https://cdn2.steamgriddb.com/hero/3b7f06487067b9aa2393a438dd095edc.png",
	}
}
