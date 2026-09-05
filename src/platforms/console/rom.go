package console

// ROM struct
type ROM struct {
	Reference      string `json:"reference"`
	Path           string `json:"path"`
	Directory      string `json:"directory"`
	File           string `json:"file"`
	Extension      string `json:"extension"`
	Name           string `json:"name"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Platform       string `json:"platform"`
	Console        string `json:"console"`
	Emulator       string `json:"emulator"`
	Program        string `json:"program"`
	StartDirectory string `json:"startDirectory"`
	Executable     string `json:"executable"`
	LaunchOptions  string `json:"launchOptions"`
}
