package steam

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
	"github.com/mateussouzaweb/nicedeck/src/steam/controller"
	"github.com/mateussouzaweb/nicedeck/src/steam/vdf"
)

type Image = shortcuts.Image
type History = shortcuts.History
type Internal = shortcuts.Shortcut

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

	// Check if VDF file exist
	exist, err = fs.FileExist(l.ShortcutsPath)
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

// Perform library and Steam setup
func (l *Library) Setup() error {

	// Write controller templates
	controllerTemplatesPaths := filepath.Join(l.BasePath, "controller_base", "templates")
	err := controller.WriteTemplates(controllerTemplatesPaths)
	if err != nil {
		return fmt.Errorf("could not perform Steam controller setup: %s", err)
	}

	// Make sure Steam on flatpak has the necessary permission
	if _, ok := GetPackage().(*linux.Flatpak); ok {
		err := GetPackage().(*linux.Flatpak).ApplyOverrides()
		if err != nil {
			return fmt.Errorf("could not perform Steam runtime setup: %s", err)
		}
	}

	return nil
}

// Sync change history to the library
func (l *Library) Sync(history *History) error {

	// Remove shortcut to sync
	if history.Action == "removed" {
		reference := history.Original
		shortcut := l.FromInternal(reference)

		for index, existing := range l.Shortcuts {
			if existing.AppID != shortcut.AppID {
				continue
			}

			// Handle shortcut assets
			err := l.Assets(reference, existing, "remove", true)
			if err != nil {
				return err
			}

			// Update library of shortcuts
			updated := make([]*Shortcut, 0)
			updated = append(updated, l.Shortcuts[:index]...)
			updated = append(updated, l.Shortcuts[index+1:]...)
			l.Shortcuts = updated
			break
		}

		return nil
	}

	// Add or update shortcut to sync
	reference := history.Result
	shortcut := l.FromInternal(history.Result)
	found := false

	// First, check if already exists the same shortcut
	// This will prevent double additions
	for index, existing := range l.Shortcuts {
		if existing.AppID != shortcut.AppID {
			continue
		}

		// Keep current value for some keys
		// These are Steam internal data, we don't modify it
		shortcut.IsHidden = existing.IsHidden
		shortcut.AllowDesktopConfig = existing.AllowDesktopConfig
		shortcut.AllowOverlay = existing.AllowOverlay
		shortcut.OpenVR = existing.OpenVR
		shortcut.DevKit = existing.DevKit
		shortcut.DevKitGameID = existing.DevKitGameID
		shortcut.DevKitOverrideAppID = existing.DevKitOverrideAppID
		shortcut.LastPlayTime = existing.LastPlayTime

		// Handle shortcut assets
		err := l.Assets(reference, shortcut, "sync", true)
		if err != nil {
			return err
		}

		// Replace shortcut at index
		l.Shortcuts[index] = shortcut
		found = true
		break
	}

	// Not found in library, can be added safely
	if !found {

		// Handle shortcut assets
		err := l.Assets(reference, shortcut, "sync", true)
		if err != nil {
			return err
		}

		// Add shortcut to the library
		l.Shortcuts = append(l.Shortcuts, shortcut)

	}

	return nil
}

// Process assets for shortcut based on action
func (l *Library) Assets(specs *Internal, shortcut *Shortcut, action string, overwrite bool) error {

	// Steam image formats:
	// - Logo: ${APPID}_logo.png
	// - Icon: ${APPID}_icon.(ico|png)
	// - Cover: ${APPID}p.(jpg|png)
	// - Banner: ${APPID}.(jpg|png)
	// - Hero: ${APPID}_hero.(jpg|png)

	// Handle images
	// Process usually means copy image from path to path
	iconImage := &Image{
		SourcePath:      specs.IconPath,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%v_icon", shortcut.AppID),
		Extensions:      []string{".png", ".ico"},
	}
	logoImage := &Image{
		SourcePath:      specs.LogoPath,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%v_logo", shortcut.AppID),
		Extensions:      []string{".png"},
	}
	coverImage := &Image{
		SourcePath:      specs.CoverPath,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%vp", shortcut.AppID),
		Extensions:      []string{".png", ".jpg"},
	}
	bannerImage := &Image{
		SourcePath:      specs.BannerPath,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%v", shortcut.AppID),
		Extensions:      []string{".png", ".jpg"},
	}
	heroImage := &Image{
		SourcePath:      specs.HeroPath,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%v_hero", shortcut.AppID),
		Extensions:      []string{".png", ".jpg"},
	}

	// Sync all images based on the action
	if action == "sync" || action == "add" {
		err := iconImage.Process(overwrite)
		if err != nil {
			return err
		}
		err = logoImage.Process(overwrite)
		if err != nil {
			return err
		}
		err = coverImage.Process(overwrite)
		if err != nil {
			return err
		}
		err = bannerImage.Process(overwrite)
		if err != nil {
			return err
		}
		err = heroImage.Process(overwrite)
		if err != nil {
			return err
		}

		// Note: Steam only stores the icon path on shortcut
		// shortcut.Logo = logoImage.TargetPath
		// shortcut.Cover = coverImage.TargetPath
		// shortcut.Banner = bannerImage.TargetPath
		// shortcut.Hero = heroImage.TargetPath
		shortcut.Icon = iconImage.TargetPath
	}

	// Remove images if specified
	if action == "remove" {
		err := iconImage.Remove()
		if err != nil {
			return err
		}
		err = logoImage.Remove()
		if err != nil {
			return err
		}
		err = coverImage.Remove()
		if err != nil {
			return err
		}
		err = bannerImage.Remove()
		if err != nil {
			return err
		}
		err = heroImage.Remove()
		if err != nil {
			return err
		}
	}

	return nil
}

// Generate a new internal shortcut from Steam shortcut
// Used to sync Steam shortcuts into main library
func (l *Library) ToInternal(shortcut *Shortcut) *Internal {

	internal := &Internal{}
	internal.ID = shortcuts.FromUint(shortcut.AppID)
	internal.Name = shortcut.AppName
	internal.StartDirectory = shortcut.StartDir
	internal.Executable = CleanExec(shortcut.Exe)
	internal.LaunchOptions = shortcut.LaunchOptions
	internal.Tags = shortcut.Tags

	return internal
}

// Generate a new Steam shortcut from internal shortcut
// Used to sync internal shortcuts into Steam library
func (l *Library) FromInternal(internal *Internal) *Shortcut {

	shortcut := &Shortcut{}
	shortcut.AppID = shortcuts.ToUint(internal.ID)
	shortcut.AppName = internal.Name
	shortcut.StartDir = internal.StartDirectory
	shortcut.Exe = EnsureExec(l.Runtime, internal.Executable)
	shortcut.LaunchOptions = internal.LaunchOptions
	shortcut.Tags = internal.Tags

	return shortcut
}
