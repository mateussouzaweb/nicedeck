package shortcuts

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// History struct
type History struct {
	Action   string
	Original *Shortcut
	Result   *Shortcut
}

// Library struct
type Library struct {
	DatabasePath string      `json:"databasePath"`
	ImagesPath   string      `json:"imagesPath"`
	Shortcuts    []*Shortcut `json:"shortcuts"`
	History      []*History  `json:"-"`
}

// Load library from database file
func (l *Library) Load(databasePath string) error {

	// Fill basic information
	l.DatabasePath = databasePath
	l.ImagesPath = fmt.Sprintf("%s/images", filepath.Dir(databasePath))
	l.Shortcuts = make([]*Shortcut, 0)
	l.History = make([]*History, 0)

	// Check if database file exist
	exist, err := fs.FileExist(databasePath)
	if err != nil {
		return err
	} else if !exist {
		return nil
	}

	// Read database file content
	content, err := os.ReadFile(databasePath)
	if err != nil {
		return err
	}

	// Retrieve information from database file content when available
	err = json.Unmarshal(content, &l)
	if err != nil {
		return err
	}

	return nil
}

// Save library of shortcuts into database file
func (l *Library) Save() error {

	// Sort library before saving into database
	err := l.Sort()
	if err != nil {
		return err
	}

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

	return nil
}

// Retrieve all shortcuts in the library
func (l *Library) All() []*Shortcut {
	return l.Shortcuts
}

// Sort shortcuts in alphabetical order
func (l *Library) Sort() error {

	sort.Slice(l.Shortcuts, func(i int, j int) bool {
		return l.Shortcuts[i].Name < l.Shortcuts[j].Name
	})

	return nil
}

// Retrieve shortcut in the library with given ID
func (l *Library) Get(ID string) *Shortcut {
	for _, item := range l.Shortcuts {
		if item.ID == ID {
			return item
		}
	}

	return &Shortcut{}
}

// Find shortcut with given name and executable combination
// Values are required and used to determine a shortcut ID
func (l *Library) Find(name string, executable string) *Shortcut {
	ID := GenerateID(name, executable)
	return l.Get(ID)
}

// Add shortcut to the library
func (l *Library) Add(shortcut *Shortcut) error {

	// Run callback on shortcut
	err := shortcut.OnCreate()
	if err != nil {
		return err
	}

	// Handle shortcut assets
	err = l.Assets(shortcut, "add")
	if err != nil {
		return err
	}

	// Add shortcut and generate history
	l.Shortcuts = append(l.Shortcuts, shortcut)
	l.History = append(l.History, &History{
		Action:   "added",
		Original: &Shortcut{},
		Result:   shortcut,
	})

	return nil
}

// Update shortcut on library
func (l *Library) Update(shortcut *Shortcut) error {

	// Check if already exist an app with the same reference
	found := false
	for index, item := range l.Shortcuts {
		if item.ID != shortcut.ID {
			continue
		}

		// Run callback on shortcut
		err := shortcut.OnUpdate()
		if err != nil {
			return err
		}

		// When no changes are detect, don't do anything
		if !reflect.DeepEqual(item, shortcut) {
			return nil
		}

		// Handle shortcut assets
		err = l.Assets(shortcut, "sync")
		if err != nil {
			return err
		}

		// Replace shortcut at index and generate history
		l.Shortcuts[index] = shortcut
		l.History = append(l.History, &History{
			Action:   "updated",
			Original: item,
			Result:   shortcut,
		})

		found = true
		break
	}

	// Append to the library if not exist
	if !found {
		return fmt.Errorf("shortcut not found")
	}

	return nil
}

// Set shortcut into library by adding or updating it
func (l *Library) Set(shortcut *Shortcut) error {

	shortcutID := shortcut.ID
	if shortcutID == "" {
		shortcutID = GenerateID(shortcut.Name, shortcut.Executable)
	}

	existing := l.Get(shortcutID)
	if existing.ID != "" {
		shortcut.ID = existing.ID
		return l.Update(shortcut)
	}

	return l.Add(shortcut)
}

// Remove shortcut from the library
func (l *Library) Remove(shortcut *Shortcut) error {

	// Instead of appending one by one
	// We detect the shortcut to remove and add others in batch
	for index, item := range l.Shortcuts {
		if item.ID != shortcut.ID {
			continue
		}

		// Run callback on shortcut
		err := shortcut.OnRemove()
		if err != nil {
			return err
		}

		// Handle shortcut assets
		err = l.Assets(shortcut, "remove")
		if err != nil {
			return err
		}

		// Update library shortcuts and history
		updated := make([]*Shortcut, 0)
		updated = append(updated, l.Shortcuts[:index]...)
		updated = append(updated, l.Shortcuts[index+1:]...)

		l.Shortcuts = updated
		l.History = append(l.History, &History{
			Action:   "removed",
			Original: shortcut,
			Result:   &Shortcut{},
		})

		break
	}

	return nil
}

// Process assets for shortcut based on action
func (l *Library) Assets(shortcut *Shortcut, action string) error {

	// Internal image formats:
	// - Logo: ${ID}_logo.png
	// - Icon: ${ID}_icon.(ico|png)
	// - Cover: ${ID}_cover.(jpg|png)
	// - Banner: ${ID}_banner.(jpg|png)
	// - Hero: ${ID}_hero.(jpg|png)

	// Handle images
	// Process usually mean download image from URL
	iconImage := &Image{
		SourcePath:      shortcut.CoverPath,
		SourceURL:       shortcut.CoverURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_icon", shortcut.ID),
		Extensions:      []string{".png", ".ico"},
	}
	logoImage := &Image{
		SourcePath:      shortcut.LogoPath,
		SourceURL:       shortcut.LogoURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_logo", shortcut.ID),
		Extensions:      []string{".png"},
	}
	coverImage := &Image{
		SourcePath:      shortcut.CoverPath,
		SourceURL:       shortcut.CoverURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_cover", shortcut.ID),
		Extensions:      []string{".png", ".jpg"},
	}
	bannerImage := &Image{
		SourcePath:      shortcut.BannerPath,
		SourceURL:       shortcut.BannerURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_banner", shortcut.ID),
		Extensions:      []string{".png", ".jpg"},
	}
	heroImage := &Image{
		SourcePath:      shortcut.HeroPath,
		SourceURL:       shortcut.HeroURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_hero", shortcut.ID),
		Extensions:      []string{".png", ".jpg"},
	}

	// Sync all images based on the action
	if action == "sync" || action == "add" {
		overwriteExisting := action == "add"

		err := iconImage.Process(overwriteExisting)
		if err != nil {
			return err
		}
		err = logoImage.Process(overwriteExisting)
		if err != nil {
			return err
		}
		err = coverImage.Process(overwriteExisting)
		if err != nil {
			return err
		}
		err = bannerImage.Process(overwriteExisting)
		if err != nil {
			return err
		}
		err = heroImage.Process(overwriteExisting)
		if err != nil {
			return err
		}

		shortcut.IconPath = iconImage.TargetPath
		shortcut.LogoPath = logoImage.TargetPath
		shortcut.CoverPath = coverImage.TargetPath
		shortcut.BannerPath = bannerImage.TargetPath
		shortcut.HeroPath = heroImage.TargetPath
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
