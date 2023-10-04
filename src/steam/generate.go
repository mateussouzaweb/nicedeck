package steam

import (
	"hash/crc32"
)

// Generate preliminarId from given key.
// Extracted from https://github.com/SteamGridDB/steam-rom-manager
func GeneratePreliminarID(key string) uint {
	crc32q := crc32.MakeTable(crc32.IEEE)
	top := crc32.Checksum([]byte(key), crc32q) | 0x80000000
	result := uint(top)<<32 | 0x02000000

	return result
}

// Generate appId - Used for big picture grids.
// Extracted from https://github.com/SteamGridDB/steam-rom-manager
func GenerateAppID(key string) uint {
	return GeneratePreliminarID(key)
}

// Generate shortcutId - Used as appId in shortcuts.vdf.
// Extracted from https://github.com/SteamGridDB/steam-rom-manager
func GenerateShortcutID(key string) uint {
	preliminar := GeneratePreliminarID(key)
	result := (preliminar >> 32) - 0x100000000
	return result
}
