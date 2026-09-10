package launcher

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging/proton"
	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// GOG game playtask struct
type GOGGamePlayTask struct {
	Category  string `json:"category"`
	IsPrimary bool   `json:"isPrimary"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Type      string `json:"type"`
}

// GOG game info struct
// Each game contains the game info file on game directory
// Declares only relevant fields to reduce struct
// File: $FOLDER/goggame-*.info
type GOGGameInfo struct {
	RootGameId string             `json:"rootGameId"`
	GameID     string             `json:"gameId"`
	Name       string             `json:"name"`
	PlayTasks  []*GOGGamePlayTask `json:"playtasks"`
}

// GOG platform definition
func GOGPlatform() (*Platform, error) {

	name := "GOG"
	launcher := "GOG Galaxy"
	programID := "gog-galaxy"
	tags := []string{"GOG"}

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

	// ReadGameInfo file content and return as struct
	readGameInfo := func(path string) (*GOGGameInfo, error) {

		info := &GOGGameInfo{
			PlayTasks: make([]*GOGGamePlayTask, 0),
		}

		err := fs.ReadJSON(path, info)
		if err != nil {
			return info, err
		}

		return info, nil
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
			libraryPath = protonPackage.RealPath("$PROGRAMS/GOG Galaxy/Games")
		} else if cli.IsMacOS() {
			libraryPath = fs.ExpandPath("$HOME/Library/Application Support/GOG Galaxy/Games")
		} else if cli.IsWindows() {
			libraryPath = fs.ExpandPath("$PROGRAMS/GOG Galaxy/Games")
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

			// Find game info files in folder
			globPattern := filepath.Join(gamePath, "goggame-*.info")
			files, err := filepath.Glob(globPattern)
			if err != nil {
				return results, err
			} else if len(files) == 0 {
				continue
			}

			// Read the game info content
			// Only the first file result is relevant
			gameFile := files[0]
			gameInfo, err := readGameInfo(gameFile)
			if err != nil {
				return results, err
			}

			// Skip when no info detected
			if gameInfo.GameID == "" {
				continue
			}

			// Determine executable and launch options
			executable := ""
			launchOptions := ""
			startDirectory := gamePath

			// Executable is read from main playTask
			for _, playTask := range gameInfo.PlayTasks {
				if playTask.IsPrimary && playTask.Path != "" {
					executable = filepath.Join(gamePath, playTask.Path)
					break
				}
			}

			// Skip when executable could not be found
			if executable == "" {
				continue
			}

			// Proton must use the wrapper as executable
			if useProton {
				launchOptions = executable
				executable = program.Package.Executable()
				startDirectory = filepath.Dir(executable)
			}

			// Add game to resulting list
			game := &Game{
				Reference:      fmt.Sprintf("%s:%s", programID, gameInfo.GameID),
				Name:           gameInfo.Name,
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
