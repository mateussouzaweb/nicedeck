package vdf

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

const (
	typeObject byte = 0x00
	typeString byte = 0x01
	typeUint32 byte = 0x02
	typeEnd    byte = 0x08
)

// VdfBinary map
type VdfBinary map[string]any

// Return key value as string
func (vdf VdfBinary) String(key string) string {
	value, _ := vdf[key].(string)
	return value
}

// Return key value as uint32
func (vdf VdfBinary) Uint32(key string) uint32 {
	value, _ := vdf[key].(uint32)
	return value
}

// Return key value as nested object
func (vdf VdfBinary) Object(key string) VdfBinary {
	value, _ := vdf[key].(VdfBinary)
	return value
}

// ParseBinary reads a binary VDF object, including its terminating marker.
func ParseBinary(content io.Reader) (VdfBinary, error) {

	reader := bufio.NewReader(content)
	readString := func() (string, error) {
		value, err := reader.ReadString(0)
		if err != nil {
			return "", fmt.Errorf("unterminated string: %w", err)
		}
		return value[:len(value)-1], nil
	}

	var parseObject func() (VdfBinary, error)
	parseObject = func() (VdfBinary, error) {
		result := VdfBinary{}
		for {

			// Detect kind and end of file
			kind, err := reader.ReadByte()
			if err != nil {
				if err == io.EOF {
					err = io.ErrUnexpectedEOF
				}
				return nil, err
			}
			if kind == typeEnd {
				return result, nil
			}

			switch kind {
			case typeObject, typeString, typeUint32:
			default:
				return nil, fmt.Errorf("unsupported binary VDF type: 0x%02x", kind)
			}

			// Key always is string
			key, err := readString()
			if err != nil {
				return nil, fmt.Errorf("reading key: %w", err)
			}

			// Retrieve appropriated value based on byte type
			switch kind {
			case typeObject:
				value, err := parseObject()
				if err != nil {
					return nil, fmt.Errorf("reading object %q: %w", key, err)
				}
				result[key] = value
			case typeString:
				value, err := readString()
				if err != nil {
					return nil, fmt.Errorf("reading string %q: %w", key, err)
				}
				result[key] = value
			case typeUint32:
				var value uint32
				err := binary.Read(reader, binary.LittleEndian, &value)
				if err != nil {
					return nil, fmt.Errorf("reading number %q: %w", key, err)
				}
				result[key] = value
			}
		}
	}

	return parseObject()
}

// WriteBinary encodes a binary VDF object with its terminating marker.
func WriteBinary(data VdfBinary) (*bytes.Buffer, error) {

	buffer := &bytes.Buffer{}
	writeString := func(value string) error {
		if strings.IndexByte(value, 0) >= 0 {
			return fmt.Errorf("string contains a NUL byte")
		}
		buffer.WriteString(value)
		buffer.WriteByte(0)
		return nil
	}

	var writeObject func(VdfBinary) error
	writeObject = func(object VdfBinary) error {
		for key, value := range object {

			// Write kind
			var kind byte
			switch value.(type) {
			case VdfBinary:
				kind = typeObject
			case string:
				kind = typeString
			case uint32:
				kind = typeUint32
			default:
				return fmt.Errorf("unsupported value type for %q: %T", key, value)
			}

			buffer.WriteByte(kind)

			// Write key
			if err := writeString(key); err != nil {
				return fmt.Errorf("writing key %q: %w", key, err)
			}

			// Write value
			switch value := value.(type) {
			case VdfBinary:
				if err := writeObject(value); err != nil {
					return err
				}
			case string:
				if err := writeString(value); err != nil {
					return fmt.Errorf("writing string %q: %w", key, err)
				}
			case uint32:
				if err := binary.Write(buffer, binary.LittleEndian, value); err != nil {
					return fmt.Errorf("writing number %q: %w", key, err)
				}
			}
		}

		// Write map end
		buffer.WriteByte(typeEnd)
		return nil
	}

	if err := writeObject(data); err != nil {
		return nil, err
	}

	return buffer, nil
}
