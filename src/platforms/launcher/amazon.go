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

// Amazon stores per game product data on __InstallData__ folder
// File is named as product_data.json
type AmazonProductData struct {
	ProductId        string `json:"ProductId"`
	InstalledVersion string `json:"InstalledVersion"`
	TargetVersion    string `json:"TargetVersion"`
	State            string `json:"State"`
}

// Each amazon game has a fuel.json file inside the game directory
// Amazon fuel.json file struct is just a JSON file
type AmazonFuelMain struct {
	Command string   `json:"Command"`
	Args    []string `json:"Args"`
}
type AmazonFuel struct {
	SchemaVersion string          `json:"SchemaVersion"`
	PostInstall   []string        `json:"PostInstall"`
	Main          *AmazonFuelMain `json:"Main"`
}

// Amazon Games platform definition
func AmazonPlatform() *Platform {

	name := "Amazon"
	launcher := "Amazon Games"
	programID := "amazon-games"
	tags := []string{"Amazon"}

	// On Linux, runs with Proton
	if cli.IsLinux() {
		tags = append(tags, "Proton")
	}

	// ReadProductData file content and return as struct
	readProductData := func(path string) (*AmazonProductData, error) {

		data := &AmazonProductData{}
		err := fs.ReadJSON(path, data)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	// readFuel file content and return as struct
	readFuel := func(path string) (*AmazonFuel, error) {

		fuel := &AmazonFuel{
			Main: &AmazonFuelMain{},
		}

		err := fs.ReadJSON(path, fuel)
		if err != nil {
			return nil, err
		}

		return fuel, nil
	}

	// Scan function to retrieve installed games from Amazon Games
	scan := func() ([]*Game, error) {

		results := []*Game{}
		libraryPath := ""

		// Get the Amazon Games program to retrieve its executable and arguments
		program, err := programs.GetProgramByID(programID)
		if err != nil {
			return results, err
		}

		// Linux (Proton) runtime requires special handling for games path
		if cli.IsLinux() {
			protonPackage := program.Package.(*proton.Proton)
			libraryPath = protonPackage.RealPath("$HOMEDRIVE/Amazon Games/Library")
		} else if cli.IsWindows() {
			libraryPath = fs.ExpandPath("$HOMEDRIVE/Amazon Games/Library")
		} else if cli.IsMacOS() {
			// There is no MacOS application yet
			// libraryPath = fs.ExpandPath("$HOME/Library/Application Support/Amazon Games/Library")
			return results, nil
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

			// Skip __InstallData__ folder
			if entry.Name() == "__InstallData__" {
				continue
			}

			gamePath := filepath.Join(libraryPath, entry.Name())
			info, err := entry.Info()
			if err != nil {
				return results, err
			} else if !info.IsDir() {
				continue
			}

			// Read fuel content
			fuelFile := filepath.Join(gamePath, "fuel.json")
			fuel, err := readFuel(fuelFile)
			if err != nil {
				return results, err
			}

			// Skip when no fuel detected
			if fuel.Main.Command == "" {
				continue
			}

			// Read product data
			productDataFile := filepath.Join(libraryPath, "__InstallData__", entry.Name(), "product_data.json")
			productData, err := readProductData(productDataFile)
			if err != nil {
				return results, err
			}

			// Skip when no product data detected
			if productData.ProductId == "" {
				continue
			}

			// Determine executable and launch options
			// Use direct game launch without Amazon wrapper
			// Alternative option: amazon-games://play/${productId}
			executable := filepath.Join(gamePath, fuel.Main.Command)
			launchOptions := strings.Join(fuel.Main.Args, " ")

			// Linux (Proton) must use the proton as executable
			if cli.IsLinux() {
				launchOptions = fmt.Sprintf("%s %s", executable, launchOptions)
				executable = program.Package.Executable()
			}

			game := &Game{
				Reference:      fmt.Sprintf("%s:%s", programID, productData.ProductId),
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
