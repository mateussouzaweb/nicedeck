package steam

import (
	"crypto/md5"
	"fmt"
	"slices"
)

// Shortcut struct
type Shortcut struct {
	AppID               uint32   `json:"appId"`
	AppName             string   `json:"appName"`
	StartDir            string   `json:"startDir"`
	Exe                 string   `json:"exe"`
	LaunchOptions       string   `json:"launchOptions"`
	ShortcutPath        string   `json:"shortcutPath"`
	Icon                string   `json:"icon"`
	IsHidden            uint32   `json:"isHidden"`
	AllowDesktopConfig  uint32   `json:"allowDesktopConfig"`
	AllowOverlay        uint32   `json:"allowOverlay"`
	OpenVR              uint32   `json:"openVr"`
	DevKit              uint32   `json:"devKit"`
	DevKitGameID        string   `json:"devKitGameId"`
	DevKitOverrideAppID uint32   `json:"devKitOverrideAppId"`
	LastPlayTime        uint32   `json:"lastPlayTime"`
	Tags                []string `json:"tags"`
	Timestamp           int64    `json:"timestamp"`
}

// Generate checksum for shortcut from relevant fields
func (s *Shortcut) Checksum() string {

	// Make sure to sort tags before generating checksum
	slices.Sort(s.Tags)

	// Create data string from relevant fields
	data := fmt.Sprintf(
		"%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v",
		s.AppName,
		s.StartDir,
		s.Exe,
		s.LaunchOptions,
		s.ShortcutPath,
		s.Icon,
		s.IsHidden,
		s.AllowDesktopConfig,
		s.AllowOverlay,
		s.OpenVR,
		s.DevKit,
		s.DevKitGameID,
		s.DevKitOverrideAppID,
		s.Tags,
	)

	checksum := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", checksum)
}
