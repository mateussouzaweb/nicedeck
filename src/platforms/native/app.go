package native

// App struct
type App struct {
	Reference      string `json:"reference"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Runtime        string `json:"runtime"`
	Platform       string `json:"platform"`
	Path           string `json:"path"`
	Directory      string `json:"directory"`
	File           string `json:"file"`
	Extension      string `json:"extension"`
	StartDirectory string `json:"startDirectory"`
	Executable     string `json:"executable"`
	LaunchOptions  string `json:"launchOptions"`
}
