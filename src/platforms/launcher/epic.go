package launcher

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging/proton"
	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// Epic manifest struct
// Each game contains the manifest JSON file on game directory
// Declares only relevant fields to reduce struct
// File: $FOLDER/$KEY.item
type EpicManifest struct {
	AppName          string `json:"AppName"`
	CatalogNamespace string `json:"CatalogNamespace"`
	CatalogItemId    string `json:"CatalogItemId"`
	DisplayName      string `json:"DisplayName"`
	InstallLocation  string `json:"InstallLocation"`
	LaunchExecutable string `json:"LaunchExecutable"`
}

// Epic Games platform definition
func EpicPlatform() *Platform {

	name := "Epic Games"
	launcher := "Epic Games Launcher"
	programID := "epic-games"
	tags := []string{"Epic Games"}

	// On Linux, runs with Proton
	useProton := cli.IsLinux()
	if useProton {
		tags = append(tags, "Proton")
	}

	// ReadManifest file content and return as struct
	readManifest := func(path string) (*EpicManifest, error) {

		manifest := &EpicManifest{}
		err := fs.ReadJSON(path, manifest)
		if err != nil {
			return manifest, err
		}

		return manifest, nil
	}

	// Scan function to retrieve installed games from launcher
	scan := func(options *Options) ([]*Game, error) {

		results := []*Game{}
		manifestsPath := ""

		// Get the program to retrieve its executable and arguments
		program, err := programs.GetProgramByID(programID)
		if err != nil {
			return results, err
		}

		// Proton runtime requires special handling for manifests path
		if useProton {
			protonPackage := program.Package.(*proton.Proton)
			manifestsPath = protonPackage.RealPath("$PROGRAMDATA/Epic/EpicGamesLauncher/Data/Manifests")
		} else if cli.IsMacOS() {
			manifestsPath = fs.ExpandPath("$HOME/Library/Application Support/Epic/EpicGamesLauncher/Data/Manifests")
		} else if cli.IsWindows() {
			manifestsPath = fs.ExpandPath("$PROGRAMDATA/Epic/EpicGamesLauncher/Data/Manifests")
		} else {
			return results, nil
		}

		// Check if the manifests directory exists
		exist, err := fs.DirectoryExist(manifestsPath)
		if err != nil {
			return results, err
		} else if !exist {
			return results, nil
		}

		// Manifest files have the extension ".item", so we only process those files
		entries, err := filepath.Glob(filepath.Join(manifestsPath, "*.item"))
		if err != nil {
			return results, err
		} else if len(entries) == 0 {
			return results, err
		}

		// Iterate through the entries and process each one
		for _, entry := range entries {

			// Read manifest content
			manifest, err := readManifest(entry)
			if err != nil {
				return results, err
			}

			// Skip when no info detected
			if manifest.AppName == "" || manifest.LaunchExecutable == "" {
				continue
			}

			// Determine executable and launch options
			launchOptions := ""
			startDirectory := manifest.InstallLocation
			executable := strings.TrimSpace(fmt.Sprintf(
				"com.epicgames.launcher://apps/%s:%s:%s?action=launch&silent=true",
				manifest.CatalogNamespace,
				manifest.CatalogItemId,
				manifest.AppName,
			))

			// Proton must use the wrapper as executable
			if useProton {
				launchOptions = executable
				executable = program.Package.Executable()
				startDirectory = filepath.Dir(executable)
			}

			// Add game to resulting list
			game := &Game{
				Reference:      fmt.Sprintf("%s:%s", programID, manifest.AppName),
				Name:           manifest.DisplayName,
				Description:    fmt.Sprintf("Installed via %s", launcher),
				Platform:       name,
				Launcher:       launcher,
				Program:        programID,
				StartDirectory: cli.Quote(startDirectory),
				Executable:     cli.Quote(executable),
				LaunchOptions:  launchOptions,
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
