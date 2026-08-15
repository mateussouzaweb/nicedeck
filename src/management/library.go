package management

import "github.com/mateussouzaweb/nicedeck/src/library"

// Init library by setting environment paths
func InitLibrary() error {
	return library.Init()
}

// Load library from config path
func LoadLibrary() error {
	return library.Load()
}

// Save library on config path
func SaveLibrary() error {
	return library.Save()
}

// Sync libraries to add, update or remove entries
func SyncLibrary() error {
	return library.Sync()
}
