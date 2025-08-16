package shortcuts

import (
	"fmt"
	"hash/crc32"
	"strconv"
)

// Convert string to uint
func ToUint(value string) uint {
	number, _ := strconv.ParseUint(value, 10, 64)
	return uint(number)
}

// Convert uint to string
func FromUint(value uint) string {
	return fmt.Sprintf("%d", value)
}

// Generate shortcut ID
// - Used as ID in library shortcut
// - Used as AppID in Steam shortcuts.vdf
func GenerateID(name string, executable string) string {
	uniqueName := []byte(executable + name)
	result := crc32.ChecksumIEEE(uniqueName) | 0x80000000
	return FromUint(uint(result))
}
