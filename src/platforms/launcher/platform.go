package launcher

import "github.com/mateussouzaweb/nicedeck/src/platforms/launcher/game"

// Scanner function type
type Scanner func() ([]*game.Game, error)

// Platform struct
type Platform struct {
	Name     string  `json:"name"`
	Launcher string  `json:"launcher"`
	Program  string  `json:"program"`
	Scan     Scanner `json:"-"`
}

// Retrieve launcher platform specs
func GetPlatforms() ([]*Platform, error) {
	platforms := []*Platform{}
	return platforms, nil
}
