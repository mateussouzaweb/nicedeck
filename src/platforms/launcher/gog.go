package launcher

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging/proton"
	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// GOG stores game information in json files
// Info file is name as "goggame-*.info" inside the game folder
// Declares only relevant fields to reduce struct
type GOGGamePlayTask struct {
	Category  string `json:"category"`
	IsPrimary bool   `json:"isPrimary"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Type      string `json:"type"`
}

type GOGGameInfo struct {
	RootGameId string             `json:"rootGameId"`
	GameID     string             `json:"gameId"`
	Name       string             `json:"name"`
	PlayTasks  []*GOGGamePlayTask `json:"playtasks"`
}

// GOG platform definition
func GOGPlatform() *Platform {

	name := "GOG"
	launcher := "GOG Galaxy"
	programID := "gog-galaxy"

	// Read game info file content and return as info struct
	readGameInfo := func(path string) (*GOGGameInfo, error) {

		info := &GOGGameInfo{
			PlayTasks: make([]*GOGGamePlayTask, 0),
		}

		err := fs.ReadJSON(path, info)
		if err != nil {
			return nil, err
		}

		return info, nil
	}

	// Scan function to retrieve installed games from GOG
	scan := func() ([]*Game, error) {

		results := []*Game{}
		gamesPath := ""

		// Get the GOG program to retrieve its executable and arguments
		program, err := programs.GetProgramByID(programID)
		if err != nil {
			return results, err
		}

		// Linux (Proton) runtime requires special handling for games path
		if cli.IsLinux() {
			protonPackage := program.Package.(*proton.Proton)
			gamesPath = protonPackage.RealPath("$PROGRAMS/GOG Galaxy/Games")
		} else if cli.IsMacOS() {
			gamesPath = fs.ExpandPath("$HOME/Library/Application Support/GOG Galaxy/Games")
		} else if cli.IsWindows() {
			gamesPath = fs.ExpandPath("$PROGRAMS/GOG Galaxy/Games")
		} else {
			return results, nil
		}

		// Check if the games directory exists
		exist, err := fs.DirectoryExist(gamesPath)
		if err != nil {
			return results, err
		} else if !exist {
			return results, nil
		}

		// Read the contents of the games directory
		entries, err := os.ReadDir(gamesPath)
		if err != nil {
			return results, err
		} else if len(entries) == 0 {
			return results, err
		}

		// Iterate through the sub folders and process each one
		for _, entry := range entries {

			gamePath := filepath.Join(gamesPath, entry.Name())
			info, err := entry.Info()
			if err != nil {
				return results, err
			} else if !info.IsDir() {
				continue
			}

			// Find game info file in folder
			files, err := filepath.Glob(filepath.Join(gamePath, "goggame-*.info"))
			if err != nil {
				return results, err
			} else if len(files) == 0 {
				continue
			}

			for _, gameFile := range files {

				// Read the info file and parse its content
				gameInfo, err := readGameInfo(gameFile)
				if err != nil {
					return results, err
				}

				// Skip if info file is empty
				if gameInfo.GameID == "" {
					continue
				} else {
					cli.Debug("Detected: %s / %s\n", launcher, gameInfo.Name)
				}

				// Determine executable and launch options
				executable := ""
				launchOptions := ""
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

				// Linux (Proton) must use the proton as executable
				if cli.IsLinux() {
					launchOptions = executable
					executable = program.Package.Executable()
				}

				game := &Game{
					Reference:      fmt.Sprintf("%s:%s", programID, gameInfo.GameID),
					Name:           gameInfo.Name,
					Description:    fmt.Sprintf("Installed via %s", launcher),
					Platform:       name,
					Launcher:       launcher,
					Program:        programID,
					StartDirectory: cli.Quote(filepath.Dir(executable)),
					Executable:     cli.Quote(executable),
					LaunchOptions:  cli.Quote(launchOptions),
				}

				results = append(results, game)
				break
			}
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
