package vdf

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// Read VDF data in buffer
func readVdf(buffer *bytes.Buffer) (Vdf, error) {

	result := Vdf{}

	// Make sure buffer has data
	data := buffer.Bytes()
	if len(data) == 0 {
		return result, errors.New("empty data")
	}

	// Make sure is a valid VDF binary content
	first := data[0]
	if first != markerMap &&
		first != markerString &&
		first != markerNumber &&
		first != markerEndOfMap {
		return result, errors.New("unexpected format")
	}

	// We expect that buffer is like and map of key and values
	// So here we try to fill this map with key and values
	result, err := readNextMap(buffer)
	if err == io.EOF {
		return result, errors.New("reached the end of the file earlier than expected, data might be corrupted")
	}
	if err != nil {
		return result, err
	}

	return result, nil
}

// Read next map on buffer
func readNextMap(buffer *bytes.Buffer) (Vdf, error) {

	result := Vdf{}

	for {

		// Read next byte
		b, err := buffer.ReadByte()
		if err != nil {
			return result, err
		}

		// Byte is end of map
		if b == markerEndOfMap {
			break
		}

		// Key always is string
		key, err := readNextString(buffer)
		if err != nil {
			return result, err
		}

		// Retrieve appropriated value based on byte type
		var value interface{}
		switch b {
		case markerMap:
			value, err = readNextMap(buffer)
		case markerNumber:
			value, err = readNextNumber(buffer)
		case markerString:
			value, err = readNextString(buffer)
		default:
			err = fmt.Errorf("unexpected byte: 0x%02x, data might be corrupted", b)
		}

		if err != nil {
			return result, err
		}

		result[key] = value
	}

	return result, nil
}

// Read next string on buffer
func readNextString(buffer *bytes.Buffer) (string, error) {

	value, err := buffer.ReadString(markerEndOfString)
	if err != nil {
		return value, err
	}

	return value[:len(value)-1], nil // Removes the null terminator
}

// Read next number on buffer
func readNextNumber(buffer *bytes.Buffer) (uint32, error) {

	var value uint32

	bf := make([]byte, 4)
	length, err := buffer.Read(bf)
	if err != nil {
		return value, err
	}
	if length != len(bf) {
		return value, errors.New("invalid number")
	}

	value = binary.LittleEndian.Uint32(bf)
	return value, nil
}
