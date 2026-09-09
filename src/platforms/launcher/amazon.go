package launcher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging/proton"
	"github.com/mateussouzaweb/nicedeck/src/programs"
)

// Amazon product data struct
// Each game contains the product data file on game directory
// File: __InstallData__/$FOLDER/product_data.json
type AmazonProductData struct {
	ProductId        string `json:"ProductId"`
	InstalledVersion string `json:"InstalledVersion"`
	TargetVersion    string `json:"TargetVersion"`
	State            string `json:"State"`
}

// Amazon fuel command struct
type AmazonFuelCommand struct {
	Command string   `json:"Command"`
	Args    []string `json:"Args"`
}

// Amazon fuel struct
// Each game contains the launch details file on game directory
// File: $FOLDER/fuel.json
type AmazonFuel struct {
	SchemaVersion string               `json:"SchemaVersion"`
	PostInstall   []*AmazonFuelCommand `json:"PostInstall"`
	Main          *AmazonFuelCommand   `json:"Main"`
}

// Amazon Games platform definition
func AmazonPlatform() *Platform {

	name := "Amazon Games"
	launcher := "Amazon Games"
	programID := "amazon-games"
	tags := []string{"Amazon Games"}

	// On Linux, runs with Proton
	useProton := cli.IsLinux()
	if useProton {
		tags = append(tags, "Proton")
	}

	// ReadProductData file content and return as struct
	readProductData := func(path string) (*AmazonProductData, error) {

		data := &AmazonProductData{}
		err := fs.ReadJSON(path, data)
		if err != nil {
			return data, err
		}

		return data, nil
	}

	// readFuel file content and return as struct
	readFuel := func(path string) (*AmazonFuel, error) {

		fuel := &AmazonFuel{
			PostInstall: make([]*AmazonFuelCommand, 0),
			Main:        &AmazonFuelCommand{},
		}

		err := fs.ReadJSON(path, fuel)
		if err != nil {
			return fuel, err
		}

		return fuel, nil
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
			libraryPath = protonPackage.RealPath("$HOMEDRIVE/Amazon Games/Library")
		} else if cli.IsWindows() {
			libraryPath = fs.ExpandPath("$HOMEDRIVE/Amazon Games/Library")
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

			// Skip __InstallData__ folder
			if entry.Name() == "__InstallData__" {
				continue
			}

			// Read fuel content
			fuelFile := filepath.Join(gamePath, "fuel.json")
			fuel, err := readFuel(fuelFile)
			if err != nil {
				return results, err
			}

			// Skip when no info detected
			if fuel.Main.Command == "" {
				continue
			}

			// Read product data content
			productDataFile := filepath.Join(libraryPath, "__InstallData__", entry.Name(), "product_data.json")
			productData, err := readProductData(productDataFile)
			if err != nil {
				return results, err
			}

			// Skip when no info detected
			if productData.ProductId == "" {
				continue
			}

			// Determine executable and launch options
			// Use direct game launch without Amazon wrapper
			// Alternative option: amazon-games://play/${productId}
			executable := filepath.Join(gamePath, fuel.Main.Command)
			launchOptions := strings.Join(fuel.Main.Args, " ")
			startDirectory := gamePath

			// Proton must use the wrapper as executable
			if useProton {
				launchOptions = fmt.Sprintf("%s %s", executable, launchOptions)
				executable = program.Package.Executable()
				startDirectory = filepath.Dir(executable)
			}

			// Add game to resulting list
			game := &Game{
				Reference:      fmt.Sprintf("%s:%s", programID, productData.ProductId),
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
		Name:     name,
		Launcher: launcher,
		Program:  programID,
		Scan:     scan,
	}
}
