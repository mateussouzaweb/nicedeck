package install

import (
	"fmt"
	"os"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/emulationstation"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

// Install Bottles
func Bottles() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub com.usebottles.bottles
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Bottles",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.usebottles.bottles.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=bottles --file-forwarding com.usebottles.bottles @@u %u @@",
		Tags:          []string{"LAUNCHERS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/449ef87e4d3fa1f1f268196b185627dd.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/92491efa7cda6552f740334c9e601855.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/8845e5d69c0f8a1d4b30334afb030214.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/123a00ca793f7db5b771574116bc061f.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/84bdc10b5cc3b036ce04a562b0e54d61.png",
	})

	return err
}

// Install EmulationStation Desktop Edition
func EmulationStationDE() error {

	// Get latest available version
	latest, err := emulationstation.GetLatestRelease()
	if err != nil {
		return err
	}

	// Download application
	directory := os.ExpandEnv("$HOME/Applications")
	executable := os.ExpandEnv("$HOME/Applications/EmulationStation-DE.AppImage")
	desktopShortcut := os.ExpandEnv("$HOME/.local/share/applications/emulationstation-de.desktop")

	err = cli.Command(fmt.Sprintf(`
		mkdir -p %s
		wget -q -O %s %s
		chmod +x %s
		`,
		directory,
		executable,
		latest,
		executable,
	)).Run()

	if err != nil {
		return err
	}

	// Write configs
	err = emulationstation.WriteConfigs()
	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "EmulationStation DE",
		Exe:           executable,
		StartDir:      directory,
		ShortcutPath:  desktopShortcut,
		LaunchOptions: "",
		Tags:          []string{"LAUNCHERS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/c0829dc52beb665d3e2fd05e36f97f35.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/c3bb9214431dec7ca7d1ebcfeca73236.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/21bd6ea21e43de6dc80e2bc8917f4ba3.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/67a900732336f1ce9d0c0496352fa9ab.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/9323f21f2098b7288267c785458548b2.png",
	})

	return err
}

// Install Heroic Games Launcher
func HeroicGamesLauncher() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub com.heroicgameslauncher.hgl
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Heroic Games Launcher",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.heroicgameslauncher.hgl.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=heroic-run --file-forwarding com.heroicgameslauncher.hgl @@u %u @@",
		Tags:          []string{"LAUNCHERS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/ae852ba7ae75fa4c5c7d186a61fcce92.png",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/6eebc030d78d41b6cbcf9067aeda9198.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/2b1c6cedeaf9571589e3dc9d51ba20e5.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/94e8e64cdefe77dcc168855c54f14acd.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/bee5ca2551bf346f067a3ac16057bc40.png",
	})

	return err
}

// Install Lutris
func Lutris() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub net.lutris.Lutris
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShortcuts(&steam.Shortcut{
		AppName:       "Lutris",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/net.lutris.Lutris.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=lutris --file-forwarding net.lutris.Lutris @@u %U @@",
		Tags:          []string{"LAUNCHERS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/8d060abe1e38ab179742bd3af495f407.png",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/bbd451c375fb5b293a9b1f082bf8d024.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3b0d861c2cf5ed4d7b139ee277c8a04a.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3c5bf5a314017c84acae32394125cf26.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/3b7f06487067b9aa2393a438dd095edc.png",
	})

	return err
}
