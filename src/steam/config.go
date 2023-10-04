package steam

import (
	"bytes"
	"fmt"
	"os"

	"github.com/mateussouzaweb/nicedeck/src/vdf"
)

type Config struct {
	ShortcutsFile string
	ArtworksPath  string
	Shortcuts     []*Shortcut
}

var _config *Config

// Use given runtime config
func Use(config *Config) (func() error, error) {

	_config = config

	// Save updated content on the shortcuts file
	saveShortcuts := func() error {

		// Create vdf from shortcuts
		shortcuts := make(vdf.Vdf)
		for index, shortcut := range _config.Shortcuts {

			tags := make(vdf.Vdf)
			for tindex, tag := range shortcut.Tags {
				position := fmt.Sprintf("%v", tindex)
				tags[position] = tag
			}

			item := vdf.Vdf{}
			item["appid"] = shortcut.AppID
			item["AppName"] = shortcut.AppName
			item["Exe"] = shortcut.Exe
			item["StartDir"] = shortcut.StartDir
			item["icon"] = shortcut.Icon
			item["ShortcutPath"] = shortcut.ShortcutPath
			item["LaunchOptions"] = shortcut.LaunchOptions
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
			shortcuts[position] = item

		}

		data := vdf.Vdf{}
		data["shortcuts"] = shortcuts

		// Transform VDF into bytes
		content, err := vdf.WriteVdf(data)
		if err != nil {
			return err
		}

		// Write content to file
		err = os.WriteFile(_config.ShortcutsFile, content.Bytes(), 0666)
		if err != nil {
			return err
		}

		return nil
	}

	// Check if file exist
	info, err := os.Stat(_config.ShortcutsFile)
	if os.IsNotExist(err) || info.IsDir() {
		return saveShortcuts, nil
	}

	// Read file content
	content, err := os.ReadFile(_config.ShortcutsFile)
	if err != nil {
		return saveShortcuts, err
	}

	// Read to resulting map
	buffer := bytes.NewBuffer(content)
	data, err := vdf.ReadVdf(buffer)
	if err != nil {
		return saveShortcuts, err
	}

	// Map to struct of shortcuts
	// We don't care about positioning
	shortcuts := data["shortcuts"].(vdf.Vdf)
	for _, item := range shortcuts {

		item := item.(vdf.Vdf)

		// Make sure has data
		if _, ok := item["appid"]; !ok {
			item["appid"] = uint(0)
		}
		if _, ok := item["AppName"]; !ok {
			item["AppName"] = ""
		}
		if _, ok := item["Exe"]; !ok {
			item["Exe"] = ""
		}
		if _, ok := item["StartDir"]; !ok {
			item["StartDir"] = ""
		}
		if _, ok := item["icon"]; !ok {
			item["icon"] = ""
		}
		if _, ok := item["ShortcutPath"]; !ok {
			item["ShortcutPath"] = ""
		}
		if _, ok := item["LaunchOptions"]; !ok {
			item["LaunchOptions"] = ""
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

		// Convert to shortcut
		var tags []string
		for _, tag := range item["tags"].(vdf.Vdf) {
			tags = append(tags, tag.(string))
		}

		shortcut := Shortcut{
			AppID:               item["appid"].(uint),
			AppName:             item["AppName"].(string),
			Exe:                 item["Exe"].(string),
			StartDir:            item["StartDir"].(string),
			Icon:                item["icon"].(string),
			ShortcutPath:        item["ShortcutPath"].(string),
			LaunchOptions:       item["LaunchOptions"].(string),
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

		_config.Shortcuts = append(_config.Shortcuts, &shortcut)

	}

	return saveShortcuts, nil
}
