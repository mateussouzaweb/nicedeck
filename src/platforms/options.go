package platforms

import (
	"regexp"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Options struct
type Options struct {
	Platforms       []string         `json:"platforms"`
	Preferences     []string         `json:"preferences"`
	ExcludePaths    []string         `json:"excludePaths"`
	ExcludePatterns []*regexp.Regexp `json:"excludePatterns"`
}

// Check if path should be excluded based on options
func (o *Options) ShouldExclude(path string) bool {

	// Check if path contains given term
	for _, pattern := range o.ExcludePaths {
		if strings.Contains(strings.ToLower(path), strings.ToLower(pattern)) {
			return true
		}
	}

	// Check against regex exclusion list
	for _, pattern := range o.ExcludePatterns {
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
	}

	// Files on these folders will be ignored
	options.ExcludePaths = []string{
		fs.NormalizePath("/Updates/"), // Updates folder
		fs.NormalizePath("/Mods/"),    // Mods folder
		fs.NormalizePath("/DLCs/"),    // DLCs folder
		fs.NormalizePath("/Others/"),  // Folder with games to ignore
		fs.NormalizePath("/Ignore/"),  // Folder with games to ignore
		fs.NormalizePath("/Hide/"),    // Folder with games to ignore
	}

	// Files with these name patterns will be ignored
	options.ExcludePatterns = []*regexp.Regexp{
		regexp.MustCompile("(?i)Disc 0?[2-9]"),     // Disc 02 - 09 of some games
		regexp.MustCompile("(?i)Disc [1-9][0-9]"),  // Disc 10 - 99 of some games
		regexp.MustCompile("(?i)Track 0?[1-9]"),    // Track 01 - 09 of some games
		regexp.MustCompile("(?i)Track [1-9][0-9]"), // Track 10 - 99 of some games
	}

	return &options
}
