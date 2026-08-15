package platforms

import (
	"github.com/mateussouzaweb/nicedeck/src/platforms/console"
	"github.com/mateussouzaweb/nicedeck/src/platforms/native"
)

// Get console platforms specs
func GetConsoles() ([]*console.Platform, error) {
	platforms := []string{}
	preferences := []string{}
	consoleOptions := console.ToOptions(platforms, preferences)
	return console.GetPlatforms(consoleOptions)
}

// Retrieve system native platform specs
func GetSystemNative() ([]*native.Platform, error) {
	platforms := []string{}
	preferences := []string{}
	nativeOptions := native.ToOptions(platforms, preferences)
	return native.GetPlatforms(nativeOptions)
}
