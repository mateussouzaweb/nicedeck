package programs

import (
	"github.com/mateussouzaweb/nicedeck/src/esde"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/proton"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
	"github.com/mateussouzaweb/nicedeck/src/programs/assets"
	"github.com/mateussouzaweb/nicedeck/src/programs/website"
)

// Installer for Amazon Games
func AmazonGames() *Program {
	return &Program{
		ID:          "amazon-games",
		Name:        "Amazon Games",
		Description: "Store and launcher for Amazon Games",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities", "Store"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://gaming.amazon.com",
		IconURL:     assets.Icon("6329d71f868e390b04af435ba2363554.png"),
		LogoURL:     assets.Logo("91996a56621122df0d9dfc51717d4f22.png"),
		CoverURL:    assets.Cover("bc5787c3784d39729b4d950eb5143cd2.png"),
		BannerURL:   assets.Banner("110c01f002b7848d931b406b6adee66c.png"),
		HeroURL:     assets.Hero("a21e85aedf84619520f0c5e30bd55042.png"),
		Package: packaging.Best(&proton.Proton{
			AppID:       "amazon-games",
			AppName:     "AmazonGames",
			Installer:   "C:/Downloads/AmazonGamesSetup.exe",
			Uninstaller: "C:/users/steamuser/AppData/Local/Amazon Games/App/Uninstall Amazon Games.exe",
			Launcher:    "C:/users/steamuser/AppData/Local/Amazon Games/App/Amazon Games.exe",
			Arguments:   packaging.NoArguments(),
			Source:      website.Link("https://download.amazongames.com/AmazonGamesSetup.exe"),
		}, &windows.WinGet{
			AppID:     "Amazon.Games",
			AppExe:    "$APPDATA/Local/Amazon Games/App/Amazon Games.exe",
			Arguments: packaging.NoArguments(),
		}),
	}
}

// Installer for Blizzard Battle.net
func BattleNet() *Program {
	return &Program{
		ID:          "battle-net",
		Name:        "Battle.net",
		Description: "Store and launcher for Blizzard Games",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities", "Store"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://us.shop.battle.net/",
		IconURL:     assets.Icon("e4bd04c70ceb80c8c6d544d75d235c8a.png"),
		LogoURL:     assets.Logo("c828b4e93bd75e8f7307dbddedea6480.png"),
		CoverURL:    assets.Cover("356c41d28e278e936b46739712043616.png"),
		BannerURL:   assets.Banner("95dc580680cdd8578951011c081121c4.png"),
		HeroURL:     assets.Hero("9f319422ca17b1082ea49820353f14ab.jpg"),
		Package: packaging.Best(&proton.Proton{
			AppID:       "battle-net",
			AppName:     "BattleNet",
			Installer:   "C:/Downloads/Battle.net-Setup.exe",
			Uninstaller: "C:/ProgramData/Battle.net/Agent/Blizzard Uninstaller.exe",
			Launcher:    "C:/Program Files (x86)/Battle.net/Battle.net.exe",
			Arguments:   packaging.NoArguments(),
			Source: website.Release(
				"https://download.battle.net/?product=bnetdesk", "",
				"https://downloader.battle.net/*os=win*version=Live",
			),
		}, &macos.Homebrew{
			AppID:     "battle-net",
			AppName:   "Battle.net.app",
			Arguments: packaging.NoArguments(),
		}, &windows.WinGet{
			AppID:  "Blizzard.BattleNet",
			AppExe: "$PROGRAMS_X86/Battle.net/Battle.net.exe",
			Arguments: &packaging.Arguments{
				Install: []string{
					"--locale=en-US",
					fs.ExpandPath("--location=\"$PROGRAMS_X86/Battle.net\""),
				},
				Remove:   []string{},
				Shortcut: []string{},
			},
		}),
	}
}

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
		HeroURL:     assets.Hero("c24f9ae141fa02c7fa1deea7e1149557.png"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.usebottles.bottles",
			Overrides: []string{fs.ExpandPath("--filesystem=$GAMES")},
			Arguments: packaging.NoArguments(),
		}),
	}
}

