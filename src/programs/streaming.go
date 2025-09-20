package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/programs/assets"
)

// Installer for Chiaki-NG
func ChiakiNG() *Program {
	return &Program{
		ID:          "chiaki-ng",
		Name:        "Chiaki-NG",
		Description: "Client for PlayStation Remote Play",
		Category:    "Streaming",
		Tags:        []string{"Gaming", "Streaming"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://streetpea.github.io/chiaki-ng/",
		IconURL:     assets.Icon("c58aa7403da471ad796cf64288404006.png"),
		LogoURL:     assets.Logo("fdd3817fb0cf38c24dd377286b1d7e41.png"),
		CoverURL:    assets.Cover("346cf5bb8dff3e90e2c4df81a83701cf.png"),
		BannerURL:   assets.Banner("9111ec4aae8cd54acf89f011eee3c164.png"),
		HeroURL:     assets.Hero("9884dfd73a2471545e0c3f8c14177a04.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "io.github.streetpea.Chiaki4deck",
			Arguments: packaging.NoArguments(),
		}, &macos.Homebrew{
			AppID:     "streetpea/streetpea/chiaki-ng",
			AppName:   "chiaki-ng.app",
			Arguments: packaging.NoArguments(),
		}, &windows.WinGet{
			AppID:     "StreetPea.chiaki-ng",
			AppExe:    "$PROGRAMS\\chiaki-ng\\chiaki-ng.exe",
			Arguments: packaging.NoArguments(),
		}),
	}
}

// Installer for GeForce NOW
func GeForceNow() *Program {
	flags := []string{}
	if cli.IsLinux() {
		flags = append(flags, "--remove-only-shortcut")
	}

	return &Program{
		ID:          "geforce-now",
		Name:        "GeForce NOW",
		Description: "Client for GeForce Now",
		Category:    "Streaming",
		Tags:        []string{"Gaming", "Streaming"},
		Flags:       flags,
		Folders:     []string{},
		Website:     "https://www.nvidia.com/geforce-now",
		IconURL:     assets.Icon("3632435cf99eec2a53ee7e4d8eeab451.png"),
		LogoURL:     assets.Logo("ee1c568adf7b9181213c80f9e917dd1f.png"),
		CoverURL:    assets.Cover("acc90c264f09d151c7a09da4c06877e8.png"),
		BannerURL:   assets.Banner("8cd586dd25cd66b50db63e51b5f44dcd.png"),
		HeroURL:     assets.Hero("5e7e6e76699ea804c65b0c37974c660c.jpg"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.google.Chrome",
			Overrides: []string{"--filesystem=/run/udev:ro"},
			Arguments: &packaging.Arguments{
				Install: []string{},
				Remove:  []string{},
				Shortcut: []string{
					"--window-size=1024,640",
					"--force-device-scale-factor=1.25",
					"--device-scale-factor=1.25",
					"--app=https://play.geforcenow.com",
				},
			},
		}, &macos.Homebrew{
			AppID:     "nvidia-geforce-now",
			AppName:   "NVIDIA GeForce NOW.app",
			Arguments: packaging.NoArguments(),
		}, &windows.WinGet{
			AppID:     "Nvidia.GeForceNow",
			AppExe:    "$APPDATA\\Local\\NVIDIA Corporation\\GeForceNOW\\CEF\\GeForceNOW.exe",
			Arguments: packaging.NoArguments(),
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
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://moonlight-stream.org",
		IconURL:     assets.Icon("ef8051ce270059a142fcb0b3e47b1cd4.png"),
		LogoURL:     assets.Logo("beb5ad322e679d0a6045c6cfc56e8b92.png"),
		CoverURL:    assets.Cover("030d60c36d51783da9e4cbb6aa5abd2c.png"),
		BannerURL:   assets.Banner("8a8f67cacf3e3d2d63614f515a2079b8.png"),
		HeroURL:     assets.Hero("0afefa2281c2f8b0b86d6332e2cdbe7d.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.moonlight_stream.Moonlight",
			Arguments: packaging.NoArguments(),
		}, &macos.Homebrew{
			AppID:     "moonlight",
			AppName:   "Moonlight.app",
			Arguments: packaging.NoArguments(),
		}, &windows.WinGet{
			AppID:     "MoonlightGameStreamingProject.Moonlight",
			AppExe:    "$PROGRAMS\\Moonlight Game Streaming\\Moonlight.exe",
			Arguments: packaging.NoArguments(),
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
		Flags:       []string{"--remove-only-shortcut"},
		Folders:     []string{},
		Website:     "https://www.xbox.com/cloud-gaming",
		IconURL:     assets.Icon("164f545c22e17e5e9298b1c84b9e3e1e.png"),
		LogoURL:     assets.Logo("e3667b435e999b653dba291634579db1.png"),
		CoverURL:    assets.Cover("8a0657375c4d4024a7d9d5cc84b3c490.png"),
		BannerURL:   assets.Banner("2b16dcbe37a15a4932affb27447d7e21.png"),
		HeroURL:     assets.Hero("f6ba16107e08c04fc684308ab18d207a.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.microsoft.Edge",
			Overrides: []string{"--filesystem=/run/udev:ro"},
			Arguments: &packaging.Arguments{
				Install: []string{},
				Remove:  []string{},
				Shortcut: []string{
					"--window-size=1024,640",
					"--force-device-scale-factor=1.25",
					"--device-scale-factor=1.25",
					"--app=https://www.xbox.com/play",
				},
			},
		}, &macos.Homebrew{
			AppID:   "microsoft-edge",
			AppName: "Microsoft Edge.app",
			Arguments: &packaging.Arguments{
				Install:  []string{},
				Remove:   []string{},
				Shortcut: []string{"--app=https://www.xbox.com/play"},
			},
		}, &windows.WinGet{
			AppID:  "Microsoft.Edge",
			AppExe: "$PROGRAMS_X86\\Microsoft\\Edge\\Application\\msedge.exe",
			Arguments: &packaging.Arguments{
				Install:  []string{},
				Remove:   []string{},
				Shortcut: []string{"--app=https://www.xbox.com/play"},
			},
		}),
	}
}
