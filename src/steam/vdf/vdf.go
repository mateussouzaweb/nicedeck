package vdf

import (
	"bytes"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"text/scanner"
)

// Text VDF tokens
const (
	textObject rune = '{'
	textString rune = scanner.String
	textEnd    rune = '}'
	textEOF    rune = scanner.EOF
)

// Vdf map (for VDF/ACF files)
type Vdf map[string]any

// Return key value as string
func (vdf Vdf) String(key string) string {
	value, _ := vdf[key].(string)
	return value
}

// Return key value as uint32, or zero if missing, invalid, or out of range
func (vdf Vdf) Uint32(key string) uint32 {
	value, err := strconv.ParseUint(vdf.String(key), 10, 32)
	if err != nil {
		return 0
	}
	return uint32(value)
}

// Return key value as nested object
func (vdf Vdf) Object(key string) Vdf {
	value, _ := vdf[key].(Vdf)
	return value
}

// Parse reads quoted text KeyValues, as used by Steam library and app manifests.
// Duplicate keys overwrite earlier values. Conditional and unquoted syntax is unsupported.
func Parse(content io.Reader) (Vdf, error) {

	var reader scanner.Scanner
	reader.Init(content)
	reader.Mode = scanner.ScanStrings | scanner.ScanComments | scanner.SkipComments

	var scanErr error
	reader.Error = func(s *scanner.Scanner, message string) {
		if scanErr == nil {
			scanErr = fmt.Errorf("%s: %s", s.Position, message)
		}
	}

	next := func() (rune, error) {
		token := reader.Scan()
		return token, scanErr
	}

	readString := func() (string, error) {
		value, err := strconv.Unquote(reader.TokenText())
		if err != nil {
			return "", fmt.Errorf("%s: invalid string: %w", reader.Position, err)
		}
		return value, nil
	}

	var parseObject func(bool) (Vdf, error)
	parseObject = func(nested bool) (Vdf, error) {
		result := Vdf{}
		for {
			token, err := next()
			if err != nil {
				return nil, err
			}

			switch token {
			case textEOF:
				if nested {
					return nil, fmt.Errorf("unexpected EOF: missing closing brace")
				}
				return result, nil
			case textEnd:
				if !nested {
					return nil, fmt.Errorf("%s: unexpected closing brace", reader.Position)
				}
				return result, nil
			case textString:
			default:
				return nil, fmt.Errorf("%s: expected quoted key", reader.Position)
			}

			key, err := readString()
			if err != nil {
				return nil, err
			}
			token, err = next()
			if err != nil {
				return nil, err
			}

			switch token {
			case textString:
				value, err := readString()
				if err != nil {
					return nil, err
				}
				result[key] = value
			case textObject:
				value, err := parseObject(true)
				if err != nil {
					return nil, err
				}
				result[key] = value
			default:
				return nil, fmt.Errorf("%s: expected string or object after %q", reader.Position, key)
			}
		}
	}

	return parseObject(false)
}

// Write encodes text VDF/ACF with sorted keys and tab indentation.
// Values must be strings or nested Vdf maps, matching Parse output.
func Write(data Vdf) (*bytes.Buffer, error) {

	buffer := &bytes.Buffer{}

	var writeObject func(Vdf, int) error
	writeObject = func(object Vdf, depth int) error {

		if depth > 128 {
			return fmt.Errorf("VDF nesting is too deep")
		}

		// Sort map keys so repeated writes produce the same output.
		keys := make([]string, 0, len(object))
		for key := range object {
			keys = append(keys, key)
		}
		slices.Sort(keys)

		indent := strings.Repeat("\t", depth)
		for _, key := range keys {

			// Quote keys and values using the escaping accepted by Parse.
			switch value := object[key].(type) {
			case string:
				buffer.WriteString(indent + strconv.Quote(key) + "\t\t" + strconv.Quote(value) + "\n")
			case Vdf:
				buffer.WriteString(indent + strconv.Quote(key) + "\n" + indent)
				buffer.WriteRune(textObject)
				buffer.WriteByte('\n')
				err := writeObject(value, depth+1)
				if err != nil {
					return fmt.Errorf("writing object %q: %w", key, err)
				}

				buffer.WriteString(indent)
				buffer.WriteRune(textEnd)
				buffer.WriteByte('\n')
			default:
				return fmt.Errorf("unsupported value type for %q: %T", key, object[key])
			}
		}

		return nil
	}

	err := writeObject(data, 0)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
