package vdf

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// Write VDF result to buffer and return it
func WriteVdf(data Vdf) (*bytes.Buffer, error) {

	buffer := bytes.NewBuffer([]byte{})

	// Iterate over keys
	for key, value := range data {
		switch value := value.(type) {
		case uint:

			// Add type
			err := buffer.WriteByte(markerNumber)
			if err != nil {
				return buffer, err
			}

			// Add key
			err = writeString(buffer, key)
			if err != nil {
				return buffer, err
			}

			// Add value
			err = writeNumber(buffer, value)
			if err != nil {
				return buffer, err
			}

		case string:

			// Add type
			err := buffer.WriteByte(markerString)
			if err != nil {
				return buffer, err
			}

			// Add key
			err = writeString(buffer, key)
			if err != nil {
				return buffer, err
			}

			// Add value
			err = writeString(buffer, value)
			if err != nil {
				return buffer, err
			}

		case Vdf:

			// Add type
			err := buffer.WriteByte(markerMap)
			if err != nil {
				return buffer, err
			}

			// Add key
			err = writeString(buffer, key)
			if err != nil {
				return buffer, err
			}

			// Add value
			mapBuffer, err := WriteVdf(value)
			if err != nil {
				return nil, err
			}

			_, err = buffer.Write(mapBuffer.Bytes())
			if err != nil {
				return nil, err
			}

		default:
			return nil, fmt.Errorf("unrecognized type: %v", value)
		}
	}

	// Add map end
	err := buffer.WriteByte(markerEndOfMap)
	if err != nil {
		return buffer, err
	}

	return buffer, nil
}

// Write a buffer with a string value and null terminator
func writeString(buffer *bytes.Buffer, value string) error {

	// Convert to bytes
	bytes := []byte(value)

	// Ensure no nulls's in the string
	for _, v := range bytes {
		if v == 0 {
			return errors.New("NUL terminator found in key")
		}
	}

	// Append null terminator
	bytes = append(bytes, 0)

	// Append to buffer
	_, err := buffer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

// Write a buffer with a string value and null terminator
func writeNumber(buffer *bytes.Buffer, value uint) error {

	// Convert to bytes
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(value))

	// Append to buffer
	_, err := buffer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