// Installer for EA App
func EAApp() *Program {
	return &Program{
		ID:          "ea-app",
		Name:        "EA App",
		Description: "Store and launcher for Electronic Arts Games",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities", "Store"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://www.ea.com/ea-app",
		IconURL:     assets.Icon("caa3d0f8578cca5d7e5daf5d1ffd1425.png"),
		LogoURL:     assets.Logo("bd81e7bf0ad319fff0a5830a22d12550.png"),
		CoverURL:    assets.Cover("67fce8ab05c7c0a28fa66b353e813cbd.png"),
		BannerURL:   assets.Banner("f1b499e8db3046ebec712209e22f830d.png"),
		HeroURL:     assets.Hero("6458ed5e1bb03b8da47c065c2f647b26.png"),
		Package: packaging.Best(&proton.Proton{
			AppID:       "ea-app",
			AppName:     "EA",
			Installer:   "C:/Downloads/EAappInstaller.exe",
			Uninstaller: "C:/Downloads/EAappInstaller.exe",
			Launcher:    "C:/Program Files/Electronic Arts/EA Desktop/EA Desktop/EADesktop.exe",
			Arguments: &packaging.Arguments{
				Install:  []string{"/quiet"},
				Remove:   []string{"/uninstall", "/quiet"},
				Shortcut: []string{},
			},
			Source: website.Release(
				"https://www.ea.com/ea-app", "",
				"https:*/EAappInstaller.exe",
			),
		}, &macos.Homebrew{
			AppID:     "ea",
			AppName:   "EA app.app",
			Arguments: packaging.NoArguments(),
		}, &windows.WinGet{
			AppID:     "ElectronicArts.EADesktop",
			AppExe:    "$PROGRAMS/Electronic Arts/EA Desktop/EA Desktop/EADesktop.exe",
			Arguments: packaging.NoArguments(),
		}),
	}
}

