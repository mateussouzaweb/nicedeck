package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Installer for Chiaki
func Chiaki() *Program {
	return &Program{
		ID:          "chiaki",
		Name:        "Chiaki",
		Description: "Client for PlayStation Remote Play",
		Category:    "Streaming",
		Tags:        []string{"Gaming", "Streaming"},
		IconURL:     "https://cdn2.steamgriddb.com/icon/3af6a013ffca8b7e22ce57d00090b754.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/65602dccbf69f8ef8aafdd4ad7b43bd4.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/99979b287fe7f91ba35ff69b8fd14233.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/5c3867d9390d85c6e708a01196d288f4.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/f2b08f23d02d5fff247a41982d44f02e.png",
		Package: packaging.Available(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "re.chiaki.Chiaki",
		}),
	}
}

// Installer for GeForce NOW
func GeForceNow() *Program {
	return &Program{
		ID:          "geforce-now",
		Name:        "GeForce NOW",
		Description: "Client for GeForce Now",
		Category:    "Streaming",
		Tags:        []string{"Gaming", "Streaming"},
		IconURL:     "https://cdn2.steamgriddb.com/icon/3632435cf99eec2a53ee7e4d8eeab451.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/ee1c568adf7b9181213c80f9e917dd1f.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/acc90c264f09d151c7a09da4c06877e8.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/8cd586dd25cd66b50db63e51b5f44dcd.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/5e7e6e76699ea804c65b0c37974c660c.jpg",
		Package: packaging.Available(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "com.google.Chrome",
			Overrides: []string{"--filesystem=/run/udev:ro"},
			Arguments: []string{
				"--window-size=1024,640",
				"--force-device-scale-factor=1.25",
				"--device-scale-factor=1.25",
				"--app=https://play.geforcenow.com",
			},
		}),
	}
}

// Installer for Moonlight Game Streaming
func MoonlightGameStreaming() *Program {
	return &Program{
		ID:          "moonlight",
		Name:        "Moonlight Game Streaming",
		Description: "Play your PC games remotely",
		Category:    "Streaming",
		Tags:        []string{"Gaming", "Streaming"},
		IconURL:     "https://cdn2.steamgriddb.com/icon/ef8051ce270059a142fcb0b3e47b1cd4.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/beb5ad322e679d0a6045c6cfc56e8b92.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/030d60c36d51783da9e4cbb6aa5abd2c.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/8a8f67cacf3e3d2d63614f515a2079b8.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/0afefa2281c2f8b0b86d6332e2cdbe7d.png",
		Package: packaging.Available(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "com.moonlight_stream.Moonlight",
		}),
	}
}

// Installer for Xbox Cloud Gaming
func XboxCloudGaming() *Program {
	return &Program{
		ID:          "xbox-cloud-gaming",
		Name:        "Xbox Cloud Gaming",
		Description: "Client for Xbox Cloud Gaming",
		Category:    "Streaming",
		Tags:        []string{"Gaming", "Streaming"},
		IconURL:     "https://cdn2.steamgriddb.com/icon/164f545c22e17e5e9298b1c84b9e3e1e.png",
		LogoURL:     "https://cdn2.steamgriddb.com/logo/e3667b435e999b653dba291634579db1.png",
		CoverURL:    "https://cdn2.steamgriddb.com/grid/8a0657375c4d4024a7d9d5cc84b3c490.png",
		BannerURL:   "https://cdn2.steamgriddb.com/grid/2b16dcbe37a15a4932affb27447d7e21.png",
		HeroURL:     "https://cdn2.steamgriddb.com/hero/f6ba16107e08c04fc684308ab18d207a.png",
		Package: packaging.Available(&packaging.Flatpak{
			Namespace: "system",
			AppID:     "com.microsoft.Edge",
			Overrides: []string{"--filesystem=/run/udev:ro"},
			Arguments: []string{
				"--window-size=1024,640",
				"--force-device-scale-factor=1.25",
				"--device-scale-factor=1.25",
				"--app=https://www.xbox.com/play",
			},
		}),
	}
}
