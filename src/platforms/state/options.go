package state

// Options struct
type Options struct {
	Action      string   `json:"action"`
	Platforms   []string `json:"platforms"`
	Preferences []string `json:"preferences"`
}

// Transform values into valid options
func ToOptions(action string, platforms []string, preferences []string) *Options {

	options := Options{
		Action:      action,
		Platforms:   platforms,
		Preferences: preferences,
	}

	return &options
}
