package desktop

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// DesktopEntry represents the standard keys in a .desktop file
type DesktopEntry struct {
	Type            string   `json:"type"`
	Name            string   `json:"name"`
	GenericName     string   `json:"genericName"`
	Comment         string   `json:"comment"`
	Exec            string   `json:"exec"`
	Icon            string   `json:"icon"`
	Terminal        bool     `json:"terminal"`
	MimeType        string   `json:"mimeType"`
	Categories      []string `json:"categories"`
	StartupNotify   bool     `json:"startupNotify"`
	StartupWMClass  string   `json:"startupWMClass"`
	OnlyShowIn      []string `json:"onlyShowIn"`
	NotShowIn       []string `json:"notShowIn"`
	TryExec         string   `json:"tryExec"`
	Hidden          bool     `json:"hidden"`
	NoDisplay       bool     `json:"noDisplay"`
	Keywords        []string `json:"keywords"`
	Actions         []string `json:"actions"`
	DBusActivatable bool     `json:"dBusActivatable"`
	Path            string   `json:"path"`
}

// Read .desktop file and return it structure
func ReadDesktopFile(path string) (*DesktopEntry, error) {

	entry := &DesktopEntry{}

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return entry, err
	}

	defer func() {
		errors.Join(err, file.Close())
	}()

	// Scan file line per line
	scanner := bufio.NewScanner(file)
	inDesktopEntry := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Determine if the content if from the desktop entry section
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			section := line[1 : len(line)-1]
			inDesktopEntry = section == "Desktop Entry"
			continue
		}

		// Ignore content from other sections
		if !inDesktopEntry {
			continue
		}

		if parts := strings.SplitN(line, "=", 2); len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			stringValue := strings.TrimSpace(parts[1])
			boolValue := strings.ToLower(stringValue) == "true"
			sliceValue := strings.Split(strings.Trim(stringValue, ";"), ";")

			switch key {
			case "Type":
				entry.Type = stringValue
			case "Name":
				entry.Name = stringValue
			case "GenericName":
				entry.GenericName = stringValue
			case "Comment":
				entry.Comment = stringValue
			case "Exec":
				entry.Exec = stringValue
			case "Icon":
				entry.Icon = stringValue
			case "Terminal":
				entry.Terminal = boolValue
			case "MimeType":
				entry.MimeType = stringValue
			case "Categories":
				entry.Categories = sliceValue
			case "StartupNotify":
				entry.StartupNotify = boolValue
			case "StartupWMClass":
				entry.StartupWMClass = stringValue
			case "OnlyShowIn":
				entry.OnlyShowIn = sliceValue
			case "NotShowIn":
				entry.NotShowIn = sliceValue
			case "TryExec":
				entry.TryExec = stringValue
			case "Hidden":
				entry.Hidden = boolValue
			case "NoDisplay":
				entry.NoDisplay = boolValue
			case "Keywords":
				entry.Keywords = sliceValue
			case "Actions":
				entry.Actions = sliceValue
			case "DBusActivatable":
				entry.DBusActivatable = boolValue
			case "Path":
				entry.Path = stringValue
			default:
				// Unknown key — skip
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return entry, err
	}

	return entry, nil
}

// Write content to a .desktop file
func WriteDesktopFile(destination string, entry *DesktopEntry) error {

	// Retrieve .desktop string line
	getString := func(key, value string) string {
		if value != "" {
			return fmt.Sprintf("%s=%s", key, value)
		}
		return ""
	}

	// Retrieve .desktop boolean line
	getBoolean := func(key string, value bool) string {
		if value {
			return fmt.Sprintf("%s=true", key)
		} else {
			return fmt.Sprintf("%s=false", key)
		}
	}

	// Retrieve .desktop slice line
	getSlice := func(key string, values []string) string {
		if len(values) > 0 {
			return fmt.Sprintf("%s=%s;", key, strings.Join(values, ";"))
		}
		return ""
	}

	// Create a result from struct parsing
	result := []string{
		"[Desktop Entry]",
		getString("Type", entry.Type),
		getString("Name", entry.Name),
		getString("GenericName", entry.GenericName),
		getString("Comment", entry.Comment),
		getString("Exec", entry.Exec),
		getString("Icon", entry.Icon),
		getBoolean("Terminal", entry.Terminal),
		getString("MimeType", entry.MimeType),
		getSlice("Categories", entry.Categories),
		getBoolean("StartupNotify", entry.StartupNotify),
		getString("StartupWMClass", entry.StartupWMClass),
		getSlice("OnlyShowIn", entry.OnlyShowIn),
		getSlice("NotShowIn", entry.NotShowIn),
		getString("TryExec", entry.TryExec),
		getBoolean("Hidden", entry.Hidden),
		getBoolean("NoDisplay", entry.NoDisplay),
		getSlice("Keywords", entry.Keywords),
		getSlice("Actions", entry.Actions),
		getBoolean("DBusActivatable", entry.DBusActivatable),
		getString("Path", entry.Path),
	}

	// Remove empty lines
	content := []string{}
	for _, line := range result {
		if line != "" {
			content = append(content, line)
		}
	}

	// Write final content to file
	err := fs.WriteFile(destination, strings.Join(content, "\n"))
	if err != nil {
		return err
	}

	return nil
}
