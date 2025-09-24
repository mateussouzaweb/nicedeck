package esde

import (
	"fmt"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/packaging/macos"
	"github.com/mateussouzaweb/nicedeck/src/packaging/windows"
)

// Get download URL from the latest release available
func GetDownloadURL(releaseType string) (string, error) {

	domain := "https://gitlab.com/es-de/emulationstation-de"
	endpoint := fmt.Sprintf("%s/-/raw/master/latest_release.json", domain)

	// Response struct
	var releases struct {
		Stable struct {
			Version  string `json:"version"`
			Packages []struct {
				Name     string `json:"name"`
				Filename string `json:"filename"`
				URL      string `json:"url"`
				Md5      string `json:"md5"`
			} `json:"packages"`
		} `json:"stable"`
	}

	// Request latest releases
	err := fs.RetrieveJSON(endpoint, &releases)
	if err != nil {
		return "", err
	}

	// Check for matching asset
	for _, release := range releases.Stable.Packages {
		if release.Name == releaseType {
			return release.URL, nil
		}
	}

	return "", fmt.Errorf("could not retrieve latest release")
}

// Return packaging source from release
func Release(releaseType string, format string) *packaging.Source {
	return &packaging.Source{
		Format: format,
		Resolver: func() (string, error) {
			return GetDownloadURL(releaseType)
		},
	}
}

// Retrieve ESDE package
func GetPackage() packaging.Package {
	return packaging.Best(&linux.AppImage{
		AppID:     "es-de",
		AppName:   "$APPLICATIONS/ES-DE/ES-DE.AppImage",
		Arguments: packaging.NoArguments(),
		Source:    Release("LinuxAppImage", "file"),
	}, &macos.Application{
		AppID:     "es-de",
		AppName:   "$APPLICATIONS/ES-DE/ES-DE.app",
		AppAlias:  "$HOME/Applications/ES-DE.app",
		Arguments: packaging.NoArguments(),
		Source:    Release("macOSApple", "dmg"),
	}, &windows.Executable{
		AppID:     "ES-DE",
		AppExe:    "$APPLICATIONS/ES-DE/ES-DE.exe",
		AppAlias:  "$START_MENU/ES-DE.lnk",
		Arguments: packaging.NoArguments(),
		Source:    Release("WindowsPortable", "zip"),
	})
}
