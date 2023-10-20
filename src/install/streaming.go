package install

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/shortcuts"
)

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
		Tags:          []string{"Utilities", "Streaming"},
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
		Tags:          []string{"Gaming", "Streaming"},
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
