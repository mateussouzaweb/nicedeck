package launcher

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging/proton"
	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// Amazon stores per game install data on __Installer folder
// File is named as installerdata.xml
// Declares only relevant fields to reduce struct
type EAInstallerDataGameTitle struct {
	Locale string `xml:"locale,attr"`
	Value  string `xml:",chardata"`
}
type EAInstallerData struct {
	XMLName    xml.Name                    `xml:"DiPManifest"`
	Version    string                      `xml:"version,attr"`
	ContentIDs []string                    `xml:"contentIDs>contentID"`
	GameTitles []*EAInstallerDataGameTitle `xml:"gameTitles>gameTitle"`
}

// EA Games platform definition
func EAPlatform() *Platform {

	name := "EA Games"
	launcher := "EA App"
	programID := "ea-app"
	tags := []string{"EA Games"}

	// On Linux, runs with Proton
	useProton := cli.IsLinux()
	if useProton {
		tags = append(tags, "Proton")
	}

	// ReadInstallerData file content and return as struct
	readInstallerData := func(path string) (*EAInstallerData, error) {

		data := &EAInstallerData{
			ContentIDs: []string{},
			GameTitles: make([]*EAInstallerDataGameTitle, 0),
		}
		err := fs.ReadXML(path, data)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	// Scan function to retrieve installed games from EA App
	scan := func() ([]*Game, error) {

		results := []*Game{}
		libraryPath := ""

		// Get the EA App program to retrieve its executable and arguments
		program, err := programs.GetProgramByID(programID)
		if err != nil {
			return results, err
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

		// Iterate through the sub folders and process each one
		for _, entry := range entries {

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

			// Skip when no content ID detected
			if len(installerData.ContentIDs) == 0 {
				continue
			}

			// Read game ID
			// First entry is the main game, secondary are DLCs
			gameID := installerData.ContentIDs[0]

			// Determine executable and launch options
			executable := fmt.Sprintf("origin2://game/launch?offerIds=%s&autoDownload=1&authCode=&cmdParams=", gameID)
			launchOptions := ""

			// Proton must use the wrapper as executable
			if useProton {
				launchOptions = fmt.Sprintf("%s %s", executable, launchOptions)
				executable = program.Package.Executable()
			}

			game := &Game{
				Reference:      fmt.Sprintf("%s:%s", programID, gameID),
				Name:           entry.Name(),
				Description:    fmt.Sprintf("Installed via %s", launcher),
				Platform:       name,
				Launcher:       launcher,
				Program:        programID,
				StartDirectory: cli.Quote(filepath.Dir(executable)),
				Executable:     cli.Quote(executable),
				LaunchOptions:  cli.Quote(launchOptions),
				Tags:           tags,
			}

			results = append(results, game)
		}

		return results, nil
	}

	return &Platform{
		Name:     name,
		Launcher: launcher,
		Program:  programID,
		Scan:     scan,
	}
}
