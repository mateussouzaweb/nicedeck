package steam

import (
	"fmt"
	"hash/crc32"
)

// Generate preliminerId from given key.
// Extracted from https://github.com/SteamGridDB/steam-rom-manager
func GeneratePreliminarId(key string) uint {

	crc32q := crc32.MakeTable(crc32.IEEE)
	top := crc32.Checksum([]byte(key), crc32q) | 0x80000000
	result := uint(top)<<32 | 0x02000000

	return result
}

// Generate appId - Used for big picture grids
func GenerateAppId(key string) string {
	preliminar := GeneratePreliminarId(key)
	return fmt.Sprintf("%v", preliminar)
}

// Generate shortcutId - Used as appId in shortcuts.vdf
func GenerateShortcutId(key string) string {
	preliminar := GeneratePreliminarId(key)
	result := (preliminar >> 32) - 0x100000000
	return fmt.Sprintf("%v", result)
}
