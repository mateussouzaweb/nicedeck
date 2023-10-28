package shortcuts

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"

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

// Add shortcut to the list
func AddShortcut(shortcuts []*Shortcut, shortcut *Shortcut) ([]*Shortcut, error) {

	// Check if already exist an app with the same reference
	found := false
	for index, item := range shortcuts {
		if item.AppID != shortcut.AppID {
			continue
		}

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

		// Merge tags to not lose current ones
		for _, tag := range item.Tags {
			if !slices.Contains(shortcut.Tags, tag) {
				shortcut.Tags = append(shortcut.Tags, tag)
			}
		}

		// Replace with new object data
		shortcuts[index] = shortcut

		found = true
		break
	}

	// Append to the list if not exist
	if !found {
		shortcuts = append(shortcuts, shortcut)
	}

	return shortcuts, nil
}

// Remove shortcut from the list
func RemoveShortcut(shortcuts []*Shortcut, shortcut *Shortcut) ([]*Shortcut, error) {

	updated := make([]*Shortcut, 0)
	found := false

	// Instead of appending one by one
	// We detect the one to remove and add others in batch
	for index, item := range shortcuts {
		if item.AppID == shortcut.AppID {
			updated = append(updated, shortcuts[:index]...)
			updated = append(updated, shortcuts[index+1:]...)
			found = true
			break
		}
	}
	if found {
		return updated, nil
	}

	// If not found, then return the same list of shortcuts
	return shortcuts, nil
}

// Sort shortcuts in alphabetical order
func SortShortcuts(shortcuts []*Shortcut) ([]*Shortcut, error) {

	sort.Slice(shortcuts, func(i int, j int) bool {
		return shortcuts[i].AppName < shortcuts[j].AppName
	})

	return shortcuts, nil
}

// Merge shortcuts lists into one
func MergeShortcuts(main []*Shortcut, extra []*Shortcut) []*Shortcut {

	// When match is detected, we always will prefer the extra content
	// When has not match, append item to the list
	for _, item := range extra {
		found := false
		for _, existing := range main {
			if existing.AppID != item.AppID {
				continue
			}

			// Get available data
			existing.AppName = item.AppName
			existing.StartDir = item.StartDir
			existing.Exe = item.Exe
			existing.LaunchOptions = item.LaunchOptions
			existing.ShortcutPath = item.ShortcutPath
			existing.Icon = item.Icon
			existing.IsHidden = item.IsHidden
			existing.AllowDesktopConfig = item.AllowDesktopConfig
			existing.AllowOverlay = item.AllowOverlay
			existing.OpenVR = item.OpenVR
			existing.Devkit = item.Devkit
			existing.DevkitGameID = item.DevkitGameID
			existing.DevkitOverrideAppID = item.DevkitOverrideAppID
			existing.FlatpakAppID = item.FlatpakAppID
			existing.LastPlayTime = item.LastPlayTime

			// Merge tags to not lose current ones
			for _, tag := range item.Tags {
				if !slices.Contains(existing.Tags, tag) {
					existing.Tags = append(existing.Tags, tag)
				}
			}

			found = true
			break
		}

		if !found {
			main = append(main, item)
		}
	}

	return main
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
