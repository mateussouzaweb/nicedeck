package steam

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
	"github.com/mateussouzaweb/nicedeck/src/packaging/linux"
	"github.com/mateussouzaweb/nicedeck/src/shortcuts"
	"github.com/mateussouzaweb/nicedeck/src/steam/controller"
	"github.com/mateussouzaweb/nicedeck/src/steam/steam"
	"github.com/mateussouzaweb/nicedeck/src/steam/vdf"
)

// Aliases
type Image = shortcuts.Image
type Internal = shortcuts.Shortcut
type Shortcut = steam.Shortcut

// Library struct
type Library struct {
	DatabasePath  string            `json:"databasePath"`
	BasePath      string            `json:"basePath"`
	Runtime       string            `json:"runtime"`
	AccountId     string            `json:"accountId"`
	AccountName   string            `json:"accountName"`
	ConfigPath    string            `json:"configPath"`
	ImagesPath    string            `json:"imagesPath"`
	ShortcutsPath string            `json:"shortcutsPath"`
	Checksums     map[uint32]string `json:"checksums"`
	Timestamps    map[uint32]int64  `json:"timestamps"`
	Shortcuts     []*Shortcut       `json:"-"`
}

// String representation of the library
func (l *Library) String() string {
	return "Steam"
}

// Operation mode of the library
func (l *Library) Mode() string {
	return "FULL"
}

// Init library
func (l *Library) Init(databasePath string) error {
	l.DatabasePath = databasePath
	return nil
}

// Load library from database file
func (l *Library) Load() error {

	// Reset and fill basic information
	l.BasePath = ""
	l.Runtime = ""
	l.AccountId = ""
	l.AccountName = ""
	l.ConfigPath = ""
	l.ImagesPath = ""
	l.ShortcutsPath = ""
	l.Checksums = make(map[uint32]string, 0)
	l.Timestamps = make(map[uint32]int64, 0)
	l.Shortcuts = make([]*Shortcut, 0)

	// Check if Steam is installed
	steamPackage := steam.GetPackage()
	installed, err := steamPackage.Installed()
	if err != nil {
		return err
	} else if !installed {
		cli.Debug("Steam installation not detected\n")
		return nil
	}

	// Read database file content
	// Missing information will be filled below
	err = fs.ReadJSON(l.DatabasePath, &l)
	if err != nil {
		return err
	}

	// Make sure data is valid
	if l.Checksums == nil {
		l.Checksums = make(map[uint32]string, 0)
	}
	if l.Timestamps == nil {
		l.Timestamps = make(map[uint32]int64, 0)
	}

	// Retrieve base path
	l.BasePath, err = steam.GetBasePath()
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
		l.ConfigPath, err = steam.GetConfigPath()
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

	// Load Steam shortcuts from VDF file
	// Shortcuts file can possible be updated by other services or Steam UI
	// Because the library represents the Steam data
	// We always read shortcuts from Steam source data
	data, err := vdf.ParseBinary(l.ShortcutsPath)
	if err != nil {
		return err
	}

	// Map to struct of shortcuts
	// We don't care about positioning
	found := data.Object("shortcuts")
	for key := range found {
		item, ok := found[key].(vdf.Vdf)
		if !ok || item == nil {
			continue
		}

		// Create tag list
		tags := []string{}
		tagValues := item.Object("tags")
		for key := range tagValues {
			tags = append(tags, tagValues.String(key))
		}

		// Convert to manageable shortcut
		// Timestamp will be filled during indexing
		shortcut := Shortcut{
			AppID:               item.Uint32("appid"),
			AppName:             item.String("AppName"),
			StartDir:            item.String("StartDir"),
			Exe:                 item.String("Exe"),
			LaunchOptions:       item.String("LaunchOptions"),
			ShortcutPath:        item.String("ShortcutPath"),
			Icon:                item.String("icon"),
			IsHidden:            item.Uint32("IsHidden"),
			AllowDesktopConfig:  item.Uint32("AllowDesktopConfig"),
			AllowOverlay:        item.Uint32("AllowOverlay"),
			OpenVR:              item.Uint32("OpenVR"),
			DevKit:              item.Uint32("Devkit"),
			DevKitGameID:        item.String("DevkitGameID"),
			DevKitOverrideAppID: item.Uint32("DevkitOverrideAppID"),
			LastPlayTime:        item.Uint32("LastPlayTime"),
			Tags:                tags,
			Timestamp:           0,
		}

		// Append to library shortcuts
		l.Shortcuts = append(l.Shortcuts, &shortcut)
	}

	// Index shortcuts to detect external changes
	// Process also clean stale metadata
	err = l.Index()
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

	// Index shortcuts to update references before saving
	err := l.Index()
	if err != nil {
		return err
	}

	// Save database state to file
	err = fs.WriteJSON(l.DatabasePath, l)
	if err != nil {
		return err
	}

	// Make sure Steam on flatpak has the necessary permission
	steamPackage := steam.GetPackage()
	if _, ok := steamPackage.(*linux.Flatpak); ok {
		err := steamPackage.(*linux.Flatpak).ApplyOverrides()
		if err != nil {
			return fmt.Errorf("could not perform Steam runtime setup: %s", err)
		}
	}

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

	// Save updated Steam shortcuts to VDF file
	err = vdf.WriteBinary(l.ShortcutsPath, data)
	if err != nil {
		return err
	}

	// Write controller template
	controllerTemplatesPaths := filepath.Join(
		l.BasePath, "controller_base", "templates",
	)

	err = controller.WriteTemplates(controllerTemplatesPaths)
	if err != nil {
		return fmt.Errorf("could not perform Steam controller setup: %s", err)
	}

	cli.Printf(cli.ColorNotice, "Note: Steam library has been updated.\n")
	cli.Printf(cli.ColorNotice, "Please restart Steam or the device to changes take effect.\n")

	return nil
}

