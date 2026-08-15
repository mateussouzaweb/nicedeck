package steam

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
)

// Retrieve Steam package
func GetPackage() packaging.Package {
	return packaging.Best(&linux.Flatpak{
		AppID:     "com.valvesoftware.Steam",
		Namespace: "system",
		Overrides: []string{
			fmt.Sprintf("--filesystem=%s", fs.ExpandPath("$GAMES")),
			"--talk-name=org.freedesktop.Flatpak",
			"--system-talk-name=org.freedesktop.NetworkManager",
		},
		Arguments: packaging.NoArguments(),
	}, &linux.Flatpak{
		AppID:     "com.valvesoftware.Steam",
		Namespace: "user",
		Overrides: []string{
			fmt.Sprintf("--filesystem=%s", fs.ExpandPath("$GAMES")),
			"--talk-name=org.freedesktop.Flatpak",
			"--system-talk-name=org.freedesktop.NetworkManager",
		},
		Arguments: packaging.NoArguments(),
	}, &linux.Binary{
		AppID:     "steam",
		Launcher:  "/usr/bin/steam",
		Arguments: packaging.NoArguments(),
	}, &linux.Binary{
		AppID:     "steam",
		Launcher:  "/usr/games/steam",
		Arguments: packaging.NoArguments(),
	}, &macos.Homebrew{
		AppID:     "steam",
		Launcher:  "/Applications/Steam.app",
		Arguments: packaging.NoArguments(),
	}, &windows.WinGet{
		AppID:     "Valve.Steam",
		Launcher:  "$PROGRAMS_X86/Steam/Steam.exe",
		Arguments: packaging.NoArguments(),
	})
}
