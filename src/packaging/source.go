package packaging

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/cli"
	"github.com/mateussouzaweb/nicedeck/src/fs"
)

type Resolver func() (string, error)

type Source struct {
	URL         string   `json:"url"`
	Destination string   `json:"destination"`
	Format      string   `json:"format"`
	Resolver    Resolver `json:"-"`
}

// Download content and extract it source into target
func (s *Source) Download(target Package) error {

	// URL can be retrieved from:
	// - Direct link in URL field
	// - Custom method when resolver is defined
	if s.URL == "" && s.Resolver != nil {
		url, err := s.Resolver()
		s.URL = url
		if err != nil {
			return err
		}
	}

	// Download based on format
	s.Destination = target.Executable()
	switch s.Format {
	case "file":
		return s.FromFile()
	case "zip":
		return s.FromZip()
	case "dmg":
		return s.FromDMG()
	}

	return nil
}

// Download source from direct file
func (s *Source) FromFile() error {
	err := fs.DownloadFile(s.URL, s.Destination, true)
	if err != nil {
		return err
	}

	return nil
}

// Download source from zip
func (s *Source) FromZip() error {

	// Download Zip
	zipFile := fmt.Sprintf("%s.zip", s.Destination)
	err := fs.DownloadFile(s.URL, zipFile, true)
	if err != nil {
		return err
	}

	// Extract ZIP
	err = fs.Unzip(zipFile, s.Destination)
	if err != nil {
		return err
	}

	// Remove ZIP file
	err = fs.RemoveFile(zipFile)
	if err != nil {
		return err
	}

	return nil
}

// Download source from DMG
func (s *Source) FromDMG() error {

	// Download file
	dmgFile := strings.Replace(s.Destination, ".app", ".dmg", 1)
	err := fs.DownloadFile(s.URL, dmgFile, true)
	if err != nil {
		return err
	}

	// Extract application from DMG
	appName := filepath.Base(s.Destination)
	script := fmt.Sprintf(
		`hdiutil attach %s
		cp /Volumes/%s %s
		hdiutil detach /Volumes/%s`,
		dmgFile,
		appName,
		s.Destination,
		appName,
	)

	err = cli.Run(script)
	if err != nil {
		return err
	}

	// Remove DMG file
	err = fs.RemoveFile(dmgFile)
	if err != nil {
		return err
	}

	return nil
}
