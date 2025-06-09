package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/programs/assets"
	"github.com/mateussouzaweb/nicedeck/src/programs/github"
)

// Installer for NiceDeck
func NiceDeck() *Program {
	return &Program{
		ID:          "nicedeck",
		Name:        "NiceDeck",
		Description: "YES, self installer and updater",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities"},
		Folders:     []string{},
		Website:     "https://github.com/mateussouzaweb/nicedeck",
		IconURL:     assets.Icon("84ad88e9ffaeb60e8a2c83b6c108debd.png"),
		LogoURL:     assets.Logo("784ac91d0f3747ed26cd45781e9f20f3.png"),
		CoverURL:    assets.Cover("4113a973e756a0a1d9ca6653dcec0462.png"),
		BannerURL:   assets.Banner("b685f1ed7e9a7b8bbd3280104179cee3.png"),
		HeroURL:     assets.Hero("4f71036295e627bf8dff9e06d8602d06.png"),
		Package: packaging.Best(&linux.Binary{
			AppID:  "nicedeck",
			AppBin: "$APPLICATIONS/NiceDeck/nicedeck",
			Source: github.Release(
				"https://github.com/mateussouzaweb/nicedeck",
				"nicedeck-linux-amd64",
			),
		}, &macos.Application{
			AppID:    "nicedeck",
			AppName:  "$APPLICATIONS/NiceDeck/nicedeck",
			AppAlias: "$HOME/Applications/NiceDeck",
			Source: github.Release(
				"https://github.com/mateussouzaweb/nicedeck",
				"nicedeck-macos-arm64",
			),
		}, &windows.Executable{
			AppID:    "NiceDeck",
			AppExe:   "$APPLICATIONS\\NiceDeck\\nicedeck.exe",
			AppAlias: "$START_MENU\\NiceDeck.lnk",
			Source: github.Release(
				"https://github.com/mateussouzaweb/nicedeck",
				"nicedeck-windows-amd64.exe",
			),
		}),
	}
}

// Installer for ProtonPlus
func ProtonPlus() *Program {
	return &Program{
		ID:          "protonplus",
		Name:        "ProtonPlus",
		Description: "Wine and Proton-based compatiblity tools manager",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities"},
		Folders:     []string{},
		Website:     "https://github.com/Vysp3r/ProtonPlus",
		IconURL:     assets.Icon("fe13849d9b9437c5a61a1760ada2a5a6.png"),
		LogoURL:     assets.Logo("4d8c150eb82579842e2d5dc5faa07999.png"),
		CoverURL:    assets.Cover("7901f04bfecd29119dfcce1c708108b1.png"),
		BannerURL:   assets.Banner("f38705891f01bda4bd16551f42ff7c0a.png"),
		HeroURL:     assets.Hero("bc6f714aa3dfeef9320a838b79515c2d.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.vysp3r.ProtonPlus",
		}),
	}
}
