package game

// Game struct
type Game struct {
	Reference      string `json:"reference"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Platform       string `json:"platform"`
	Launcher       string `json:"launcher"`
	Program        string `json:"program"`
	StartDirectory string `json:"startDirectory"`
	Executable     string `json:"executable"`
	LaunchOptions  string `json:"launchOptions"`
}
