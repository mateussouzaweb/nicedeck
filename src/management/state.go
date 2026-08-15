package management

import "github.com/mateussouzaweb/nicedeck/src/platforms/state"

// Sync state of each platform
func SyncState(options *state.Options) error {
	return state.SyncState(options)
}

// Return synchronizable items based on given options
func GetState(options *state.Options) ([]*state.Synchronizable, error) {
	return state.GetSynchronizables(options)
}
