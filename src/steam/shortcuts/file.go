package shortcuts

import (
	"bytes"
	"fmt"
	"os"

	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/steam/vdf"
)

// Load shortcuts from file
func LoadFromFile(shortcutsFile string) ([]*Shortcut, error) {

	var shortcuts []*Shortcut

	// Check if file exist
	exist, err := fs.FileExist(shortcutsFile)
	if err != nil {
		return shortcuts, err
	} else if !exist {
		return shortcuts, nil
	}

	// Read file content
	content, err := os.ReadFile(shortcutsFile)
	if err != nil {
		return shortcuts, err
	}

	// Read to resulting map
	buffer := bytes.NewBuffer(content)
	data, err := vdf.ReadVdf(buffer)
	if err != nil {
		return shortcuts, err
	}

	// Map to struct of shortcuts
	// We don't care about positioning
	found := data["shortcuts"].(vdf.Vdf)
	for _, item := range found {

		item := item.(vdf.Vdf)

		// Make sure has data
		if _, ok := item["appid"]; !ok {
			item["appid"] = uint(0)
		}
		if _, ok := item["AppName"]; !ok {
			item["AppName"] = ""
		}
		if _, ok := item["StartDir"]; !ok {
			item["StartDir"] = ""
		}
		if _, ok := item["Exe"]; !ok {
			item["Exe"] = ""
		}
		if _, ok := item["LaunchOptions"]; !ok {
			item["LaunchOptions"] = ""
		}
		if _, ok := item["ShortcutPath"]; !ok {
			item["ShortcutPath"] = ""
		}
		if _, ok := item["icon"]; !ok {
			item["icon"] = ""
		}
		if _, ok := item["IsHidden"]; !ok {
			item["IsHidden"] = uint(0)
		}
		if _, ok := item["AllowDesktopConfig"]; !ok {
			item["AllowDesktopConfig"] = uint(0)
		}
		if _, ok := item["AllowOverlay"]; !ok {
			item["AllowOverlay"] = uint(0)
		}
		if _, ok := item["OpenVR"]; !ok {
			item["OpenVR"] = uint(0)
		}
		if _, ok := item["Devkit"]; !ok {
			item["Devkit"] = uint(0)
		}
		if _, ok := item["DevkitGameID"]; !ok {
			item["DevkitGameID"] = ""
		}
		if _, ok := item["DevkitOverrideAppID"]; !ok {
			item["DevkitOverrideAppID"] = uint(0)
		}
		if _, ok := item["FlatpakAppID"]; !ok {
			item["FlatpakAppID"] = ""
		}
		if _, ok := item["LastPlayTime"]; !ok {
			item["LastPlayTime"] = uint(0)
		}
		if _, ok := item["tags"]; !ok {
			item["tags"] = vdf.Vdf{}
		}

		// Create tag list
		var tags []string
		for _, tag := range item["tags"].(vdf.Vdf) {
			tags = append(tags, tag.(string))
		}

		// Convert to manageable shortcut
		shortcut := Shortcut{
			AppID:               item["appid"].(uint),
			AppName:             item["AppName"].(string),
			StartDir:            item["StartDir"].(string),
			Exe:                 item["Exe"].(string),
			LaunchOptions:       item["LaunchOptions"].(string),
			ShortcutPath:        item["ShortcutPath"].(string),
			Icon:                item["icon"].(string),
			IsHidden:            item["IsHidden"].(uint),
			AllowDesktopConfig:  item["AllowDesktopConfig"].(uint),
			AllowOverlay:        item["AllowOverlay"].(uint),
			OpenVR:              item["OpenVR"].(uint),
			Devkit:              item["Devkit"].(uint),
			DevkitGameID:        item["DevkitGameID"].(string),
			DevkitOverrideAppID: item["DevkitOverrideAppID"].(uint),
			FlatpakAppID:        item["FlatpakAppID"].(string),
			LastPlayTime:        item["LastPlayTime"].(uint),
			Tags:                tags,
		}

		shortcuts = append(shortcuts, &shortcut)

	}

	// Sort list of shortcuts
	return SortShortcuts(shortcuts)
}

// Save shortcuts list to shortcuts file
func SaveToFile(shortcuts []*Shortcut, destinationFile string) error {

	// Create vdf from shortcuts
	items := make(vdf.Vdf)
	for index, shortcut := range shortcuts {

		tags := make(vdf.Vdf)
		for tagIndex, tag := range shortcut.Tags {
			position := fmt.Sprintf("%v", tagIndex)
			tags[position] = tag
		}

		item := vdf.Vdf{}
		item["appid"] = shortcut.AppID
		item["AppName"] = shortcut.AppName
		item["StartDir"] = shortcut.StartDir
		item["Exe"] = shortcut.Exe
		item["LaunchOptions"] = shortcut.LaunchOptions
		item["ShortcutPath"] = shortcut.ShortcutPath
		item["icon"] = shortcut.Icon
		item["IsHidden"] = shortcut.IsHidden
		item["AllowDesktopConfig"] = shortcut.AllowDesktopConfig
		item["AllowOverlay"] = shortcut.AllowOverlay
		item["OpenVR"] = shortcut.OpenVR
		item["Devkit"] = shortcut.Devkit
		item["DevkitGameID"] = shortcut.DevkitGameID
		item["DevkitOverrideAppID"] = shortcut.DevkitOverrideAppID
		item["FlatpakAppID"] = shortcut.FlatpakAppID
		item["LastPlayTime"] = shortcut.LastPlayTime
		item["tags"] = tags

		position := fmt.Sprintf("%v", index)
		items[position] = item

	}

	data := vdf.Vdf{}
	data["shortcuts"] = items

	// Transform VDF into bytes
	content, err := vdf.WriteVdf(data)
	if err != nil {
		return err
	}

	// Write content to file
	err = os.WriteFile(destinationFile, content.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}
