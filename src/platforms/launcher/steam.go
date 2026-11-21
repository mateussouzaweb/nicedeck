package launcher

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/programs"
	"github.com/mateussouzaweb/nicedeck/src/steam/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/vdf"
)

// Steam platform definition
func SteamPlatform() (*Platform, error) {

	name := "Steam"
	launcher := "Steam"
	programID := "steam"
	tags := []string{"Steam"}

	// List of excluded known tools and runtimes
	excludedApps := map[string]bool{
		"2230260": true, // Proton Next
		"1826330": true, // Proton EasyAntiCheat Runtime
		"1161040": true, // Proton BattlEye Runtime
		"1070560": true, // Steam Linux Runtime 1.0 (scout)
		"1391110": true, // Steam Linux Runtime 2.0 (soldier)
		"1628350": true, // Steam Linux Runtime 3.0 (sniper)
		"4183110": true, // Steam Linux Runtime 4.0
		"228980":  true, // Steamworks Common Redistributables
		"3086180": true, // Proton Voice Files
		"250820":  true, // SteamVR
		"858280":  true, // Proton 3.7
		"961940":  true, // Proton 3.16
		"1054830": true, // Proton 4.2
		"1113280": true, // Proton 4.11
		"1245040": true, // Proton 5.0
		"1420170": true, // Proton 5.13
		"1580130": true, // Proton 6.3
		"1887720": true, // Proton 7.0
		"2348590": true, // Proton 8.0
		"2805730": true, // Proton 9.0
		"1493710": true, // Proton Experimental
		"4427310": true, // Proton Experimental (ARM64)
		"2180100": true, // Proton Hotfix
		"3658110": true, // Proton 10.0
		"4628710": true, // Proton 11.0
		"4628740": true, // Proton 11.0 (ARM64)
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

	// Read Steam library and installation manifest file
	readManifest := func(path string) (vdf.Vdf, error) {
		data, err := vdf.ParseText(path)
		if err != nil {
			return data, fmt.Errorf("error while reading Steam manifest %s: %w", path, err)
		}

		return data, nil
	}

	// Scan installed apps
	scan := func(options *Options) ([]*Game, error) {

		results := []*Game{}

		if !available {
			return results, nil
		}

		includeMissing := slices.Contains(options.Preferences, "include-missing")
		if !installed && !includeMissing {
			return results, nil
		}

		basePath, err := steam.GetBasePath()
		if err != nil {
			return results, err
		} else if basePath == "" {
			return results, nil
		}

		// Read libraries locations from libraryfolders.vdf
		// Locations are nested on libraryfolders.[].path
		libraries := []string{}
		librariesFile := filepath.Join(basePath, "steamapps", "libraryfolders.vdf")
		librariesManifest, err := readManifest(librariesFile)
		if err != nil && !os.IsNotExist(err) {
			return results, err
		}

		for _, value := range librariesManifest.Object("libraryfolders") {
			library, ok := value.(vdf.Vdf)
			if !ok {
				continue
			}

			path := library.String("path")
			if path != "" && !slices.Contains(libraries, path) {
				libraries = append(libraries, path)
			}
		}

		slices.Sort(libraries)
		processed := map[string]bool{}

		for _, library := range libraries {

			// Find list of app manifests detected in the library
			pattern := filepath.Join(library, "steamapps", "appmanifest_*.acf")
			entries, err := filepath.Glob(pattern)
			if err != nil {
				return results, err
			} else if len(entries) == 0 {
				return results, err
			}

			// Process each manifest entry
			for _, entry := range entries {

				// Read the manifest file
				manifest, err := readManifest(entry)
				if os.IsNotExist(err) {
					continue
				} else if err != nil {
					return results, err
				}

				// Skip if excluded or processed
				app := manifest.Object("AppState")
				appID := app.String("appid")
				if processed[appID] || excludedApps[appID] {
					continue
				}

				// Skip if name is not detected
				appName := app.String("name")
				if appName == "" {
					continue
				}

				launchOptions := ""
				startDirectory := basePath
				executable := strings.TrimSpace(fmt.Sprintf("steam://rungameid/%s", appID))

				game := &Game{
					Reference:      fmt.Sprintf("%s:%s", programID, appID),
					Name:           appName,
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
				processed[appID] = true
			}
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
