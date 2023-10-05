package steam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mateussouzaweb/nicedeck/src/vdf"
)

type Config struct {
	ArtworksPath  string      `json:"artworksPath"`
	DebugFile     string      `json:"debugFile"`
	ShortcutsFile string      `json:"shortcutsFile"`
	Shortcuts     []*Shortcut `json:"shortcuts"`
}

func (c *Config) LoadShortcuts() error {

	// Check if file exist
	info, err := os.Stat(c.ShortcutsFile)
	if os.IsNotExist(err) || info.IsDir() {
		return nil
	}

	// Read file content
	content, err := os.ReadFile(c.ShortcutsFile)
	if err != nil {
		return err
	}

	// Read to resulting map
	buffer := bytes.NewBuffer(content)
	data, err := vdf.ReadVdf(buffer)
	if err != nil {
		return err
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

		c.Shortcuts = append(c.Shortcuts, &shortcut)

	}

	return nil
}

// Add shortcut to config
func (c *Config) AddShortcut(shortcut *Shortcut) error {

	// Check if already exist an app with the same reference
	found := false
	for index, item := range c.Shortcuts {
		if item.AppID == shortcut.AppID {

			// Keep current value for some keys
			shortcut.IsHidden = item.IsHidden
			shortcut.AllowDesktopConfig = item.AllowDesktopConfig
			shortcut.AllowOverlay = item.AllowOverlay
			shortcut.OpenVR = item.OpenVR
			shortcut.Devkit = item.Devkit
			shortcut.DevkitGameID = item.DevkitGameID
			shortcut.DevkitOverrideAppID = item.DevkitOverrideAppID
			shortcut.FlatpakAppID = item.FlatpakAppID
			shortcut.LastPlayTime = item.LastPlayTime
			shortcut.Tags = item.Tags

			// Replace with new object data
			c.Shortcuts[index] = shortcut

			found = true
			break
		}
	}

	// Append to the list if not exist
	if !found {
		c.Shortcuts = append(c.Shortcuts, shortcut)
	}

	return nil
}

// Save updated content on the shortcuts file
func (c *Config) SaveShotcuts() error {

	// Create vdf from shortcuts
	shortcuts := make(vdf.Vdf)
	for index, shortcut := range c.Shortcuts {

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
	err = os.WriteFile(c.ShortcutsFile, content.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

// Save updated content on the debug file
func (c *Config) SaveDebug() error {

	// Save JSON copy for debugging
	jsonContent, err := json.Marshal(c)
	if err != nil {
		return err
	}

	// Write JSON content to file
	err = os.WriteFile(c.DebugFile, jsonContent, 0666)
	if err != nil {
		return err
	}

	return nil
}

var _config *Config

// Use given runtime config
func Use(config *Config) (func() error, error) {
	_config = config

	err := _config.LoadShortcuts()
	save := func() error {
		err := _config.SaveDebug()
		if err != nil {
			return err
		}
		err = _config.SaveShotcuts()
		if err != nil {
			return err
		}
		return nil
	}

	return save, err
}
