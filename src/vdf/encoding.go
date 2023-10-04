package vdf

// Marker of VDF binaries
const (
	markerMap         byte = 0x00
	markerString      byte = 0x01
	markerNumber      byte = 0x02
	markerEndOfMap    byte = 0x08
	markerEndOfString byte = 0x00
)

type Vdf map[string]interface{}
