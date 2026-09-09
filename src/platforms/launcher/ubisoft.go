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

// Ubisoft install state struct
// Each game contains the install state file on game directory
// Content is generated from Protobuf and must be decoded
// Declares only relevant fields to reduce struct
// File: $FOLDER/uplay_install.state
type UbisoftInstallState struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Ubisoft Games platform definition
func UbisoftPlatform() *Platform {

	name := "Ubisoft Games"
	launcher := "Ubisoft Connect"
	programID := "ubisoft-connect"
	tags := []string{"Ubisoft Games"}

	// On Linux, runs with Proton
	useProton := cli.IsLinux()
	if useProton {
		tags = append(tags, "Proton")
	}

	// ReadInstallState file content and return as struct
	readInstallState := func(path string) (*UbisoftInstallState, error) {

		data := &UbisoftInstallState{}
		content, err := fs.ReadProtoBuf(path)
		if err != nil {
			return data, err
		}

		// Protobuf decoding use generic keys
		// We match content based on decoded values
		if _, ok := content["28"]; ok {
			data.ID = content["28"]
		}
		if _, ok := content["22"]; ok {
			data.Title = content["22"]
		}

		return data, nil
	}

	// Scan function to retrieve installed games from launcher
	scan := func() ([]*Game, error) {

		results := []*Game{}
		libraryPath := ""

		// Get the program to retrieve its executable and arguments
		program, err := programs.GetProgramByID(programID)
		if err != nil {
			return results, err
		}

		// Proton runtime requires special handling for library path
		if useProton {
			protonPackage := program.Package.(*proton.Proton)
			libraryPath = protonPackage.RealPath("$PROGRAMS_X86/Ubisoft/Ubisoft Game Launcher/games")
		} else if cli.IsWindows() {
			libraryPath = fs.ExpandPath("$PROGRAMS_X86/Ubisoft/Ubisoft Game Launcher/games")
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

			// Read install state content
			installStateFile := filepath.Join(gamePath, "uplay_install.state")
			installState, err := readInstallState(installStateFile)
			if err != nil {
				return results, err
			}

			// Skip when no info detected
			if installState.ID == "" {
				continue
			}

			// Determine executable and launch options
			executable := fmt.Sprintf("uplay://launch/%s/0", installState.ID)
			launchOptions := ""

			// Proton must use the wrapper as executable
			if useProton {
				launchOptions = executable
				executable = program.Package.Executable()
			}

			// Add game to resulting list
			game := &Game{
				Reference:      fmt.Sprintf("%s:%s", programID, installState.ID),
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
