package launcher

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging/proton"
	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// EA installer data game title struct
type EAInstallerDataGameTitle struct {
	Locale string `xml:"locale,attr"`
	Value  string `xml:",chardata"`
}

// EA installer data struct
// Each game contains the installer data file on game directory
// Declares only relevant fields to reduce struct
// File: $FOLDER/__Installer/installerdata.xml
type EAInstallerData struct {
	XMLName    xml.Name                    `xml:"DiPManifest"`
	Version    string                      `xml:"version,attr"`
	ContentIDs []string                    `xml:"contentIDs>contentID"`
	GameTitles []*EAInstallerDataGameTitle `xml:"gameTitles>gameTitle"`
}

// EA Games platform definition
func EAPlatform() (*Platform, error) {

	name := "EA Games"
	launcher := "EA App"
	programID := "ea-app"
	tags := []string{"EA Games"}

	// On Linux, runs with Proton
	useProton := cli.IsLinux()
	if useProton {
		tags = append(tags, "Proton")
	}

	// Get the program details
	program, err := programs.GetProgramByID(programID)
	if err != nil {
		return &Platform{}, err
	}

	available := program.Package.Available()
	installed, err := program.Package.Installed()
	if err != nil {
		return &Platform{}, err
	}

	// ReadInstallerData file content and return as struct
	readInstallerData := func(path string) (*EAInstallerData, error) {

		data := &EAInstallerData{
			ContentIDs: make([]string, 0),
			GameTitles: make([]*EAInstallerDataGameTitle, 0),
		}

		err := fs.ReadXML(path, data)
		if err != nil {
			return data, err
		}

		return data, nil
	}

	// Scan function to retrieve installed games from launcher
	scan := func(options *Options) ([]*Game, error) {

		results := []*Game{}
		libraryPath := ""

		// Skip if program not available
		if !available {
			return results, nil
		}

		// Include missing option will add games even when the launcher is missing
		includeMissing := slices.Contains(options.Preferences, "include-missing")
		if !installed && !includeMissing {
			return results, nil
		}

		// Proton runtime requires special handling for library path
		if useProton {
			protonPackage := program.Package.(*proton.Proton)
			libraryPath = protonPackage.RealPath("$PROGRAMS/EA Games")
		} else if cli.IsMacOS() {
			libraryPath = fs.ExpandPath("$HOME/Library/Application Support/EA Games")
		} else if cli.IsWindows() {
			libraryPath = fs.ExpandPath("$PROGRAMS/EA Games")
		} else {
			return results, nil
		}

		// Check if the library directory exists
		exist, err := fs.DirectoryExist(libraryPath)
		if err != nil {
			return results, err
		} else if !exist {
			return results, nil
		}

		// Read the contents of the library directory
		entries, err := os.ReadDir(libraryPath)
		if err != nil {
			return results, err
		} else if len(entries) == 0 {
			return results, err
		}

		// Iterate through the entries and process each one
		for _, entry := range entries {

			// Accept only folders
			gamePath := filepath.Join(libraryPath, entry.Name())
			info, err := entry.Info()
			if err != nil {
				return results, err
			} else if !info.IsDir() {
				continue
			}

			// Read installer data content
			installerDataFile := filepath.Join(gamePath, "__Installer", "installerdata.xml")
			installerData, err := readInstallerData(installerDataFile)
			if err != nil {
				return results, err
			}

			// Skip when no info detected
			if len(installerData.ContentIDs) == 0 {
				continue
			}

			// Determine game ID
			// First entry is the main game, secondary are DLCs
			gameID := installerData.ContentIDs[0]

			// Determine executable and launch options
			executable := fmt.Sprintf("origin2://game/launch?offerIds=%s&autoDownload=1&authCode=&cmdParams=", gameID)
			startDirectory := gamePath
			launchOptions := ""

			// Proton must use the wrapper as executable
			if useProton {
				launchOptions = executable
				executable = program.Package.Executable()
				startDirectory = filepath.Dir(executable)
			}

			// Add game to resulting list
			game := &Game{
				Reference:      fmt.Sprintf("%s:%s", programID, gameID),
				Name:           entry.Name(),
				Description:    fmt.Sprintf("Installed via %s", launcher),
				Platform:       name,
				Launcher:       launcher,
				Program:        programID,
				StartDirectory: cli.Quote(startDirectory),
				Executable:     cli.Quote(executable),
				LaunchOptions:  cli.Quote(launchOptions),
				Tags:           tags,
			}

			results = append(results, game)
		}

		return results, nil
	}

	return &Platform{
		Name:      name,
		Launcher:  launcher,
		Program:   programID,
		Available: available,
		Installed: installed,
		Scan:      scan,
	}, nil
}
