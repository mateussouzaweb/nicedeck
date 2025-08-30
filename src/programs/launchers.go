package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/programs/assets"
	"github.com/mateussouzaweb/nicedeck/src/programs/esde"
)

// Installer for Bottles
func Bottles() *Program {
	return &Program{
		ID:          "bottles",
		Name:        "Bottles",
		Description: "Run Windows in a Bottle",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://usebottles.com",
		IconURL:     assets.Icon("7c5c040ae5d810d39deebbc55a06ff3f.png"),
		LogoURL:     assets.Logo("92491efa7cda6552f740334c9e601855.png"),
		CoverURL:    assets.Cover("8845e5d69c0f8a1d4b30334afb030214.png"),
		BannerURL:   assets.Banner("123a00ca793f7db5b771574116bc061f.png"),
		HeroURL:     assets.Hero("84bdc10b5cc3b036ce04a562b0e54d61.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.usebottles.bottles",
		}),
	}
}

// Installer for ES-DE
func ESDE() *Program {
	return &Program{
		ID:          "es-de",
		Name:        "ES-DE",
		Description: "Frontend for browsing and launching emulated games",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities"},
		Flags:       []string{},
		Folders:     []string{"$APPLICATIONS"},
		Website:     "https://es-de.org",
		IconURL:     assets.Icon("85ebad98d8a178be8baf16929526446e.png"),
		LogoURL:     assets.Logo("c3bb9214431dec7ca7d1ebcfeca73236.png"),
		CoverURL:    assets.Cover("21bd6ea21e43de6dc80e2bc8917f4ba3.png"),
		BannerURL:   assets.Banner("67a900732336f1ce9d0c0496352fa9ab.png"),
		HeroURL:     assets.Hero("9323f21f2098b7288267c785458548b2.png"),
		OnInstall:   esde.WriteSettings,
		Package: packaging.Best(&linux.AppImage{
			AppID:   "es-de",
			AppName: "$APPLICATIONS/ES-DE/ES-DE.AppImage",
			Source:  esde.Release("LinuxAppImage", "file"),
		}, &macos.Application{
			AppID:    "es-de",
			AppName:  "$APPLICATIONS/ES-DE/ES-DE.app",
			AppAlias: "$HOME/Applications/ES-DE.app",
			Source:   esde.Release("macOSApple", "dmg"),
		}, &windows.Executable{
			AppID:    "ES-DE",
			AppExe:   "$APPLICATIONS\\ES-DE\\ES-DE.exe",
			AppAlias: "$START_MENU\\ES-DE.lnk",
			Source:   esde.Release("WindowsPortable", "zip"),
		}),
	}
}

// Installer for Heroic Games Launcher
func HeroicGamesLauncher() *Program {
	return &Program{
		ID:          "heroic-games",
		Name:        "Heroic Games Launcher",
		Description: "Launcher for Epic Games, GOG and Prime Gaming",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://heroicgameslauncher.com",
		IconURL:     assets.Icon("ae852ba7ae75fa4c5c7d186a61fcce92.png"),
		LogoURL:     assets.Logo("6eebc030d78d41b6cbcf9067aeda9198.png"),
		CoverURL:    assets.Cover("2b1c6cedeaf9571589e3dc9d51ba20e5.png"),
		BannerURL:   assets.Banner("94e8e64cdefe77dcc168855c54f14acd.png"),
		HeroURL:     assets.Hero("bee5ca2551bf346f067a3ac16057bc40.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.heroicgameslauncher.hgl",
		}, &macos.Homebrew{
			AppID:   "heroic",
			AppName: "Heroic.app",
		}, &windows.WinGet{
			AppID:  "HeroicGamesLauncher.HeroicGamesLauncher",
			AppExe: "$APPDATA\\Local\\Programs\\heroic\\Heroic.exe",
		}),
	}
}

// Installer for Lutris
func Lutris() *Program {
	return &Program{
		ID:          "lutris",
		Name:        "Lutris",
		Description: "Play all your games on Linux",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://lutris.net",
		IconURL:     assets.Icon("8d060abe1e38ab179742bd3af495f407.png"),
		LogoURL:     assets.Logo("bbd451c375fb5b293a9b1f082bf8d024.png"),
		CoverURL:    assets.Cover("3b0d861c2cf5ed4d7b139ee277c8a04a.png"),
		BannerURL:   assets.Banner("3c5bf5a314017c84acae32394125cf26.png"),
		HeroURL:     assets.Hero("3b7f06487067b9aa2393a438dd095edc.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "net.lutris.Lutris",
		}),
	}
}
