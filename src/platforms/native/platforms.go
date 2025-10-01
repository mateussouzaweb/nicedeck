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
	Extensions     string `json:"extensions"`
	StartDirectory string `json:"startDirectory"`
	Executable     string `json:"executable"`
	LaunchOptions  string `json:"launchOptions"`
}

// Retrieve system platform specs
func GetPlatforms(options *Options) ([]*Platform, error) {

	platforms := []*Platform{}

	// Linux
	if cli.IsLinux() {
		platforms = append(platforms, &Platform{
			Name:           "Linux",
			Runtime:        "Native",
			Extensions:     ".AppImage .desktop .sh",
			StartDirectory: "${DIRECTORY}",
			Executable:     "${ROM}",
			LaunchOptions:  "",
		})
	}

	// Linux (Proton)
	proton := proton.Proton{
		Arguments: packaging.NoArguments(),
		Source:    &packaging.Source{},
	}

	proton.Launcher = proton.Executable()
	protonInstalled, err := proton.Installed()
	if err != nil {
		return platforms, err
	}

	if proton.Available() && protonInstalled {
		platforms = append(platforms, &Platform{
			Name:           "Linux",
			Runtime:        "Proton",
			Extensions:     ".exe .msi .bat .cmd",
			StartDirectory: proton.ProtonPath(),
			Executable:     proton.Executable(),
			LaunchOptions:  "${ROM}",
		})
	}

	// MacOS
	if cli.IsMacOS() {
		platforms = append(platforms, &Platform{
			Name:           "MacOS",
			Runtime:        "Native",
			Extensions:     ".app .sh",
			StartDirectory: "${DIRECTORY}",
			Executable:     "${ROM}",
			LaunchOptions:  "",
		})
	}

	// Windows
	if cli.IsWindows() {
		platforms = append(platforms, &Platform{
			Name:           "Windows",
			Runtime:        "Native",
			Extensions:     ".exe .msi .bat .cmd",
			StartDirectory: "${DIRECTORY}",
			Executable:     "${ROM}",
			LaunchOptions:  "",
		})
	}

	return platforms, nil
}
