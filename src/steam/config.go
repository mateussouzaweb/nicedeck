package steam

import (
	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/vdf"
)

type Config struct {
	ShortcutsFile string
	ArtworksPath  string
	Shortcuts     Shortcuts
}

var _config *Config

// Use given runtime config
func Use(config *Config) (func() error, error) {

	_config = config

	// Save updated content for the shortcuts file
	saveShortcuts := func() error {

		// Encode content to bytes
		content, err := vdf.Marshal(_config.Shortcuts)
		if err != nil {
			return err
		}

		// Write content to file
		err = cli.WriteFile(_config.ShortcutsFile, content, 0666)
		if err != nil {
			return err
		}

		return nil
	}

	// Check if file exist
	if !cli.ExistFile(_config.ShortcutsFile) {
		return saveShortcuts, nil
	}

	// Read file content
	content, err := cli.ReadFile(_config.ShortcutsFile)
	if err != nil {
		return saveShortcuts, err
	}

	// Map to struct
	err = vdf.Unmarshal(content, &_config.Shortcuts)
	if err != nil {
		return saveShortcuts, err
	}

	return saveShortcuts, nil
}
