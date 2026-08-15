package platforms

import (
	"github.com/mateussouzaweb/nicedeck/src/platforms/console"
	"github.com/mateussouzaweb/nicedeck/src/platforms/native"
)

// Get console platforms specs
func GetConsoles() ([]*console.Platform, error) {
	return console.GetPlatforms()
}

// Retrieve system native platform specs
func GetSystemNative() ([]*native.Platform, error) {
	return native.GetPlatforms()
}
