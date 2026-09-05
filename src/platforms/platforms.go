package platforms

import (
	"github.com/mateussouzaweb/nicedeck/src/platforms/console"
	"github.com/mateussouzaweb/nicedeck/src/platforms/launcher"
	"github.com/mateussouzaweb/nicedeck/src/platforms/native"
)

// Get console platforms specs
func GetConsoles() ([]*console.Platform, error) {
	return console.GetPlatforms()
}

// Get launcher platforms specs
func GetLaunchers() ([]*launcher.Platform, error) {
	return launcher.GetPlatforms()
}

// Get system native platforms specs
func GetSystemNative() ([]*native.Platform, error) {
	return native.GetPlatforms()
}
