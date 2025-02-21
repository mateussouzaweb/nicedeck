package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/programs/assets"
)

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
