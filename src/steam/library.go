package steam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
	"github.com/mateussouzaweb/nicedeck/src/steam/vdf"
)

// Library struct
type Library struct {
	DatabasePath  string      `json:"databasePath"`
	BasePath      string      `json:"basePath"`
	Runtime       string      `json:"runtime"`
	AccountId     string      `json:"accountId"`
	AccountName   string      `json:"accountName"`
	ConfigPath    string      `json:"configPath"`
	ImagesPath    string      `json:"imagesPath"`
	ShortcutsPath string      `json:"shortcutsPath"`
	Shortcuts     []*Shortcut `json:"-"`
}

// Load library from database file
func (l *Library) Load(databasePath string) error {

	// Fill basic information
	l.DatabasePath = databasePath
	l.Shortcuts = make([]*Shortcut, 0)

	// Check if database file exist
	exist, err := fs.FileExist(databasePath)
	if err != nil {
		return err
	} else if exist {

		// Read database file content
		content, err := os.ReadFile(databasePath)
		if err != nil {
			return err
		}

		// Retrieve information from database file content when available
		// Missing information will be filled below
		err = json.Unmarshal(content, &l)
		if err != nil {
			return err
		}

	}

	// Retrieve base path
	if l.BasePath == "" {
		l.BasePath, err = GetPath()
		if err != nil {
			return fmt.Errorf("could not detect Steam installation: %s", err)
		}
	}

	// Show message based on Steam detection
	if l.BasePath == "" {
		cli.Printf(cli.ColorWarn, "Steam installation was not detected.\n")
		cli.Printf(cli.ColorWarn, "Please make sure to install and login into Steam first.\n")
	}

	// Check how is Steam running
	if l.Runtime == "" {
		l.Runtime, err = GetRuntime()
		if err != nil {
			return fmt.Errorf("could not determine Steam runtime: %s", err)
		}
	}

	// Retrieve user config path
	if l.ConfigPath == "" {
		l.ConfigPath, err = GetConfigPath()
		if err != nil {
			return fmt.Errorf("could not detect Steam user config path: %s", err)
		}
	}

	// Shortcuts images folder path
	if l.ImagesPath == "" {
		l.ImagesPath = filepath.Join(l.ConfigPath, "grid")
	}

	// Shortcuts file path
	if l.ShortcutsPath == "" {
		l.ShortcutsPath = filepath.Join(l.ConfigPath, "shortcuts.vdf")
	}

	// Determine user account id
	if l.AccountId == "" {
		accountPath := filepath.Dir(l.ConfigPath)
		accountId := filepath.Base(accountPath)
		l.AccountId = accountId
	}

	// Determine user account name
	if l.AccountName == "" {
		l.AccountName = l.AccountId
	}

	// Load Steam shortcuts from VDF file
	// Shortcuts file can possible be updated by other services or Steam UI
	// Because the library represents the Steam data
	// We always read shortcuts from Steam source data
	err = l.LoadShortcuts()
	if err != nil {
		return err
	}

	return nil
}

// Load shortcuts from VDF file into library
func (l *Library) LoadShortcuts() error {

	// Check if VDF file exist
	exist, err := fs.FileExist(l.ShortcutsPath)
	if err != nil {
		return err
	} else if !exist {
		return nil
	}

	// Read VDF file content
	content, err := os.ReadFile(l.ShortcutsPath)
	if err != nil {
		return err
	}

	// Read VDF to resulting map
	buffer := bytes.NewBuffer(content)
	data, err := vdf.ReadVdf(buffer)
	if err != nil {
		return err
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
			DevKit:              item["Devkit"].(uint),
			DevKitGameID:        item["DevkitGameID"].(string),
			DevKitOverrideAppID: item["DevkitOverrideAppID"].(uint),
			LastPlayTime:        item["LastPlayTime"].(uint),
			Tags:                tags,
		}

		l.Shortcuts = append(l.Shortcuts, &shortcut)
	}

	return nil
}

