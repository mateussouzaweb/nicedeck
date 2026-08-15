package settings

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/platforms"
)

// Sample XML for es_systems.xml
// <?xml version="1.0"?>
// <systemList>
//     <system>
//         <name>n3ds</name>
//         <fullname>Nintendo 3DS</fullname>
//         <platform>n3ds</platform>
//         <theme>n3ds</theme>
//         <path>%ROMPATH%/3DS</path>
//         <extension>.3ds .3DS .3dsx .3DSX .app .APP .axf .AXF .cci .CCI .cxi .CXI .elf .ELF .7z .7Z .zip .ZIP</extension>
//         <command label="AZAHAR">%EMULATOR_AZAHAR% %ROM%</command>
//     </system>
// </systemList>

// SystemList represents the root element of es_systems.xml
type SystemList struct {
	XMLName xml.Name `xml:"systemList"`
	Systems []System `xml:"system"`
}

// System represents a single game system configuration
type System struct {
	XMLName   xml.Name  `xml:"system"`
	Name      string    `xml:"name"`
	FullName  string    `xml:"fullname"`
	Path      string    `xml:"path"`
	Platform  string    `xml:"platform"`
	Theme     string    `xml:"theme"`
	Extension string    `xml:"extension"`
	Commands  []Command `xml:"command"`
}

// Command represents an emulator command configuration
type Command struct {
	XMLName xml.Name `xml:"command"`
	Label   string   `xml:"label,attr"`
	Value   string   `xml:",chardata"`
}

// Platform to Theme mapping for ES-DE
var platformThemeMap = map[string]string{
	"DC":     "dreamcast",
	"GBA":    "gba",
	"GC":     "gc",
	"3DS":    "n3ds",
	"N64":    "n64",
	"NDS":    "nds",
	"PS1":    "psx",
	"PS2":    "ps2",
	"PS3":    "ps3",
	"PS4":    "ps4",
	"PSP":    "psp",
	"PSVITA": "psvita",
	"SWITCH": "switch",
	"WII":    "wii",
	"WIIU":   "wiiu",
	"XBOX":   "xbox",
	"X360":   "xbox360",
}

// Write systems for ES-DE
func WriteSystems(destinationPath string) error {

	// Get platforms from console package
	consolePlatforms, err := platforms.GetConsoles()
	if err != nil {
		return err
	}

	// ES Systems
	systemList := SystemList{
		Systems: []System{},
	}

	for _, platform := range consolePlatforms {

		extensions := []string{}
		commands := []Command{}

		for _, emulator := range platform.Emulators {

			// Include extensions for the emulator on platform
			extensions = append(extensions, strings.ToLower(emulator.Extensions))
			extensions = append(extensions, strings.ToUpper(emulator.Extensions))

			// Add commands for the emulator on platform
			commandValue := fmt.Sprintf("%%EMULATOR_%s%% %s", strings.ToUpper(emulator.Name), emulator.LaunchOptions)
			commandValue = strings.ReplaceAll(commandValue, "${ROM}", "%ROM%") // Ensure %ROM% is preserved
			commands = append(commands, Command{
				Label: strings.ToUpper(emulator.Name),
				Value: commandValue,
			})

		}

		extensionsList := strings.Split(strings.Join(extensions, " "), " ")
		extensionsList = slices.Compact(extensionsList)

		theme, ok := platformThemeMap[platform.Name]
		if !ok {
			theme = strings.ToLower(platform.Name)
			theme = strings.ReplaceAll(theme, " ", "")
			theme = strings.ReplaceAll(theme, "-", "")
		}

		systemList.Systems = append(systemList.Systems, System{
			Name:      strings.ToLower(platform.Name),
			FullName:  strings.ToLower(platform.Console),
			Platform:  theme,
			Theme:     theme,
			Path:      filepath.Join("%ROMPATH%", platform.Folder),
			Extension: strings.Join(extensionsList, " "),
			Commands:  commands,
		})

	}

	// Write to XML file
	esSystemsPath := filepath.Join(destinationPath, "custom_systems", "es_systems.xml")
	err = fs.WriteXML(esSystemsPath, systemList)
	if err != nil {
		return err
	}

	return nil
}
