package shortcuts

import (
	"hash/crc32"
)

// Generate shortcutID - Used as appID in shortcuts.vdf
func GenerateShortcutID(exe string, appName string) uint {
	uniqueName := []byte(exe + appName)
	result := crc32.ChecksumIEEE(uniqueName) | 0x80000000
	return uint(result)
}

// Generate appID - Used for big picture grids
func GenerateAppID(exe string, appName string) uint {
	value := GenerateShortcutID(exe, appName)
	result := uint(value)<<32 | 0x02000000
	return result
}
