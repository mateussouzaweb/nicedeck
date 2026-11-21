package vdf

import (
	"strconv"
)

// Vdf map shared by text and binary VDF/ACF files
type Vdf map[string]any

// Return key value as string
func (vdf Vdf) String(key string) string {
	value, _ := vdf[key].(string)
	return value
}

// Return a binary uint32 or parse a text number, or zero if missing, invalid, or out of range
func (vdf Vdf) Uint32(key string) uint32 {
	if value, ok := vdf[key].(uint32); ok {
		return value
	}

	value, err := strconv.ParseUint(vdf.String(key), 10, 32)
	if err != nil {
		return 0
	}
	return uint32(value)
}

// Return a non-nil nested object, creating and storing it when needed.
// A nil receiver returns an empty object without storing it.
func (vdf Vdf) Object(key string) Vdf {
	value, _ := vdf[key].(Vdf)
	if value == nil {
		value = Vdf{}
	}
	return value
}