// Save library
func (l *Library) Save() error {

	// Convert database state to JSON representation
	jsonContent, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}

	// Make sure destination folder path exist
	err = os.MkdirAll(filepath.Dir(l.DatabasePath), 0774)
	if err != nil {
		return err
	}

	// Write JSON content to database file
	err = os.WriteFile(l.DatabasePath, jsonContent, 0666)
	if err != nil {
		return err
	}

	// Save updated Steam shortcuts to VDF file
	err = l.SaveShortcuts()
	if err != nil {
		return err
	}

	return nil
}

// Save library shortcuts on VDF file
func (l *Library) SaveShortcuts() error {

	// Create VDF specs from shortcuts
	items := make(vdf.Vdf)
	for index, shortcut := range l.Shortcuts {

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
		item["Devkit"] = shortcut.DevKit
		item["DevkitGameID"] = shortcut.DevKitGameID
		item["DevkitOverrideAppID"] = shortcut.DevKitOverrideAppID
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

	// Write content to VDF file
	err = os.WriteFile(l.ShortcutsPath, content.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

// Sync change history to the library
func (l *Library) Sync(history *shortcuts.History) error {

	// Remove shortcut to sync
	if history.Action == "removed" {
		shortcut := l.FromInternal(history.Original)

		for index, existing := range l.Shortcuts {
			if existing.AppID != shortcut.AppID {
				continue
			}

			updated := make([]*Shortcut, 0)
			updated = append(updated, l.Shortcuts[:index]...)
			updated = append(updated, l.Shortcuts[index+1:]...)
			l.Shortcuts = updated
			break
		}

		return nil
	}

	// Add or update shortcut to sync
	shortcut := l.FromInternal(history.Result)
	found := false

	// First, check if already exists the same shortcut
	// This will prevent double additions
	for index, existing := range l.Shortcuts {
		if existing.AppID != shortcut.AppID {
			continue
		}

		// Keep current value for some keys
		shortcut.IsHidden = existing.IsHidden
		shortcut.AllowDesktopConfig = existing.AllowDesktopConfig
		shortcut.AllowOverlay = existing.AllowOverlay
		shortcut.OpenVR = existing.OpenVR
		shortcut.DevKit = existing.DevKit
		shortcut.DevKitGameID = existing.DevKitGameID
		shortcut.DevKitOverrideAppID = existing.DevKitOverrideAppID
		shortcut.LastPlayTime = existing.LastPlayTime

		// Replace shortcut at index
		l.Shortcuts[index] = shortcut
		found = true
		break
	}

	// Can be added safely
	if !found {
		l.Shortcuts = append(l.Shortcuts, shortcut)
	}

	return nil
}

// Generate a new internal shortcut from Steam shortcut
// Used to merge Steam shortcuts into main library
func (l *Library) ToInternal(specs *Shortcut) *shortcuts.Shortcut {

	shortcut := &shortcuts.Shortcut{}
	shortcut.ID = shortcuts.FromUint(specs.AppID)
	shortcut.Name = specs.AppName
	shortcut.StartDirectory = specs.StartDir
	shortcut.Executable = specs.Exe
	shortcut.LaunchOptions = specs.LaunchOptions
	shortcut.Tags = specs.Tags

	// Merge tags to not lose current ones
	for _, tag := range specs.Tags {
		if !slices.Contains(shortcut.Tags, tag) {
			shortcut.Tags = append(shortcut.Tags, tag)
		}
	}

	return shortcut
}

// Generate a Steam shortcut from internal shortcut
// Used to sync internal shortcuts into Steam library
func (l *Library) FromInternal(specs *shortcuts.Shortcut) *Shortcut {

	shortcut := &Shortcut{}
	shortcut.AppID = shortcuts.ToUint(specs.ID)
	shortcut.AppName = specs.Name
	shortcut.StartDir = specs.StartDirectory
	shortcut.Exe = specs.Executable
	shortcut.LaunchOptions = specs.LaunchOptions
	shortcut.Tags = specs.Tags

	return shortcut
}
