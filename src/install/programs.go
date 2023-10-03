package install

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/steam"
)

// Install Bottles
func Bottles() error {

	err := cli.Command(`
		flatpak install -y flathub com.usebottles.bottles
	`).Run()

	if err != nil {
		return err
	}

	// TODO: appID
	err = steam.DownloadArtworks(
		"4210646725",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/449ef87e4d3fa1f1f268196b185627dd.ico",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/92491efa7cda6552f740334c9e601855.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/8845e5d69c0f8a1d4b30334afb030214.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/123a00ca793f7db5b771574116bc061f.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/84bdc10b5cc3b036ce04a562b0e54d61.png",
	)

	return err
}

// Install Firefox
func Firefox() error {

	err := cli.Command(`
		flatpak install -y flathub org.mozilla.firefox
	`).Run()

	if err != nil {
		return err
	}

	err = steam.DownloadArtworks(
		"3384410319",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/f968fdc88852a4a3a27a81fe3f57bfc5.ico",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/43285a8b542fcdc35377439e05dcb04f.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/4529f985441a035ae4a107b8862ba4dd.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/9384fe92aef7ea0128be2c916ed07cea.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/a318166b8539611449bf21ddc297a783.png",
	)

	return err
}

// Install Google Chrome
func GoogleChrome() error {

	err := cli.Command(`
		flatpak install -y flathub com.google.Chrome
	`).Run()

	if err != nil {
		return err
	}

	err = steam.DownloadArtworks(
		"4210646725",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/3941c4358616274ac2436eacf67fae05.ico",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/3b049d0f6cbf5421d399f156807b8657.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/d45c26607db83f6f14b09dd70123913b.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/d40c243072a2d2957b3484e775f1f925.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/cae83cfcb1d8a2a4bb17bd1446fb1cee.png",
	)

	return err
}

// Install Heroic Games Launcher
func HeroicGamesLauncher() error {

	err := cli.Command(`
		flatpak install -y flathub com.heroicgameslauncher.hgl
	`).Run()

	if err != nil {
		return err
	}

	err = steam.DownloadArtworks(
		"2728092030",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/ae852ba7ae75fa4c5c7d186a61fcce92.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/6eebc030d78d41b6cbcf9067aeda9198.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/2b1c6cedeaf9571589e3dc9d51ba20e5.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/94e8e64cdefe77dcc168855c54f14acd.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/bee5ca2551bf346f067a3ac16057bc40.png",
	)

	return err
}

// Install Jellyfin Media Player
func JellyfinMediaPlayer() error {

	err := cli.Command(`
		flatpak install -y flathub com.github.iwalton3.jellyfin-media-player
	`).Run()

	if err != nil {
		return err
	}

	err = steam.DownloadArtworks(
		"3372214690",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/bbe2977a4c5b136df752894d93b44c72.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/c84389bbba219be3e13b80f9376a0db7.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/43174d6e1d2f2791af4925de27e813af.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/717a3aabd77351296bbf24f7274a4d6e.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/5e38c6c14e095dd7b30db8c0fdba643a.png",
	)

	return err
}

// Install Lutris
func Lutris() error {

	err := cli.Command(`
		flatpak install -y flathub net.lutris.Lutris
	`).Run()

	if err != nil {
		return err
	}

	// TODO: appID
	err = steam.DownloadArtworks(
		"4210646725",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/8d060abe1e38ab179742bd3af495f407.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/bbd451c375fb5b293a9b1f082bf8d024.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3b0d861c2cf5ed4d7b139ee277c8a04a.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/3c5bf5a314017c84acae32394125cf26.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/3b7f06487067b9aa2393a438dd095edc.png",
	)

	return err
}

// Install Moonlight
func Moonlight() error {

	err := cli.Command(`
		flatpak install -y flathub com.moonlight_stream.Moonlight
	`).Run()

	if err != nil {
		return err
	}

	err = steam.DownloadArtworks(
		"2258966675",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/ef8051ce270059a142fcb0b3e47b1cd4.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/beb5ad322e679d0a6045c6cfc56e8b92.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/030d60c36d51783da9e4cbb6aa5abd2c.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/8a8f67cacf3e3d2d63614f515a2079b8.png",
		"https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/0afefa2281c2f8b0b86d6332e2cdbe7d.png",
	)

	return err
}