// Check if library is available to perform operations
func (l *Library) Available() bool {
	return l.BasePath != "" && l.AccountId != ""
}

// Index shortcuts to detect external changes
func (l *Library) Index() error {

	// Skip indexing if library is not available
	if !l.Available() {
		return nil
	}

	// Take each shortcut and make change verification
	// Use internal checksum and timestamp to detect changes
	// Necessary because Steam does not store change timestamp per shortcut
	// Finally, clean up any checksum and timestamp that no longer exist
	processed := make(map[uint32]bool, 0)
	timestamp := time.Now().UTC().Unix()

	for _, shortcut := range l.Shortcuts {
		appID := shortcut.AppID
		checksum := shortcut.Checksum()
		processed[appID] = true

		// When checksum exists and differ, shortcut was updated
		// When checksum does not exist, consider as new shortcut
		if value, ok := l.Checksums[appID]; ok {
			if value != checksum {
				l.Checksums[appID] = checksum
				l.Timestamps[appID] = timestamp
			}
		} else {
			l.Checksums[appID] = checksum
			l.Timestamps[appID] = timestamp
		}
	}

	// Remove entries that no longer exist
	for appID := range l.Checksums {
		if _, ok := processed[appID]; !ok {
			delete(l.Checksums, appID)
			delete(l.Timestamps, appID)
		}
	}

	// Update timestamp in all shortcuts
	for _, shortcut := range l.Shortcuts {
		shortcut.Timestamp = l.Timestamps[shortcut.AppID]
	}

	return nil
}

// Export shortcuts to internal format
func (l *Library) Export() []*Internal {

	results := make([]*Internal, 0)

	// Skip export if library is not available
	if !l.Available() {
		return results
	}

	// Generate a new internal shortcut from Steam shortcut
	// Used to sync Steam shortcuts into internal library
	for _, shortcut := range l.Shortcuts {
		internal := &Internal{}
		internal.ID = shortcuts.FromUint(uint(shortcut.AppID))
		internal.Name = shortcut.AppName
		internal.StartDirectory = shortcut.StartDir
		internal.Executable = CleanExec(shortcut.Exe)
		internal.LaunchOptions = shortcut.LaunchOptions
		internal.Tags = shortcut.Tags
		internal.Timestamp = shortcut.Timestamp

		results = append(results, internal)
	}

	return results
}

// Transform reference shortcut into Steam shortcut
func (l *Library) Transform(reference *Internal) (*Shortcut, error) {

	shortcut := &Shortcut{}
	shortcut.AppID = uint32(shortcuts.ToUint(reference.ID))
	shortcut.AppName = reference.Name
	shortcut.StartDir = reference.StartDirectory
	shortcut.Exe = EnsureExec(l.Runtime, reference.Executable)
	shortcut.LaunchOptions = reference.LaunchOptions
	shortcut.Tags = reference.Tags
	shortcut.Timestamp = reference.Timestamp

	return shortcut, nil
}

