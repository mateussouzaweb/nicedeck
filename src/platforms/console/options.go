package console

import (
	"regexp"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Options struct
type Options struct {
	Platforms   []string         `json:"platforms"`
	Preferences []string         `json:"preferences"`
	RootPath    string           `json:"rootPath"`
	Exclude     []*regexp.Regexp `json:"exclude"`
}

// Check if path should be excluded based on options
func (o *Options) ShouldExclude(path string) bool {

	for _, pattern := range o.Exclude {
		if pattern.MatchString(path) {
			return true
		}
	}

	return false
}

// Transform values into valid options
func ToOptions(platforms []string, preferences []string) *Options {

	options := Options{
		Platforms:   platforms,
		Preferences: preferences,
		RootPath:    fs.ExpandPath("$ROMS"),
	}

	// Files with these name patterns will be ignored
	compile := regexp.MustCompile
	options.Exclude = []*regexp.Regexp{
		compile("(?i)[\\\\/]Updates[\\\\/]"), // Updates folder
		compile("(?i)[\\\\/]Mods[\\\\/]"),    // Mods folder
		compile("(?i)[\\\\/]DLCs[\\\\/]"),    // DLCs folder
		compile("(?i)[\\\\/]Others[\\\\/]"),  // Folder with games to ignore
		compile("(?i)[\\\\/]Ignore[\\\\/]"),  // Folder with games to ignore
		compile("(?i)[\\\\/]Hide[\\\\/]"),    // Folder with games to ignore
		compile("(?i)Disc 0?[2-9]"),          // Disc 02 - 09 of some games
		compile("(?i)Disc [1-9][0-9]"),       // Disc 10 - 99 of some games
		compile("(?i)Track 0?[1-9]"),         // Track 01 - 09 of some games
		compile("(?i)Track [1-9][0-9]"),      // Track 10 - 99 of some games
	}

	return &options
}
