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

// EpicGames stores games installed in manifest files
// Manifest file is name as "$NAME.item", but they are just like JSON files
// Declares only relevant fields to reduce struct
type EpicManifest struct {
	AppName          string `json:"AppName"`
	CatalogNamespace string `json:"CatalogNamespace"`
	CatalogItemId    string `json:"CatalogItemId"`
	DisplayName      string `json:"DisplayName"`
	InstallLocation  string `json:"InstallLocation"`
	LaunchExecutable string `json:"LaunchExecutable"`
}

// Epic Games Launcher platform definition
func EpicPlatform() *Platform {

	name := "Epic Games"
	launcher := "Epic Games Launcher"
	programID := "epic-games"

	// ReadManifest file content and return as Manifest struct
	readManifest := func(path string) (*EpicManifest, error) {

		manifest := &EpicManifest{}
		err := fs.ReadJSON(path, manifest)
		if err != nil {
			return nil, err
		}

		return manifest, nil
	}

	// Scan function to retrieve installed games from Epic Games Launcher
	scan := func() ([]*Game, error) {

		results := []*Game{}
		manifestsPath := ""

		// Get the Epic Games Launcher program to retrieve its executable and arguments
		program, err := programs.GetProgramByID(programID)
		if err != nil {
			return results, err
		}

		executable := program.Package.Executable()
		args := program.Package.Args()

		// MacOS must use the executable inside the app bundle
		if cli.IsMacOS() {
			executable = filepath.Join(executable, "Contents", "MacOS", "EpicGamesLauncher")
		}

		// Linux (Proton) runtime requires special handling for the manifest path
		if cli.IsLinux() {
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

		// Iterate through the manifest files and process each one
		for _, entry := range entries {

			// Read the manifest file and parse its content
			manifest, err := readManifest(entry)
			if err != nil {
				return results, err
			}

			// Skip if manifest is empty
			if manifest.AppName == "" || manifest.LaunchExecutable == "" {
				continue
			} else {
				cli.Debug("Detected: %s / %s\n", launcher, manifest.AppName)
			}

			// Create a new game object based on the manifest
			launchOptions := strings.TrimSpace(fmt.Sprintf(
				"%s com.epicgames.launcher://apps/%s:%s:%s?action=launch&silent=true",
				strings.Join(args, " "),
				manifest.CatalogNamespace,
				manifest.CatalogItemId,
				manifest.AppName,
			))

			game := &Game{
				Reference:      fmt.Sprintf("%s:%s", programID, manifest.AppName),
				Name:           manifest.DisplayName,
				Description:    fmt.Sprintf("Installed via %s", launcher),
				Platform:       name,
				Launcher:       launcher,
				Program:        programID,
				StartDirectory: cli.Quote(filepath.Dir(executable)),
				Executable:     cli.Quote(executable),
				LaunchOptions:  launchOptions,
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
