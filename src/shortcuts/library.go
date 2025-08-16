package shortcuts

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Library struct
type Library struct {
	SourceFile string      `json:"sourceFile"`
	ImagesPath string      `json:"imagesPath"`
	Shortcuts  []*Shortcut `json:"shortcuts"`
}

// Merge callback with rules to apply between target and source
type MergeCallback func(target *Shortcut, source *Shortcut)

// Load library from file
func (l *Library) Load(sourceFile string) error {

	// Fill basic information
	l.SourceFile = sourceFile
	l.ImagesPath = fmt.Sprintf("%s/images", filepath.Dir(sourceFile))
	l.Shortcuts = make([]*Shortcut, 0)

	// Check if file exist
	exist, err := fs.FileExist(sourceFile)
	if err != nil {
		return err
	} else if !exist {
		return nil
	}

	// Read file content
	content, err := os.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	// Retrieve information from file content when available
	err = json.Unmarshal(content, &l.Shortcuts)
	if err != nil {
		return err
	}

	return nil
}

// Save library of shortcuts into file
func (l *Library) Save() error {

	// Sort data before saving
	err := l.Sort()
	if err != nil {
		return err
	}

	// Convert config state to JSON representation
	jsonContent, err := json.MarshalIndent(l.Shortcuts, "", "  ")
	if err != nil {
		return err
	}

	// Make sure destination folder path exist
	err = os.MkdirAll(filepath.Dir(l.SourceFile), 0774)
	if err != nil {
		return err
	}

	// Write JSON content to config file
	err = os.WriteFile(l.SourceFile, jsonContent, 0666)
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

// Add shortcut to the library
func (l *Library) Add(shortcut *Shortcut) error {

	shortcut.ImagesPath = l.ImagesPath
	err := shortcut.OnCreate()
	if err != nil {
		return err
	}

	l.Shortcuts = append(l.Shortcuts, shortcut)
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

// Update shortcut on library
func (l *Library) Update(shortcut *Shortcut) error {

	// Check if already exist an app with the same reference
	found := false
	for index, item := range l.Shortcuts {
		if item.ID != shortcut.ID {
			continue
		}

		// Run callback on shortcut
		shortcut.ImagesPath = l.ImagesPath
		err := shortcut.OnUpdate()
		if err != nil {
			return err
		}

		// Replace with new object data
		l.Shortcuts[index] = shortcut
		found = true
		break
	}

	// Append to the library if not exist
	if !found {
		return fmt.Errorf("shortcut not found")
	}

	return nil
}

// Add or update shortcut into library
func (l *Library) AddOrUpdate(shortcut *Shortcut) error {

	existing := l.Find(shortcut.Name, shortcut.Executable)
	if existing.ID != "" {
		shortcut.ID = existing.ID
		return l.Update(shortcut)
	}

	return l.Add(shortcut)
}

// Merge shortcuts libraries into one
func (l *Library) Merge(extra []*Shortcut, callback MergeCallback) error {

	// When match is detected, call callback to merge data
	// When has not match, append the item to the list
	for _, item := range extra {
		found := false
		for _, existing := range l.Shortcuts {
			if existing.ID != item.ID {
				continue
			}

			// Run merge callback
			callback(existing, item)

			// Run internal callback on shortcut
			existing.ImagesPath = l.ImagesPath
			err := existing.OnMerge()
			if err != nil {
				return err
			}

			found = true
			break
		}

		// Append to the library if not exist
		if !found {
			l.Add(item)
		}
	}

	return nil
}

// Remove shortcut from the library
func (l *Library) Remove(shortcut *Shortcut) error {

	updated := make([]*Shortcut, 0)
	found := false

	// Instead of appending one by one
	// We detect the shortcut to remove and add others in batch
	for index, item := range l.Shortcuts {
		if item.ID != shortcut.ID {
			continue
		}

		shortcut.ImagesPath = l.ImagesPath
		err := shortcut.OnRemove()
		if err != nil {
			return err
		}

		updated = append(updated, l.Shortcuts[:index]...)
		updated = append(updated, l.Shortcuts[index+1:]...)
		found = true
		break
	}

	// If found, then update the library
	if found {
		l.Shortcuts = updated
	}

	return nil
}
