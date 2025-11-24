package steam

import (
	"bytes"
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
	Timestamp     int64       `json:"-"`
}

// Load library from database file
func (l *Library) Load(databasePath string) error {

	// Reset and fill basic information
	l.DatabasePath = databasePath
	l.BasePath = ""
	l.Runtime = ""
	l.AccountId = ""
	l.AccountName = ""
	l.ConfigPath = ""
	l.ImagesPath = ""
	l.ShortcutsPath = ""
	l.Shortcuts = make([]*Shortcut, 0)
	l.Timestamp = 0

	// Check if Steam is installed
	steamPackage := GetPackage()
	installed, err := steamPackage.Installed()
	if err != nil {
		return err
	} else if !installed {
		return nil
	}

	// Read database file content
	// Missing information will be filled below
	err = fs.ReadJSON(databasePath, &l)
	if err != nil {
		return err
	}

	// Retrieve base path
	l.BasePath, err = GetBasePath()
	if err != nil {
		return err
	}

	// Check how is Steam running
	l.Runtime = steamPackage.Runtime()
	if l.Runtime == "none" {
		return fmt.Errorf("could not determine Steam runtime")
	}

	// Retrieve user config path
	if l.ConfigPath == "" {
		l.ConfigPath, err = GetConfigPath()
		if err != nil {
			return err
		}
	}

	// When no user config path is detect
	// Gracefully ignore processing for this library
	if l.ConfigPath == "" {
		cli.Printf(cli.ColorWarn, "Note: Steam account was not detected.\n")
		cli.Printf(cli.ColorWarn, "Please make sure to login into Steam first to sync library.\n")
		return nil
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

	// Load shortcuts from old format for migrated reasons
	// When using this model, we avoid loading the shortcuts from VDF file
	// @deprecated and will be removed in future versions
	var deprecatedFile = filepath.Join(l.ConfigPath, "niceconfig.json")
	var deprecatedConfig struct {
		Shortcuts []*Shortcut `json:"shortcuts"`
	}

	err = fs.ReadJSON(deprecatedFile, &deprecatedConfig)
	if err != nil {
		return err
	} else if len(deprecatedConfig.Shortcuts) > 0 {
		l.Shortcuts = deprecatedConfig.Shortcuts
		return nil
	}

	// Load Steam shortcuts from VDF file
	// Shortcuts file can possible be updated by other services or Steam UI
	// Because the library represents the Steam data
	// We always read shortcuts from Steam source data

	// Check if VDF file exist
	exist, err := fs.FileExist(l.ShortcutsPath)
	if err != nil {
		return err
	} else if !exist {
		return nil
	}

	cli.Debug("Reading VDF at %s\n", l.ShortcutsPath)

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
		tags := []string{}
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

	// Read VDF modified time and use as timestamp reference
	l.Timestamp, err = fs.ModificationTime(l.ShortcutsPath)
	if err != nil {
		return err
	}

	return nil
}

// Save library
func (l *Library) Save() error {

	// Check if can save library
	if l.BasePath == "" || l.AccountId == "" {
		return nil
	}

	// Save database state to file
	err := fs.WriteJSON(l.DatabasePath, l)
	if err != nil {
		return err
	}

	// Make sure Steam on flatpak has the necessary permission
	steamPackage := GetPackage()
	if _, ok := steamPackage.(*linux.Flatpak); ok {
		err := steamPackage.(*linux.Flatpak).ApplyOverrides()
		if err != nil {
			return fmt.Errorf("could not perform Steam runtime setup: %s", err)
		}
	}

	// Write controller template
	controllerTemplatesPaths := filepath.Join(
		l.BasePath, "controller_base", "templates",
	)

	err = controller.WriteTemplates(controllerTemplatesPaths)
	if err != nil {
		return fmt.Errorf("could not perform Steam controller setup: %s", err)
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

	cli.Debug("Writing VDF at %s\n", l.ShortcutsPath)

	// Transform VDF into bytes
	content, err := vdf.WriteVdf(data)
	if err != nil {
		return err
	}

	// Write content to VDF file
	err = fs.WriteFile(l.ShortcutsPath, content.String())
	if err != nil {
		return err
	}

	// Move deprecated file to avoid loading it again
	// @deprecated and will be removed in future versions
	deprecatedFile := filepath.Join(l.ConfigPath, "niceconfig.json")
	deprecatedDestination := fmt.Sprintf("%s.deprecated", deprecatedFile)
	err = fs.MoveFile(deprecatedFile, deprecatedDestination)
	if err != nil {
		return err
	}

	cli.Printf(cli.ColorNotice, "Note: Steam library has been updated.\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	return nil
}

// Sync shortcuts between libraries
func (l *Library) Sync(Shortcuts *shortcuts.Library) error {

	// Check if can sync library
	if l.BasePath == "" || l.AccountId == "" {
		return nil
	}

	// Generate a new internal shortcut from Steam shortcut
	// Used to sync Steam shortcuts into main library
	toInternal := func(shortcut *Shortcut) *Internal {

		internal := &Internal{}
		internal.ID = shortcuts.FromUint(shortcut.AppID)
		internal.Name = shortcut.AppName
		internal.StartDirectory = shortcut.StartDir
		internal.Executable = CleanExec(shortcut.Exe)
		internal.LaunchOptions = shortcut.LaunchOptions
		internal.ShortcutPath = shortcut.ShortcutPath
		internal.Tags = shortcut.Tags

		// Extended specs
		// @deprecated and will be removed in future versions
		internal.Description = shortcut.Description
		internal.RelativePath = shortcut.RelativePath
		internal.IconURL = shortcut.IconURL
		internal.LogoURL = shortcut.LogoURL
		internal.CoverURL = shortcut.CoverURL
		internal.BannerURL = shortcut.BannerURL
		internal.HeroURL = shortcut.HeroURL

		return internal
	}

	// Generate a new Steam shortcut from internal shortcut
	// Used to sync internal shortcuts into Steam library
	fromInternal := func(internal *Internal) *Shortcut {

		shortcut := &Shortcut{}
		shortcut.AppID = shortcuts.ToUint(internal.ID)
		shortcut.AppName = internal.Name
		shortcut.StartDir = internal.StartDirectory
		shortcut.Exe = EnsureExec(l.Runtime, internal.Executable)
		shortcut.LaunchOptions = internal.LaunchOptions
		shortcut.ShortcutPath = internal.ShortcutPath
		shortcut.Tags = internal.Tags

		return shortcut
	}

	// Sync Steam shortcuts to main library
	// Libraries must have at least one minute difference between timestamps
	// In such case, Steam is considered as newest library reference
	if l.Timestamp > Shortcuts.Timestamp && l.Timestamp > Shortcuts.Timestamp-60 {
		cli.Debug("Synchronizing Steam library to main library.\n")

		// Add or update shortcuts to sync
		processed := make(map[string]bool)
		for _, shortcut := range l.Shortcuts {
			internal := toInternal(shortcut)
			existing := Shortcuts.Get(internal.ID)

			if existing.ID != "" {
				existing.Merge(internal)
				internal = existing
			}

			err := Shortcuts.Set(internal, false)
			if err != nil {
				return err
			}

			processed[internal.ID] = true
		}

		// Remove shortcuts based on processed ones to sync
		for _, internal := range Shortcuts.All() {
			if _, ok := processed[internal.ID]; !ok {
				err := Shortcuts.Remove(internal)
				if err != nil {
					return err
				}
			}
		}

	}

	// Sync main library to Steam shortcuts
	// Libraries must have at least one minute difference between timestamps
	// In such case, main library is considered as newest library reference
	if Shortcuts.Timestamp > l.Timestamp && Shortcuts.Timestamp > l.Timestamp-60 {
		cli.Debug("Synchronizing main library to Steam library.\n")

		// Add or update shortcuts to sync
		processed := make(map[uint]bool)
		for _, reference := range Shortcuts.All() {
			shortcut := fromInternal(reference)
			found := false

			// First, check if already exists the same shortcut
			// This will prevent double additions
			for index, existing := range l.Shortcuts {
				if existing.AppID != shortcut.AppID {
					continue
				}

				cli.Debug("Updating Steam shortcut: %v\n", shortcut.AppID)

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
				processed[shortcut.AppID] = true
				found = true
				break
			}

			// Not found in library, can be added safely
			if !found {

				cli.Debug("Adding Steam shortcut: %v\n", shortcut.AppID)

				// Handle shortcut assets
				err := l.Assets(reference, shortcut, "sync", true)
				if err != nil {
					return err
				}

				// Add shortcut to the library
				l.Shortcuts = append(l.Shortcuts, shortcut)
				processed[shortcut.AppID] = true

			}

		}

		// Remove shortcuts based on processed ones to sync
		for index, shortcut := range l.Shortcuts {
			if _, ok := processed[shortcut.AppID]; !ok {

				cli.Debug("Removing Steam shortcut: %v\n", shortcut.AppID)

				// Handle shortcut assets
				reference := toInternal(shortcut)
				err := l.Assets(reference, shortcut, "remove", true)
				if err != nil {
					return err
				}

				// Empty the shortcut at index
				// We will remove empty shortcuts later
				l.Shortcuts[index] = &Shortcut{}

			}
		}

		// Remove empty shortcuts if any
		// Method is optimized to remove multiple items from slice
		updated := make([]*Shortcut, 0)
		for _, shortcut := range l.Shortcuts {
			if shortcut.AppID != 0 {
				updated = append(updated, shortcut)
			}
		}

		l.Shortcuts = updated

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
