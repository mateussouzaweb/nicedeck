package install

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
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
	err = steam.AddToShotcuts(&steam.Shortcut{
		AppName:       "Bottles",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.usebottles.bottles.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=bottles --file-forwarding com.usebottles.bottles @@u %u @@",
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/449ef87e4d3fa1f1f268196b185627dd.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/92491efa7cda6552f740334c9e601855.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/8845e5d69c0f8a1d4b30334afb030214.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/123a00ca793f7db5b771574116bc061f.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/84bdc10b5cc3b036ce04a562b0e54d61.png",
	})

	return err
}

// Install Firefox
func Firefox() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub org.mozilla.firefox
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShotcuts(&steam.Shortcut{
		AppName:       "Firefox",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/org.mozilla.firefox.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=firefox --file-forwarding org.mozilla.firefox @@u %u @@",
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/f968fdc88852a4a3a27a81fe3f57bfc5.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/43285a8b542fcdc35377439e05dcb04f.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/4529f985441a035ae4a107b8862ba4dd.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/9384fe92aef7ea0128be2c916ed07cea.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/a318166b8539611449bf21ddc297a783.png",
	})

	return err
}

// Install Google Chrome
func GoogleChrome() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub com.google.Chrome
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShotcuts(&steam.Shortcut{
		AppName:       "Google Chrome",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.google.Chrome.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=/app/bin/chrome --file-forwarding com.google.Chrome @@u %U @@",
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/3941c4358616274ac2436eacf67fae05.ico",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/3b049d0f6cbf5421d399f156807b8657.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/d45c26607db83f6f14b09dd70123913b.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/d40c243072a2d2957b3484e775f1f925.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/cae83cfcb1d8a2a4bb17bd1446fb1cee.png",
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
	err = steam.AddToShotcuts(&steam.Shortcut{
		AppName:       "Heroic Games Launcher",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.heroicgameslauncher.hgl.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=heroic-run --file-forwarding com.heroicgameslauncher.hgl @@u %u @@",
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/ae852ba7ae75fa4c5c7d186a61fcce92.png",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/6eebc030d78d41b6cbcf9067aeda9198.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/2b1c6cedeaf9571589e3dc9d51ba20e5.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/94e8e64cdefe77dcc168855c54f14acd.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/bee5ca2551bf346f067a3ac16057bc40.png",
	})

	return err
}

// Install Jellyfin Media Player
func JellyfinMediaPlayer() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub com.github.iwalton3.jellyfin-media-player
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShotcuts(&steam.Shortcut{
		AppName:       "Jellyfin Media Player",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.github.iwalton3.jellyfin-media-player.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=jellyfinmediaplayer com.github.iwalton3.jellyfin-media-player",
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/bbe2977a4c5b136df752894d93b44c72.png",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/c84389bbba219be3e13b80f9376a0db7.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/43174d6e1d2f2791af4925de27e813af.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/717a3aabd77351296bbf24f7274a4d6e.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/5e38c6c14e095dd7b30db8c0fdba643a.png",
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
	err = steam.AddToShotcuts(&steam.Shortcut{
		AppName:       "Lutris",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/net.lutris.Lutris.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=lutris --file-forwarding net.lutris.Lutris @@u %U @@",
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/8d060abe1e38ab179742bd3af495f407.png",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/bbd451c375fb5b293a9b1f082bf8d024.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3b0d861c2cf5ed4d7b139ee277c8a04a.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3c5bf5a314017c84acae32394125cf26.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/3b7f06487067b9aa2393a438dd095edc.png",
	})

	return err
}

// Install Moonlight Game Streaming
func MoonlightGameStreaming() error {

	// Install from flatpak
	err := cli.Command(`
		flatpak install -y flathub com.moonlight_stream.Moonlight
	`).Run()

	if err != nil {
		return err
	}

	// Add to steam
	err = steam.AddToShotcuts(&steam.Shortcut{
		AppName:       "Moonlight Game Streaming",
		Exe:           "/usr/bin/flatpak",
		StartDir:      "/usr/bin/",
		ShortcutPath:  "/var/lib/flatpak/exports/share/applications/com.moonlight_stream.Moonlight.desktop",
		LaunchOptions: "run --branch=stable --arch=x86_64 --command=moonlight com.moonlight_stream.Moonlight",
		IconURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/ef8051ce270059a142fcb0b3e47b1cd4.png",
		LogoURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/beb5ad322e679d0a6045c6cfc56e8b92.png",
		CoverURL:      "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/030d60c36d51783da9e4cbb6aa5abd2c.png",
		BannerURL:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/8a8f67cacf3e3d2d63614f515a2079b8.png",
		HeroURL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/0afefa2281c2f8b0b86d6332e2cdbe7d.png",
	})

	return err
}
