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
		Tags:          []string{"PROGRAMS"},
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
		Tags:          []string{"PROGRAMS"},
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

// Install Jellyfin Media Player
func JellyfinMediaPlayer() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install --or-update --assumeyes --noninteractive flathub com.github.iwalton3.jellyfin-media-player
	`).Run()

	if err != nil {
		return err
	}

	// Add to Steam
	err = steam.AddToShortcuts(&shortcuts.Shortcut{
		AppName:       "Jellyfin Media Player",
		StartDir:      "/var/lib/flatpak/exports/bin/",
		Exe:           "/var/lib/flatpak/exports/bin/com.github.iwalton3.jellyfin-media-player",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.github.iwalton3.jellyfin-media-player.desktop",
		LaunchOptions: "",
		Tags:          []string{"PROGRAMS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/bbe2977a4c5b136df752894d93b44c72.png",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/c84389bbba219be3e13b80f9376a0db7.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/43174d6e1d2f2791af4925de27e813af.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/717a3aabd77351296bbf24f7274a4d6e.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/5e38c6c14e095dd7b30db8c0fdba643a.png",
	})

	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Jellyfin Media Player installed!\n")
	return nil
}

// Install Moonlight Game Streaming
func MoonlightGameStreaming() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install --or-update --assumeyes --noninteractive flathub com.moonlight_stream.Moonlight
	`).Run()

	if err != nil {
		return err
	}

	// Add to Steam
	err = steam.AddToShortcuts(&shortcuts.Shortcut{
		AppName:       "Moonlight Game Streaming",
		StartDir:      "/var/lib/flatpak/exports/bin/",
		Exe:           "/var/lib/flatpak/exports/bin/com.moonlight_stream.Moonlight",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.moonlight_stream.Moonlight.desktop",
		LaunchOptions: "",
		Tags:          []string{"PROGRAMS"},
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/ef8051ce270059a142fcb0b3e47b1cd4.png",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/beb5ad322e679d0a6045c6cfc56e8b92.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/030d60c36d51783da9e4cbb6aa5abd2c.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/8a8f67cacf3e3d2d63614f515a2079b8.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/0afefa2281c2f8b0b86d6332e2cdbe7d.png",
	})

	if err != nil {
		return err
	}

	cli.Printf(cli.ColorSuccess, "Moonlight Game Streaming installed!\n")
	return nil
}
