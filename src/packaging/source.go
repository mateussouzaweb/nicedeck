package packaging

import (
	"fmt"

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
	case "tar.gz":
		return s.FromTarGZ()
	case "zip":
		return s.FromZip()
	case "7z":
		return s.From7z()
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

// Download source from .tar.gz
func (s *Source) FromTarGZ() error {

	// Download file
	archiveFile := fmt.Sprintf("%s.tar.gz", s.Destination)
	err := fs.DownloadFile(s.URL, archiveFile, true)
	if err != nil {
		return err
	}

	// Print warning message
	cli.Printf(cli.ColorWarn, "WARNING\n")
	cli.Printf(cli.ColorWarn, "Unable to extract from .tar.gz file.\n")
	cli.Printf(cli.ColorWarn, "Please manually extract the program.\n")
	cli.Printf(cli.ColorWarn, "Archive file: %s\n", archiveFile)
	cli.Printf(cli.ColorWarn, "Expected executable: %s\n", s.Destination)

	return nil
}

// Download source from .zip
func (s *Source) FromZip() error {

	// Download Zip
	archiveFile := fmt.Sprintf("%s.zip", s.Destination)
	err := fs.DownloadFile(s.URL, archiveFile, true)
	if err != nil {
		return err
	}

	// Print warning message
	cli.Printf(cli.ColorWarn, "WARNING\n")
	cli.Printf(cli.ColorWarn, "Unable to extract from .zip file.\n")
	cli.Printf(cli.ColorWarn, "Please manually extract the program.\n")
	cli.Printf(cli.ColorWarn, "Archive file: %s\n", archiveFile)
	cli.Printf(cli.ColorWarn, "Expected executable: %s\n", s.Destination)

	return nil
}

// Download source from .7z
func (s *Source) From7z() error {

	// Download file
	archiveFile := fmt.Sprintf("%s.7z", s.Destination)
	err := fs.DownloadFile(s.URL, archiveFile, true)
	if err != nil {
		return err
	}

	// Print warning message
	cli.Printf(cli.ColorWarn, "WARNING\n")
	cli.Printf(cli.ColorWarn, "Unable to extract from .7z file.\n")
	cli.Printf(cli.ColorWarn, "Please manually extract the program.\n")
	cli.Printf(cli.ColorWarn, "Archive file: %s\n", archiveFile)
	cli.Printf(cli.ColorWarn, "Expected executable: %s\n", s.Destination)

	return nil
}

// Download source from .dmg
func (s *Source) FromDMG() error {

	// Download file
	dmgFile := fmt.Sprintf("%s.dmg", s.Destination)
	err := fs.DownloadFile(s.URL, dmgFile, true)
	if err != nil {
		return err
	}

	// Print warning message
	cli.Printf(cli.ColorWarn, "WARNING\n")
	cli.Printf(cli.ColorWarn, "Unable to extract from .dmg file.\n")
	cli.Printf(cli.ColorWarn, "Please manually extract the program.\n")
	cli.Printf(cli.ColorWarn, "DMG file: %s\n", dmgFile)
	cli.Printf(cli.ColorWarn, "Expected executable: %s\n", s.Destination)

	return nil
}
