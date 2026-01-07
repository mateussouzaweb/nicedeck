package windows

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"unicode/utf16"
)

// Decode UTF-16 (LE or BE) bytes into string
func DecodeUTF16(data []byte, byteOrder binary.ByteOrder) string {
	if len(data)%2 != 0 {
		data = data[:len(data)-1]
	}

	u16 := make([]uint16, len(data)/2)
	for i := 0; i < len(u16); i++ {
		u16[i] = byteOrder.Uint16(data[i*2 : i*2+2])
	}

	return string(utf16.Decode(u16))
}

// Extracts path from .lnk file
func ReadLnkPath(path string) (string, error) {

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer func() {
		errors.Join(err, file.Close())
	}()

	// Read raw file contents
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Read BOM header
	reader := bufio.NewReader(file)
	bom := make([]byte, 2)
	_, err = reader.Read(bom)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("error reading BOM: %w", err)
	}

	// Decode bytes into string based on encoded type
	decoded := ""
	if bom[0] == 0xFE && bom[1] == 0xFF {
		decoded = DecodeUTF16(data, binary.BigEndian)
	} else if bom[0] == 0xFF && bom[1] == 0xFE {
		decoded = DecodeUTF16(data, binary.LittleEndian)
	} else {
		decoded = string(data)
	}

	// Search for Windows-style paths (e.g. C:\..., D:\...)
	pattern := regexp.MustCompile(`[A-Z]:\\[^\x00\r\n]{3,}`)
	match := pattern.FindString(decoded)
	if match == "" {
		return "", fmt.Errorf("no Windows path found in .lnk file")
	}

	return match, nil
}
