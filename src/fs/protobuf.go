package fs

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Protobuf wire types
const (
	wireVarint  = 0 // int32, int64, uint32, uint64, sint32, sint64, bool, enum
	wireFixed64 = 1 // fixed64, sfixed64, double
	wireLength  = 2 // string, bytes, embedded messages, packed repeated fields
	wireFixed32 = 5 // fixed32, sfixed32, float
)

// ReadProtoBuf and decodes content as a protobuf-wire-format message
// DecodeProtoBuf bytes content as a protobuf-wire-format message
// Keys are dotted field-number paths reflecting nesting, e.g. "3.1", "4.4.6.1"
// Repeated fields get an index suffix: "6", "6[1]", "6[2]", ...
// All values (varint, fixed32/64, strings, bytes) are returned as plain text.
func ReadProtoBuf(path string) (map[string]string, error) {

	// Check if file exist
	result := make(map[string]string)
	exist, err := FileExist(path)
	if err != nil {
		return result, err
	} else if exist {
		cli.Debug("Reading ProtonBuf %s\n", path)
	}

	// Read file bytes
	data, err := os.ReadFile(path)
	if err != nil {
		return result, err
	}

	// readVarint reads a base-128 varint from the start of b.
	// Returns the decoded value and the number of bytes consumed,
	// or (0, -1) if b doesn't contain a complete, valid varint.
	readVarint := func(b []byte) (value uint64, bytesRead int) {
		var shift uint
		for i := 0; i < len(b) && i < 10; i++ {
			value |= uint64(b[i]&0x7f) << shift
			if b[i]&0x80 == 0 {
				return value, i + 1 // high bit clear: this was the last byte
			}
			shift += 7
		}
		return 0, -1 // ran out of bytes without finding a terminating byte
	}

	// leUint interprets b as a little-endian unsigned integer.
	// Used to turn raw fixed32/fixed64 bytes into a printable number.
	leUint := func(b []byte) uint64 {
		var value uint64
		for i := len(b) - 1; i >= 0; i-- {
			value = value<<8 | uint64(b[i])
		}
		return value
	}

	// isPrintableUTF8 reports whether b looks like human-readable text
	isPrintableUTF8 := func(b []byte) bool {
		if len(b) == 0 {
			return false
		}

		printable, total := 0, 0
		for _, r := range string(b) {
			total++
			if r == unicode.ReplacementChar {
				return false // invalid UTF-8 byte sequence
			}
			if unicode.IsPrint(r) || unicode.IsSpace(r) {
				printable++
			}
		}

		const printableThreshold = 0.95
		return float64(printable)/float64(total) > printableThreshold
	}

	// hexString renders b as space-separated hex bytes, e.g. "de ad be ef".
	// Used as the last-resort representation when a value is neither a
	// valid nested message nor printable text.
	hexString := func(b []byte) string {
		value := ""
		for i, c := range b {
			if i > 0 {
				value += " "
			}
			value += strconv.FormatInt(int64(c), 16)
		}
		return value
	}

	// looksLikeMessage does a cheap dry-run parse of b: valid tags, valid
	// wire types, and lengths that stay in bounds. It exists so that a
	// length-delimited field is only treated as a nested message when the
	// bytes are actually shaped like one, rather than misreading an
	// arbitrary binary blob or string as nested protobuf.
	looksLikeMessage := func(b []byte) bool {
		position, fields := 0, 0
		for position < len(b) {
			tag, n := readVarint(b[position:])
			if n <= 0 {
				return false
			}

			fieldNum, wt := tag>>3, tag&0x7
			if fieldNum == 0 || wt > 5 || wt == 3 || wt == 4 {
				return false // field 0 is invalid; wire types 3/4 are legacy groups, unsupported
			}

			position += n
			switch wt {
			case wireVarint:
				_, n := readVarint(b[position:])
				if n <= 0 {
					return false
				}
				position += n
			case wireFixed64:
				position += 8
			case wireFixed32:
				position += 4
			case wireLength:
				length, n := readVarint(b[position:])
				if n <= 0 {
					return false
				}
				position += n + int(length)
			}

			if position > len(b) {
				return false
			}

			fields++
		}

		return fields > 0
	}

	// seen tracks how many times each dotted key has already been written,
	// so repeated fields become "key", "key[1]", "key[2]", ... instead of
	// overwriting one another in result.
	seen := make(map[string]int)
	nextKey := func(key string) string {
		count, exists := seen[key]
		if !exists {
			seen[key] = 0
			return key
		}

		seen[key] = count + 1
		return fmt.Sprintf("%s[%d]", key, count+1)
	}

	// decode walks data as a protobuf message, writing flattened key/value
	// pairs into result under prefix. Returns false if data isn't valid wire format.
	var decode func(data []byte, prefix string) bool
	decode = func(data []byte, prefix string) bool {
		position, fieldCount := 0, 0

		for position < len(data) {
			// Every field starts with a varint tag encoding
			// (field number << 3) | wire type.
			tag, n := readVarint(data[position:])
			if n <= 0 {
				// No more valid tag bytes: that's fine only if we're
				// exactly at the end and already decoded something.
				return fieldCount > 0 && position == len(data)
			}
			position += n

			fieldNum, wt := tag>>3, tag&0x7
			if fieldNum == 0 {
				return false
			}

			key := strconv.FormatUint(fieldNum, 10)
			if prefix != "" {
				key = prefix + "." + key
			}

			switch wt {
			case wireVarint:
				val, n := readVarint(data[position:])
				if n <= 0 {
					return false
				}
				position += n
				result[nextKey(key)] = strconv.FormatUint(val, 10)

			case wireFixed64:
				if position+8 > len(data) {
					return false
				}
				result[nextKey(key)] = strconv.FormatUint(leUint(data[position:position+8]), 10)
				position += 8

			case wireFixed32:
				if position+4 > len(data) {
					return false
				}
				result[nextKey(key)] = strconv.FormatUint(leUint(data[position:position+4]), 10)
				position += 4

			case wireLength:
				// This is either a nested message, a UTF-8 string, or raw
				// bytes. Try each interpretation in that order and keep
				// whichever one fits.
				length, n := readVarint(data[position:])
				if n <= 0 || position+n+int(length) > len(data) {
					return false
				}
				position += n
				chunk := data[position : position+int(length)]
				position += int(length)

				indexedKey := nextKey(key)
				switch {
				case len(chunk) > 0 && looksLikeMessage(chunk) && decode(chunk, indexedKey):
					// Nested message fields were written directly into
					// result by the recursive decode call above.
				case isPrintableUTF8(chunk):
					result[indexedKey] = string(chunk)
				default:
					result[indexedKey] = hexString(chunk)
				}

			default: // groups (3, 4) and any other unknown wire type: unsupported
				return false
			}

			fieldCount++
		}

		return true
	}

	if !decode(data, "") {
		return result, fmt.Errorf("data does not look like a valid protobuf message")
	}

	return result, nil
}
