package install

// Installer for Chiaki
func Chiaki() *Program {
	return &Program{
		ID:           "chiaki",
		Name:         "Chiaki",
		Description:  "Client for PlayStation Remote Play",
		Tags:         []string{"Gaming", "Streaming"},
		FlatpakAppID: "re.chiaki.Chiaki",
		IconURL:      "https://cdn2.steamgriddb.com/icon/3af6a013ffca8b7e22ce57d00090b754.png",
		LogoURL:      "https://cdn2.steamgriddb.com/logo/65602dccbf69f8ef8aafdd4ad7b43bd4.png",
		CoverURL:     "https://cdn2.steamgriddb.com/grid/99979b287fe7f91ba35ff69b8fd14233.png",
		BannerURL:    "https://cdn2.steamgriddb.com/grid/5c3867d9390d85c6e708a01196d288f4.png",
		HeroURL:      "https://cdn2.steamgriddb.com/hero/f2b08f23d02d5fff247a41982d44f02e.png",
	}
}

// Installer for Jellyfin Media Player
func JellyfinMediaPlayer() *Program {
	return &Program{
		ID:           "jellyfin",
		Name:         "Jellyfin Media Player",
		Description:  "Client for Jellyfin Server",
		Tags:         []string{"Utilities", "Streaming"},
		FlatpakAppID: "com.github.iwalton3.jellyfin-media-player",
		IconURL:      "https://cdn2.steamgriddb.com/icon/bbe2977a4c5b136df752894d93b44c72.png",
		LogoURL:      "https://cdn2.steamgriddb.com/logo/c84389bbba219be3e13b80f9376a0db7.png",
		CoverURL:     "https://cdn2.steamgriddb.com/grid/43174d6e1d2f2791af4925de27e813af.png",
		BannerURL:    "https://cdn2.steamgriddb.com/grid/717a3aabd77351296bbf24f7274a4d6e.png",
		HeroURL:      "https://cdn2.steamgriddb.com/hero/5e38c6c14e095dd7b30db8c0fdba643a.png",
	}
}

// Installer for Moonlight Game Streaming
func MoonlightGameStreaming() *Program {
	return &Program{
		ID:           "moonlight",
		Name:         "Moonlight Game Streaming",
		Description:  "Play your PC games remotely",
		Tags:         []string{"Gaming", "Streaming"},
		FlatpakAppID: "com.moonlight_stream.Moonlight",
		IconURL:      "https://cdn2.steamgriddb.com/icon/ef8051ce270059a142fcb0b3e47b1cd4.png",
		LogoURL:      "https://cdn2.steamgriddb.com/logo/beb5ad322e679d0a6045c6cfc56e8b92.png",
		CoverURL:     "https://cdn2.steamgriddb.com/grid/030d60c36d51783da9e4cbb6aa5abd2c.png",
		BannerURL:    "https://cdn2.steamgriddb.com/grid/8a8f67cacf3e3d2d63614f515a2079b8.png",
		HeroURL:      "https://cdn2.steamgriddb.com/hero/0afefa2281c2f8b0b86d6332e2cdbe7d.png",
	}
}
