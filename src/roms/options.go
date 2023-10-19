package roms

import (
	"slices"
	"strings"
)

// Options struct
type Options struct {
	Platforms  []string `json:"platforms"`
	Rebuild    bool     `json:"rebuild"`
	UseRyujinx bool     `json:"useRyujinx"`
}

// Transform values into valid options
func ToOptions(platforms string, preferences string, rebuild bool) *Options {

	options := Options{
		Rebuild:    rebuild,
		UseRyujinx: false,
	}

	if platforms != "" {
		options.Platforms = strings.Split(strings.ToUpper(platforms), ",")
	}

	if preferences != "" {
		keys := strings.Split(preferences, ",")
		if slices.Contains(keys, "use-ryujinx") {
			options.UseRyujinx = true
		}
	}

	return &options
}