// Installer for Epic Games
func EpicGames() *Program {
	return &Program{
		ID:          "epic-games",
		Name:        "Epic Games",
		Description: "Store and launcher for Epic Games",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities", "Store"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://store.epicgames.com",
		IconURL:     assets.Icon("5945f39f9f61fe06bf2a8e2624462548.png"),
		LogoURL:     assets.Logo("8402e5776dfcbac2fe148bb2c35528fc.png"),
		CoverURL:    assets.Cover("67f56a2fe648cfdb82822bfdc360ef6a.png"),
		BannerURL:   assets.Banner("02d7e610ae675ae3be88626d18fa7999.png"),
		HeroURL:     assets.Hero("164fbf608021ece8933758ee2b28dd7d.png"),
		Package: packaging.Best(&proton.Proton{
			AppID:       "epic-games",
			AppName:     "EpicGames",
			Installer:   "C:/Downloads/EpicGamesLauncherInstaller.msi",
			Uninstaller: "C:/Downloads/EpicGamesLauncherInstaller.msi",
			Launcher:    "C:/Program Files (x86)/Epic Games/Launcher/Portal/Binaries/Win32/EpicGamesLauncher.exe",
			Arguments: &packaging.Arguments{
				Install:  []string{"-opengl"},
				Remove:   []string{"-opengl"},
				Shortcut: []string{"-opengl"},
			},
			Source: website.Link("https://launcher-public-service-prod06.ol.epicgames.com/launcher/api/installer/download/EpicGamesLauncherInstaller.msi"),
		}, &macos.Homebrew{
			AppID:     "epic-games",
			AppName:   "Epic Games Launcher.app",
			Arguments: packaging.NoArguments(),
		}, &windows.WinGet{
			AppID:     "EpicGames.EpicGamesLauncher",
			AppExe:    "$PROGRAMS_X86/Epic Games/Launcher/Portal/Binaries/Win32/EpicGamesLauncher.exe",
			Arguments: packaging.NoArguments(),
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
		Package:     esde.GetPackage(),
	}
}

// Installer for GOG Galaxy
func GOGGalaxy() *Program {
	return &Program{
		ID:          "gog-galaxy",
		Name:        "GOG Galaxy",
		Description: "Store and launcher for GOG Games",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities", "Store"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://www.gog.com/galaxy",
		IconURL:     assets.Icon("d0f5edad9ac19abed9e235c0fe0aa59f.png"),
		LogoURL:     assets.Logo("bb3eb7da5111a89afb18524c385b4ee2.png"),
		CoverURL:    assets.Cover("c3d13ca6a5797b92dcaf18529d9d795f.png"),
		BannerURL:   assets.Banner("5f77d1e72f72a5ea4cfd99b4a21e7fdd.png"),
		HeroURL:     assets.Hero("01ccb68a74dd1edfbccbd76d86dbd51f.png"),
		Package: packaging.Best(&proton.Proton{
			AppID:       "gog-galaxy",
			AppName:     "GOG",
			Installer:   "C:/Downloads/GOG_Galaxy_2.0.exe",
			Uninstaller: "C:/Program Files (x86)/GOG Galaxy/unins000.exe",
			Launcher:    "C:/Program Files (x86)/GOG Galaxy/GalaxyClient.exe",
			Arguments: &packaging.Arguments{
				Install:  []string{"/silent"},
				Remove:   []string{"/SILENT"},
				Shortcut: []string{},
			},
			Source: website.Release(
				"https://www.gog.com/galaxy", "",
				"https:*/download/GOG_Galaxy_2.0.exe",
			),
		}, &macos.Homebrew{
			AppID:     "gog-galaxy",
			AppName:   "GOG Galaxy.app",
			Arguments: packaging.NoArguments(),
		}, &windows.WinGet{
			AppID:     "GOG.Galaxy",
			AppExe:    "$PROGRAMS_X86/GOG Galaxy/GalaxyClient.exe",
			Arguments: packaging.NoArguments(),
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
		Tags:        []string{"Gaming", "Utilities", "Store"},
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
			Overrides: []string{fs.ExpandPath("--filesystem=$GAMES")},
			Arguments: packaging.NoArguments(),
		}, &macos.Homebrew{
			AppID:     "heroic",
			AppName:   "Heroic.app",
			Arguments: packaging.NoArguments(),
		}, &windows.WinGet{
			AppID:     "HeroicGamesLauncher.HeroicGamesLauncher",
			AppExe:    "$APPDATA/Local/Programs/heroic/Heroic.exe",
			Arguments: packaging.NoArguments(),
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
			Overrides: []string{fs.ExpandPath("--filesystem=$GAMES")},
			Arguments: packaging.NoArguments(),
		}),
	}
}

// Installer for Valve Steam
func Steam() *Program {
	return &Program{
		ID:          "steam",
		Name:        "Steam",
		Description: "Store and launcher for Steam Games",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities", "Store"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://store.steampowered.com",
		IconURL:     assets.Icon("d5510b2e793187ba82840ef35588cd10.png"),
		LogoURL:     assets.Logo("2472ed54742b7773cffc40910063839b.png"),
		CoverURL:    assets.Cover("c5174327c1975f78b7ffc788ed60b80e.png"),
		BannerURL:   assets.Banner("0e18441e60c88b9af7ebde5cdf65a23a.jpg"),
		HeroURL:     assets.Hero("63ca87b524b54b70a2bb83a5d20909c0.jpg"),
		Package: packaging.Best(&linux.Flatpak{
			Namespace: "system",
			AppID:     "com.valvesoftware.Steam",
			Overrides: []string{
				fs.ExpandPath("--filesystem=$GAMES"),
				"--talk-name=org.freedesktop.Flatpak",
			},
			Arguments: packaging.NoArguments(),
		}, &linux.Binary{
			AppID:     "steam",
			AppBin:    "/usr/bin/steam",
			Arguments: packaging.NoArguments(),
		}, &macos.Homebrew{
			AppID:     "steam",
			AppName:   "Steam.app",
			Arguments: packaging.NoArguments(),
		}, &windows.WinGet{
			AppID:     "Valve.Steam",
			AppExe:    "$PROGRAMS_X86/Steam/Steam.exe",
			Arguments: packaging.NoArguments(),
		}),
	}
}

// Installer for Ubisoft Connect
func UbisoftConnect() *Program {
	return &Program{
		ID:          "ubisoft-connect",
		Name:        "Ubisoft Connect",
		Description: "Store and launcher for Ubisoft Games",
		Category:    "Gaming",
		Tags:        []string{"Gaming", "Utilities", "Store"},
		Flags:       []string{},
		Folders:     []string{},
		Website:     "https://www.ubisoft.com/ubisoft-connect",
		IconURL:     assets.Icon("1f3435c7390fbc03bdc3621c420c7ebe.png"),
		LogoURL:     assets.Logo("021b8947656eb84e4c641506215777c8.png"),
		CoverURL:    assets.Cover("09d966b427fe08f5674b7e22a58bce8b.jpg"),
		BannerURL:   assets.Banner("5070c1f86e4885d73865919ce537fd21.png"),
		HeroURL:     assets.Hero("b1d49d65692f373bd3ae6ed4af9eda30.png"),
		Package: packaging.Best(&proton.Proton{
			AppID:       "ubisoft-connect",
			AppName:     "Ubisoft",
			Installer:   "C:/Downloads/UbisoftConnectInstaller.exe",
			Uninstaller: "C:/Program Files (x86)/Ubisoft/Ubisoft Game Launcher/Uninstall.exe",
			Launcher:    "C:/Program Files (x86)/Ubisoft/Ubisoft Game Launcher/UbisoftConnect.exe",
			Arguments: &packaging.Arguments{
				Install:  []string{"/S"},
				Remove:   []string{},
				Shortcut: []string{},
			},
			Source: website.Link("https://static3.cdn.ubi.com/orbit/launcher_installer/UbisoftConnectInstaller.exe"),
		}, &windows.WinGet{
			AppID:     "Ubisoft.Connect",
			AppExe:    "$PROGRAMS_X86/Ubisoft/Ubisoft Game Launcher/UbisoftConnect.exe",
			Arguments: packaging.NoArguments(),
		}),
	}
}
