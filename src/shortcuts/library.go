package shortcuts

import (
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
	l.ImagesPath = filepath.Join(filepath.Dir(databasePath), "images")
	l.Shortcuts = make([]*Shortcut, 0)
	l.History = make([]*History, 0)

	// Read database file content
	err := fs.ReadJSON(databasePath, &l)
	if err != nil {
		return err
	}

	// Make sure images path exists
	err = os.MkdirAll(l.ImagesPath, 0755)
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

	// Save database state to file
	err = fs.WriteJSON(l.DatabasePath, l)
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

	// Handle shortcut assets
	err := l.Assets(shortcut, "sync", true)
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
func (l *Library) Update(shortcut *Shortcut, overwriteAssets bool) error {

	// Check if already exist an app with the same reference
	found := false
	for index, item := range l.Shortcuts {
		if item.ID != shortcut.ID {
			continue
		}

		// When no changes are detect, don't do anything
		if !reflect.DeepEqual(item, shortcut) {
			return nil
		}

		// Handle shortcut assets
		err := l.Assets(shortcut, "sync", overwriteAssets)
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
func (l *Library) Set(shortcut *Shortcut, overwriteAssets bool) error {
	existing := l.Get(shortcut.ID)

	if existing.ID != "" {
		return l.Update(shortcut, overwriteAssets)
	} else {
		return l.Add(shortcut)
	}
}

// Remove shortcut from the library
func (l *Library) Remove(shortcut *Shortcut) error {

	// Instead of appending one by one
	// We detect the shortcut to remove and add others in batch
	for index, item := range l.Shortcuts {
		if item.ID != shortcut.ID {
			continue
		}

		// Handle shortcut assets
		err := l.Assets(shortcut, "remove", true)
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
func (l *Library) Assets(shortcut *Shortcut, action string, overwrite bool) error {

	// Internal image formats:
	// - Logo: ${ID}_logo.png
	// - Icon: ${ID}_icon.(ico|png)
	// - Cover: ${ID}_cover.(jpg|png)
	// - Banner: ${ID}_banner.(jpg|png)
	// - Hero: ${ID}_hero.(jpg|png)

	// Handle images
	// Process usually mean download image from URL
	iconImage := &Image{
		SourceURL:       shortcut.IconURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_icon", shortcut.ID),
		Extensions:      []string{".png", ".ico"},
	}
	logoImage := &Image{
		SourceURL:       shortcut.LogoURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_logo", shortcut.ID),
		Extensions:      []string{".png"},
	}
	coverImage := &Image{
		SourceURL:       shortcut.CoverURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_cover", shortcut.ID),
		Extensions:      []string{".png", ".jpg"},
	}
	bannerImage := &Image{
		SourceURL:       shortcut.BannerURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_banner", shortcut.ID),
		Extensions:      []string{".png", ".jpg"},
	}
	heroImage := &Image{
		SourceURL:       shortcut.HeroURL,
		TargetDirectory: l.ImagesPath,
		TargetName:      fmt.Sprintf("%s_hero", shortcut.ID),
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
