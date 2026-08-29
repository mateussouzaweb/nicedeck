package windows

import (
	"bytes"
	"encoding/binary"
	"image/png"
	"os"
)

// IconDir represents the 6-byte ICO file header
type IconDir struct {
	Reserved uint16 `json:"reserved"`
	Type     uint16 `json:"type"`
	Count    uint16 `json:"count"`
}

// IconDirEntry represents the 16-byte image directory entry
type IconDirEntry struct {
	Width       uint8  `json:"width"`
	Height      uint8  `json:"height"`
	ColorCount  uint8  `json:"colorCount"`
	Reserved    uint8  `json:"reserved"`
	Planes      uint16 `json:"planes"`
	BitCount    uint16 `json:"bitCount"`
	BytesInRes  uint32 `json:"bytesInRes"`
	ImageOffset uint32 `json:"imageOffset"`
}

// Convert PNG icon to ICO format for Windows shortcuts
func ConvertPngToIco(source string, destination string) error {

	// Read .png file data
	pngFile, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	pngReader := bytes.NewReader(pngFile)
	config, err := png.DecodeConfig(pngReader)
	if err != nil {
		return err
	}

	// Create destination .ico file
	iconFile, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer iconFile.Close()

	// Write the ICO Header (6 bytes)
	iconDir := IconDir{
		Reserved: 0,
		Type:     1, // 1 for ICO
		Count:    1, // 1 image
	}

	err = binary.Write(iconFile, binary.LittleEndian, iconDir)
	if err != nil {
		return err
	}

	// Write the ICO Image Entry (16 bytes)
	width := config.Width
	height := config.Height

	if width >= 256 {
		width = 0
	}
	if height >= 256 {
		height = 0
	}

	iconEntry := IconDirEntry{
		Width:       uint8(width),
		Height:      uint8(height),
		ColorCount:  0,
		Reserved:    0,
		Planes:      1,
		BitCount:    32,
		BytesInRes:  uint32(len(pngFile)),
		ImageOffset: 22, // Header (6) + Entry (16)
	}

	err = binary.Write(iconFile, binary.LittleEndian, iconEntry)
	if err != nil {
		return err
	}

	// Write the raw PNG bytes directly
	_, err = iconFile.Write(pngFile)
	if err != nil {
		return err
	}

	return nil
}
