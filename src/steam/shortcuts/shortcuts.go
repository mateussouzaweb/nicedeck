package shortcuts

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"

	"github.com/mateussouzaweb/nicedeck/src/steam/vdf"
)

// Load shortcuts from file
func LoadFromFile(shortcutsFile string) ([]*Shortcut, error) {

	var shortcuts []*Shortcut

	// Check if file exist
	info, err := os.Stat(shortcutsFile)
	if os.IsNotExist(err) || info.IsDir() {
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
		if _, ok := item["Exe"]; !ok {
			item["Exe"] = ""
		}
		if _, ok := item["StartDir"]; !ok {
			item["StartDir"] = ""
		}
		if _, ok := item["icon"]; !ok {
			item["icon"] = ""
		}
		if _, ok := item["IconURL"]; !ok {
			item["IconURL"] = ""
		}
		if _, ok := item["Logo"]; !ok {
			item["Logo"] = ""
		}
		if _, ok := item["LogoURL"]; !ok {
			item["LogoURL"] = ""
		}
		if _, ok := item["Cover"]; !ok {
			item["Cover"] = ""
		}
		if _, ok := item["CoverURL"]; !ok {
			item["CoverURL"] = ""
		}
		if _, ok := item["Banner"]; !ok {
			item["Banner"] = ""
		}
		if _, ok := item["BannerURL"]; !ok {
			item["BannerURL"] = ""
		}
		if _, ok := item["Hero"]; !ok {
			item["Hero"] = ""
		}
		if _, ok := item["HeroURL"]; !ok {
			item["HeroURL"] = ""
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

		// Create uppercase variation for strange keys
		item["Icon"] = item["icon"]
		item["AppID"] = item["appid"]
		item["Tags"] = item["tags"]

		// Create tag list
		var tags []string
		for _, tag := range item["Tags"].(vdf.Vdf) {
			tags = append(tags, tag.(string))
		}

		// Convert to manageable shortcut
		shortcut := Shortcut{
			AppID:               item["AppID"].(uint),
			AppName:             item["AppName"].(string),
			Exe:                 item["Exe"].(string),
			StartDir:            item["StartDir"].(string),
			Icon:                item["Icon"].(string),
			IconURL:             item["IconURL"].(string),
			Logo:                item["Logo"].(string),
			LogoURL:             item["LogoURL"].(string),
			Cover:               item["Cover"].(string),
			CoverURL:            item["CoverURL"].(string),
			Banner:              item["Banner"].(string),
			BannerURL:           item["BannerURL"].(string),
			Hero:                item["Hero"].(string),
			HeroURL:             item["HeroURL"].(string),
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
		shortcut.FlatpakAppID = item.FlatpakAppID
		shortcut.LastPlayTime = item.LastPlayTime

		// Merge tags to not lose current ones
		shortcut.Tags = append(shortcut.Tags, item.Tags...)
		shortcut.Tags = slices.Compact(shortcut.Tags)

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

// Sort shortcuts in alphabetical order
func SortShortcuts(shortcuts []*Shortcut) ([]*Shortcut, error) {

	sort.Slice(shortcuts, func(i int, j int) bool {
		return shortcuts[i].AppName < shortcuts[j].AppName
	})

	return shortcuts, nil
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
		item["AppName"] = shortcut.AppName
		item["Exe"] = shortcut.Exe
		item["StartDir"] = shortcut.StartDir
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

		// item["AppID"] = shortcut.AppID
		// item["Icon"] = shortcut.Icon
		// item["IconURL"] = shortcut.IconURL
		// item["Logo"] = shortcut.Logo
		// item["LogoURL"] = shortcut.LogoURL
		// item["Cover"] = shortcut.Cover
		// item["CoverURL"] = shortcut.CoverURL
		// item["Banner"] = shortcut.Banner
		// item["BannerURL"] = shortcut.BannerURL
		// item["Hero"] = shortcut.Hero
		// item["HeroURL"] = shortcut.HeroURL
		// item["Tags"] = tags

		// Keys required to be lowercase
		item["appid"] = shortcut.AppID
		item["icon"] = shortcut.Icon
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
