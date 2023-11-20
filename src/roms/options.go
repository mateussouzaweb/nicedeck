package roms

import (
	"slices"
)

// Options struct
type Options struct {
	Platforms  []string `json:"platforms"`
	Rebuild    bool     `json:"rebuild"`
	UseRyujinx bool     `json:"useRyujinx"`
}

// Transform values into valid options
func ToOptions(platforms []string, preferences []string, rebuild bool) *Options {

	options := Options{
		Platforms:  platforms,
		Rebuild:    rebuild,
		UseRyujinx: false,
	}

	if len(preferences) > 0 {
		if slices.Contains(preferences, "use-ryujinx") {
			options.UseRyujinx = true
		}
	}

	return &options
}
