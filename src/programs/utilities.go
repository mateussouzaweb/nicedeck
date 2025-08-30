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
		Flags:       []string{"--remove-only-shortcut"},
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
		Description: "Wine and Proton-based compatibility tools manager",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://github.com/Vysp3r/ProtonPlus",
		IconURL:     assets.Icon("1e3b2280c3c02e5e9a89259bdf6e887c.png"),
		LogoURL:     assets.Logo("14cf97275c0fa8b72799cf97ad3a8cea.png"),
		CoverURL:    assets.Cover("cc98bda0fd790ef327868e0848be1a1b.png"),
		BannerURL:   assets.Banner("217caa6f1839e19191655c0ed782754a.png"),
		HeroURL:     assets.Hero("76d89f46aab1d736a03ca408f0c5ef50.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.vysp3r.ProtonPlus",
		}),
	}
}
