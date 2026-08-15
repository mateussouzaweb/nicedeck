package settings

import (
	"encoding/xml"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Sample XML for es_settings.xml
// <?xml version="1.0"?>
// <string name="InputControllerType" value="xbox360" />
// <string name="ROMDirectory" value="$ROMS/" />
// <string name="Theme" value="linear-es-de" />
// <string name="ThemeAspectRatio" value="automatic" />
// <string name="ThemeColorScheme" value="dark" />
// <string name="ThemeSet" value="linear-es-de" />
// <string name="ThemeTransitions" value="automatic" />
// <string name="ThemeVariant" value="withoutVideos" />

// Setting represents a single configuration string
type Setting struct {
	XMLName xml.Name `xml:"string"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

// Write settings for ES-DE
func WriteSettings(destinationPath string) error {

	settings := []Setting{
		{
			Name:  "InputControllerType",
			Value: "xbox360",
		},
		{
			Name:  "ROMDirectory",
			Value: fs.ExpandPath("$ROMS/"),
		},
		{
			Name:  "Theme",
			Value: "linear-es-de",
		},
		{
			Name:  "ThemeAspectRatio",
			Value: "automatic",
		},
		{
			Name:  "ThemeColorScheme",
			Value: "dark",
		},
		{
			Name:  "ThemeSet",
			Value: "linear-es-de",
		},
		{
			Name:  "ThemeTransitions",
			Value: "automatic",
		},
		{
			Name:  "ThemeVariant",
			Value: "withoutVideos",
		},
	}

	settingsPath := filepath.Join(destinationPath, "settings", "es_settings.xml")
	err := fs.WriteXML(settingsPath, settings)
	if err != nil {
		return err
	}

	return nil
}
