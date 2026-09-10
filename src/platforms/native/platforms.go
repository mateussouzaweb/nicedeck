package native

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/proton"
)

// Platform struct
type Platform struct {
	Name           string `json:"name"`
	Runtime        string `json:"runtime"`
	Available      bool   `json:"available"`
	Installed      bool   `json:"installed"`
	Extensions     string `json:"extensions"`
	StartDirectory string `json:"startDirectory"`
	Executable     string `json:"executable"`
	LaunchOptions  string `json:"launchOptions"`
}

// Retrieve system platform specs
func GetPlatforms() ([]*Platform, error) {

	platforms := []*Platform{}

	// Linux
	if cli.IsLinux() {
		platforms = append(platforms, &Platform{
			Name:           "Linux",
			Runtime:        "Native",
			Available:      true,
			Installed:      true,
			Extensions:     ".AppImage .desktop .sh",
			StartDirectory: "${DIRECTORY}",
			Executable:     "${PATH}",
			LaunchOptions:  "",
		})
	}

	// Linux (Proton)
	proton := proton.Proton{
		Arguments: packaging.NoArguments(),
		Source:    &packaging.Source{},
	}

	if proton.Available() {
		proton.Launcher = proton.Executable()
		protonInstalled, err := proton.Installed()
		if err != nil {
			return platforms, err
		}

		platforms = append(platforms, &Platform{
			Name:           "Linux",
			Runtime:        "Proton",
			Available:      true,
			Installed:      protonInstalled,
			Extensions:     ".exe .msi .bat .cmd",
			StartDirectory: proton.DataPath(),
			Executable:     proton.Executable(),
			LaunchOptions:  "${PATH}",
		})
	}

	// MacOS
	if cli.IsMacOS() {
		platforms = append(platforms, &Platform{
			Name:           "MacOS",
			Runtime:        "Native",
			Available:      true,
			Installed:      true,
			Extensions:     ".app .sh",
			StartDirectory: "${DIRECTORY}",
			Executable:     "${PATH}",
			LaunchOptions:  "",
		})
	}

	// Windows
	if cli.IsWindows() {
		platforms = append(platforms, &Platform{
			Name:           "Windows",
			Runtime:        "Native",
			Available:      true,
			Installed:      true,
			Extensions:     ".exe .msi .bat .cmd .lnk",
			StartDirectory: "${DIRECTORY}",
			Executable:     "${PATH}",
			LaunchOptions:  "",
		})
	}

	return platforms, nil
}
