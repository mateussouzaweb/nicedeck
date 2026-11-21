package vdf

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

const (
	typeObject byte = 0x00
	typeString byte = 0x01
	typeUint32 byte = 0x02
	typeEnd    byte = 0x08
)

// ParseBinary reads a binary VDF object, including its terminating marker
func ParseBinary(path string) (Vdf, error) {
	data := Vdf{}

	// Check if file exist
	exist, err := fs.FileExist(path)
	if err != nil {
		return data, err
	} else if !exist {
		return data, nil
	}

	// Read VDF file content
	cli.Debug("Reading VDF at %s\n", path)
	content, err := os.ReadFile(path)
	if err != nil {
		return data, err
	}

	// Read VDF to resulting map
	buffer := bytes.NewBuffer(content)
	reader := bufio.NewReader(buffer)
	readString := func() (string, error) {
		value, err := reader.ReadString(0)
		if err != nil {
			return "", fmt.Errorf("unterminated string error: %w", err)
		}
		return value[:len(value)-1], nil
	}

	var parseObject func() (Vdf, error)
	parseObject = func() (Vdf, error) {
		result := Vdf{}
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
				return nil, fmt.Errorf("error reading key: %w", err)
			}

			// Retrieve appropriated value based on byte type
			switch kind {
			case typeObject:
				value, err := parseObject()
				if err != nil {
					return nil, fmt.Errorf("error reading object %q: %w", key, err)
				}
				result[key] = value
			case typeString:
				value, err := readString()
				if err != nil {
					return nil, fmt.Errorf("error reading string %q: %w", key, err)
				}
				result[key] = value
			case typeUint32:
				var value uint32
				err := binary.Read(reader, binary.LittleEndian, &value)
				if err != nil {
					return nil, fmt.Errorf("error reading number %q: %w", key, err)
				}
				result[key] = value
			}
		}
	}

	return parseObject()
}

// WriteBinary encodes a binary VDF object with its terminating marker
func WriteBinary(path string, data Vdf) error {

	buffer := &bytes.Buffer{}
	writeString := func(value string) error {
		if strings.IndexByte(value, 0) >= 0 {
			return fmt.Errorf("string contains a NUL byte")
		}
		buffer.WriteString(value)
		buffer.WriteByte(0)
		return nil
	}

	var writeObject func(Vdf) error
	writeObject = func(object Vdf) error {
		for key, value := range object {

			// Write kind
			var kind byte
			switch value.(type) {
			case Vdf:
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
			case Vdf:
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

	// Transform VDF into bytes
	err := writeObject(data)
	if err != nil {
		return err
	}

	// Write content to VDF file
	cli.Debug("Writing VDF at %s\n", path)
	err = fs.WriteFile(path, buffer.String())
	if err != nil {
		return err
	}

	return nil
}