// Check if shortcut is supported by the library
func (l *Library) Supports(reference *Internal) bool {
	return reference.Platform != "Steam"
}

// Add shortcut to the library
func (l *Library) Add(reference *Internal) error {

	// Skip if library is not available
	if !l.Available() {
		return nil
	}

	// Transform reference shortcut into Steam shortcut
	shortcut, err := l.Transform(reference)
	if err != nil {
		return err
	}

	cli.Debug("Adding shortcut to Steam: %v\n", shortcut.AppID)

	// Handle assets only for custom shortcuts
	err = l.Assets(reference, shortcut, "sync", true)
	if err != nil {
		return err
	}

	l.Shortcuts = append(l.Shortcuts, shortcut)
	l.Checksums[shortcut.AppID] = shortcut.Checksum()
	l.Timestamps[shortcut.AppID] = shortcut.Timestamp

	return nil
}

// Update shortcut on library
func (l *Library) Update(reference *Internal, overwriteAssets bool) error {

	// Skip if library is not available
	if !l.Available() {
		return nil
	}

	// Transform reference shortcut into Steam shortcut
	shortcut, err := l.Transform(reference)
	if err != nil {
		return err
	}

	// Find and update existing shortcut
	for index, existing := range l.Shortcuts {
		if existing.AppID != shortcut.AppID {
			continue
		}

		cli.Debug("Updating shortcut in Steam: %v\n", existing.AppID)

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
		err := l.Assets(reference, shortcut, "sync", overwriteAssets)
		if err != nil {
			return err
		}

		// Replace shortcut at index and update references
		l.Shortcuts[index] = shortcut
		l.Checksums[shortcut.AppID] = shortcut.Checksum()
		l.Timestamps[shortcut.AppID] = shortcut.Timestamp

		break
	}

	return nil
}

// Set shortcut on library to either add or update
func (l *Library) Set(reference *Internal, overwriteAssets bool) error {

	// Skip if library is not available
	if !l.Available() {
		return nil
	}

	// Transform reference shortcut into Steam shortcut
	shortcut, err := l.Transform(reference)
	if err != nil {
		return err
	}

	// Use checksums presence for faster lookup
	if _, ok := l.Checksums[shortcut.AppID]; !ok {
		return l.Add(reference)
	} else {
		return l.Update(reference, overwriteAssets)
	}
}

// Remove shortcut from the library
func (l *Library) Remove(reference *Internal) error {

	// Skip if library is not available
	if !l.Available() {
		return nil
	}

	// Transform reference shortcut into Steam shortcut
	shortcut, err := l.Transform(reference)
	if err != nil {
		return err
	}

	// Find and remove shortcut
	for index, existing := range l.Shortcuts {
		if existing.AppID != shortcut.AppID {
			continue
		}

		cli.Debug("Removing shortcut in Steam: %v\n", shortcut.AppID)

		// Handle shortcut assets
		err := l.Assets(reference, existing, "remove", true)
		if err != nil {
			return err
		}

		// Update library of shortcuts
		// Method is optimized to remove single item from slice
		updated := make([]*Shortcut, 0)
		updated = append(updated, l.Shortcuts[:index]...)
		updated = append(updated, l.Shortcuts[index+1:]...)
		l.Shortcuts = updated
		break
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
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%v_icon", shortcut.AppID),
		Extensions:      []string{".png", ".ico"},
	}
	logoImage := &Image{
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%v_logo", shortcut.AppID),
		Extensions:      []string{".png"},
	}
	coverImage := &Image{
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%vp", shortcut.AppID),
		Extensions:      []string{".png", ".jpg"},
	}
	bannerImage := &Image{
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%v", shortcut.AppID),
		Extensions:      []string{".png", ".jpg"},
	}
	heroImage := &Image{
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%v_hero", shortcut.AppID),
		Extensions:      []string{".png", ".jpg"},
	}

	// Sync all images based on the action
	if action == "sync" || action == "add" {
		err := iconImage.Process(specs.IconPath, overwrite)
		if err != nil {
			return err
		}
		err = logoImage.Process(specs.LogoPath, overwrite)
		if err != nil {
			return err
		}
		err = coverImage.Process(specs.CoverPath, overwrite)
		if err != nil {
			return err
		}
		err = bannerImage.Process(specs.BannerPath, overwrite)
		if err != nil {
			return err
		}
		err = heroImage.Process(specs.HeroPath, overwrite)
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
