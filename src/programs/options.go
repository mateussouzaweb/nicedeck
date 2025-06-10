package programs

// Options struct
type Options struct {
	Programs    []string `json:"programs"`
	Preferences []string `json:"preferences"`
}

// Transform values into valid options
func ToOptions(programs []string, preferences []string) *Options {

	options := Options{
		Programs:    programs,
		Preferences: preferences,
	}

	return &options
}
