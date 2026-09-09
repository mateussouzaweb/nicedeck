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

// Battle.net product database struct
// Each game contains the product database file on game directory
// Content is generated from Protobuf and must be decoded
// Declares only relevant fields to reduce struct
// File: $FOLDER/.product.db
type BattleNetProductDB struct {
	Code        string `json:"code"`
	Version     string `json:"version"`
	InstallPath string `json:"installPath"`
}

// Battle.net platform definition
func BattleNetPlatform() *Platform {

	name := "Battle.net"
	launcher := "Battle.net"
	programID := "battle-net"
	tags := []string{"Battle.net"}

	// On Linux, runs with Proton
	useProton := cli.IsLinux()
	if useProton {
		tags = append(tags, "Proton")
	}

	// ReadProductDB file content and return as struct
	readProductDB := func(path string) (*BattleNetProductDB, error) {

		data := &BattleNetProductDB{}
		content, err := fs.ReadProtoBuf(path)
		if err != nil {
			return data, err
		}

		// Protobuf decoding use generic keys
		// We match content based on decoded values
		if _, ok := content["1"]; ok {
			data.Code = content["1"]
		}
		if _, ok := content["4.1.7"]; ok {
			data.Version = content["4.1.7"]
		}
		if _, ok := content["3.1"]; ok {
			data.InstallPath = content["3.1"]
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
			libraryPath = protonPackage.RealPath("$PROGRAMS_X86")
		} else if cli.IsMacOS() {
			libraryPath = fs.ExpandPath("$HOME/Library/Application Support/Battle.net")
		} else if cli.IsWindows() {
			libraryPath = fs.ExpandPath("$PROGRAMS_X86")
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

			// Skip Battle.net folder
			if entry.Name() == "Battle.net" {
				continue
			}

			// Since library folder is just generic
			// We need to check for the .product.db presence on subfolder
			// Installed games always has the .product.db file
			productDBFile := filepath.Join(gamePath, ".product.db")
			exist, err := fs.FileExist(productDBFile)
			if err != nil {
				return results, err
			} else if !exist {
				continue
			}

			// Read manifest data content
			productDB, err := readProductDB(productDBFile)
			if err != nil {
				return results, err
			}

			// Skip when no info detected
			if productDB.Code == "" {
				continue
			}

			// Determine executable and launch options
			executable := fmt.Sprintf("battlenet://%s", productDB.Code)
			launchOptions := ""

			// Proton must use the wrapper as executable
			if useProton {
				launchOptions = executable
				executable = program.Package.Executable()
			}

			// Add game to resulting list
			game := &Game{
				Reference:      fmt.Sprintf("%s:%s", programID, productDB.Code),
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
