package settings

import (
	"encoding/xml"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/platforms/console"
)

// Sample XML for es_find_rules.xml
// <?xml version="1.0"?>
// <ruleList>
//     <emulator name="OS-SHELL">
//         <rule type="systempath">
//             <entry>zsh</entry>
//             <entry>bash</entry>
//             <entry>sh</entry>
//             <entry>cmd.exe</entry>
//         </rule>
//     </emulator>
//     <emulator name="AZAHAR">
//         <rule type="staticpath">
//             <entry>$EMULATORS/Azahar/azahar.AppImage</entry>
//         </rule>
//     </emulator>
// </ruleList>

// Represents the root element of es_find_rules.xml
type RuleList struct {
	XMLName   xml.Name   `xml:"ruleList"`
	Emulators []Emulator `xml:"emulator"`
}

// Emulator represents a single emulator configuration
type Emulator struct {
	XMLName xml.Name `xml:"emulator"`
	Name    string   `xml:"name,attr"`
	Rules   []Rule   `xml:"rule"`
}

// Rule represents a single find rule configuration
type Rule struct {
	XMLName xml.Name `xml:"rule"`
	Type    string   `xml:"type,attr"`
	Entries []string `xml:"entry"`
}

// Write find rules for ES-DE
func WriteFindRules(destinationPath string) error {

	// Get platforms from console package
	options := &console.Options{}
	platforms, err := console.GetPlatforms(options)
	if err != nil {
		return err
	}

	// ES Find Rules
	ruleList := RuleList{
		XMLName: xml.Name{Local: "ruleList"},
		Emulators: []Emulator{{
			XMLName: xml.Name{Local: "emulator"},
			Name:    "OS-SHELL",
			Rules: []Rule{{
				XMLName: xml.Name{Local: "rule"},
				Type:    "systempath",
				Entries: []string{
					"zsh",
					"bash",
					"sh",
					"cmd.exe",
				},
			}},
		}},
	}

	// Add emulators to RuleList
	for _, platform := range platforms {
		for _, emulator := range platform.Emulators {
			ruleList.Emulators = append(ruleList.Emulators, Emulator{
				XMLName: xml.Name{Local: "emulator"},
				Name:    strings.ToUpper(emulator.Name),
				Rules: []Rule{
					{
						XMLName: xml.Name{Local: "rule"},
						Type:    "staticpath",
						Entries: []string{fs.ExpandPath(emulator.Executable)},
					},
				},
			})
		}
	}

	// Write to XML file
	esFindRulesPath := filepath.Join(destinationPath, "custom_systems", "es_find_rules.xml")
	err = fs.WriteXML(esFindRulesPath, ruleList)
	if err != nil {
		return err
	}

	return nil
}
