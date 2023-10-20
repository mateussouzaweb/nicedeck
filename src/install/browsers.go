package install

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

// Install Firefox
func Firefox() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install --or-update --assumeyes --noninteractive flathub org.mozilla.firefox
	`).Run()

	if err != nil {
		return err
	}

	// Add to Steam
	err = steam.AddToShortcuts(&shortcuts.Shortcut{
		AppName:       "Firefox",
		StartDir:      "/var/lib/flatpak/exports/bin/",
		Exe:           "/var/lib/flatpak/exports/bin/org.mozilla.firefox",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/org.mozilla.firefox.desktop",
		LaunchOptions: "",
		Tags:          []string{"Utilities"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/f968fdc88852a4a3a27a81fe3f57bfc5.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/43285a8b542fcdc35377439e05dcb04f.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/4529f985441a035ae4a107b8862ba4dd.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/9384fe92aef7ea0128be2c916ed07cea.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/a318166b8539611449bf21ddc297a783.png",
	})

	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Firefox installed!\n")
	return nil
}

// Install Google Chrome
func GoogleChrome() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install --or-update --assumeyes --noninteractive flathub com.google.Chrome
	`).Run()

	if err != nil {
		return err
	}

	// Add to Steam
	err = steam.AddToShortcuts(&shortcuts.Shortcut{
		AppName:       "Google Chrome",
		StartDir:      "/var/lib/flatpak/exports/bin/",
		Exe:           "/var/lib/flatpak/exports/bin/com.google.Chrome",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.google.Chrome.desktop",
		LaunchOptions: "",
		Tags:          []string{"Utilities"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/3941c4358616274ac2436eacf67fae05.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/3b049d0f6cbf5421d399f156807b8657.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/d45c26607db83f6f14b09dd70123913b.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/d40c243072a2d2957b3484e775f1f925.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/cae83cfcb1d8a2a4bb17bd1446fb1cee.png",
	})

	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Google Chrome installed!\n")
	return nil
}
