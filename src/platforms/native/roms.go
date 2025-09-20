package native

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// ROM struct
type ROM struct {
	Path           string `json:"path"`
	Directory      string `json:"directory"`
	File           string `json:"file"`
	Extension      string `json:"extension"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Runtime        string `json:"runtime"`
	Program        string `json:"program"`
	StartDirectory string `json:"startDirectory"`
	Executable     string `json:"executable"`
	LaunchOptions  string `json:"launchOptions"`
}

// Find ROM details based on the given path
func ParseROM(path string, options *Options) (*ROM, error) {

	rom := &ROM{}

	// Parse basic ROM information
	directory := filepath.Dir(path)
	file := filepath.Base(path)
	extension := filepath.Ext(path)
	name := strings.TrimSuffix(file, extension)

	// Find available platforms
	platforms, err := GetPlatforms(options)
	if err != nil {
		return rom, err
	}

	// Replace placeholder with real data on given value
	replaceData := func(value string) string {
		value = strings.Replace(value, "${ROM}", cli.Quote(path), 1)
		value = strings.Replace(value, "${DIRECTORY}", cli.Quote(directory), 1)
		return value
	}

	// Find runtime based on file data
	for _, platform := range platforms {
		if !slices.Contains(platform.Extensions, extension) {
			continue
		}

		// Valid, fill ROM data with platform
		rom.Path = path
		rom.Directory = directory
		rom.File = file
		rom.Extension = extension
		rom.Name = name
		rom.Description = fmt.Sprintf("%s executable", platform.Runtime)
		rom.Runtime = platform.Runtime
		rom.Program = strings.ToLower(platform.Runtime)
		rom.StartDirectory = replaceData(platform.StartDirectory)
		rom.Executable = replaceData(platform.Executable)
		rom.LaunchOptions = replaceData(platform.LaunchOptions)
		break
	}

	return rom, nil
}
