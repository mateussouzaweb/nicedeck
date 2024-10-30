package shortcuts

import (
	"slices"
	"sort"
)

// Retrieve shortcut in the list with given appID
func GetShortcut(shortcuts []*Shortcut, appID uint) *Shortcut {

	for _, item := range shortcuts {
		if item.AppID == appID {
			return item
		}
	}

	return &Shortcut{}
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

// Merge callback with rules to apply between target and source
type MergeCallback func(target *Shortcut, source *Shortcut)

// Merge shortcuts lists into one
func MergeShortcuts(main []*Shortcut, extra []*Shortcut, mergeCallback MergeCallback, appendWhenNotFound bool) []*Shortcut {

	// When match is detected, call callback to merge data
	// When has not match, it can append the item to the list if option is enabled
	for _, item := range extra {
		found := false
		for _, existing := range main {
			if existing.AppID != item.AppID {
				continue
			}

			mergeCallback(existing, item)
			found = true
			break
		}

		if !found && appendWhenNotFound {
			main = append(main, item)
		}
	}

	return main
}
