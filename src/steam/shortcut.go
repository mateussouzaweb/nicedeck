package steam

import "github.com/mateussouzaweb/nicedeck/src/shortcuts"

// Shortcut struct
type Shortcut struct {
	AppID               uint     `json:"appId"`
	AppName             string   `json:"appName"`
	StartDir            string   `json:"startDir"`
	Exe                 string   `json:"exe"`
	LaunchOptions       string   `json:"launchOptions"`
	ShortcutPath        string   `json:"shortcutPath"`
	Icon                string   `json:"icon"`
	IsHidden            uint     `json:"isHidden"`
	AllowDesktopConfig  uint     `json:"allowDesktopConfig"`
	AllowOverlay        uint     `json:"allowOverlay"`
	OpenVR              uint     `json:"openVr"`
	DevKit              uint     `json:"devkit"`
	DevKitGameID        string   `json:"devkitGameId"`
	DevKitOverrideAppID uint     `json:"devkitOverrideAppId"`
	LastPlayTime        uint     `json:"lastPlayTime"`
	Tags                []string `json:"tags"`
}

// Perform action when creating the shortcut
func (s *Shortcut) OnCreate() error {

	// Ensure to have a valid shortcut appId
	if s.AppID == 0 {
		appID := shortcuts.GenerateID(s.AppName, s.Exe)
		s.AppID = shortcuts.ToUint(appID)
	}

	return nil
}

// Perform action when updating the shortcut
func (s *Shortcut) OnUpdate() error {
	return nil
}

// Perform action when removing the shortcut
func (s *Shortcut) OnRemove() error {
	return nil
}
