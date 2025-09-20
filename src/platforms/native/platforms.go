package native

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
)

// Platform struct
type Platform struct {
	Runtime        string
	Extensions     []string
	StartDirectory string
	Executable     string
	LaunchOptions  string
}

// Retrieve system platform specs
func GetPlatforms(options *Options) ([]*Platform, error) {

	platforms := []*Platform{}

	// Linux
	if cli.IsLinux() {
		platforms = append(platforms, &Platform{
			Runtime:        "Native",
			Extensions:     []string{".AppImage", ".desktop", ".sh"},
			StartDirectory: "${DIRECTORY}",
			Executable:     "${ROM}",
			LaunchOptions:  "",
		})
	}

	// Linux (Proton)
	proton := linux.Proton{
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
			Runtime:        "Proton",
			Extensions:     []string{".exe", ".msi", ".bat", ".cmd"},
			StartDirectory: proton.ProtonPath(),
			Executable:     proton.Executable(),
			LaunchOptions:  "${ROM}",
		})
	}

	// MacOS
	if cli.IsMacOS() {
		platforms = append(platforms, &Platform{
			Runtime:        "Native",
			Extensions:     []string{".app", ".sh"},
			StartDirectory: "${DIRECTORY}",
			Executable:     "${ROM}",
			LaunchOptions:  "",
		})
	}

	// Windows
	if cli.IsWindows() {
		platforms = append(platforms, &Platform{
			Runtime:        "Native",
			Extensions:     []string{".exe", ".msi", ".bat", ".cmd"},
			StartDirectory: "${DIRECTORY}",
			Executable:     "${ROM}",
			LaunchOptions:  "",
		})
	}

	return platforms, nil
}
