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

// Add shortcut to the library
func (l *Library) Add(shortcut *Shortcut) error {

	shortcut.ImagesPath = l.ImagesPath
	err := shortcut.OnCreate()
	if err != nil {
		return err
	}

	l.Shortcuts = append(l.Shortcuts, shortcut)
	l.History = append(l.History, &History{
		Action:   "added",
		Original: &Shortcut{},
		Result:   shortcut,
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

		// When no changes are detect, don't do anything
		if !reflect.DeepEqual(item, shortcut) {
			return nil
		}

		// Replace object at index and generate history
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

// Add or update shortcut into library
func (l *Library) AddOrUpdate(shortcut *Shortcut) error {

	existing := l.Find(shortcut.Name, shortcut.Executable)
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

		shortcut.ImagesPath = l.ImagesPath
		err := shortcut.OnRemove()
		if err != nil {
			return err
		}

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
