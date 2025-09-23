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

	cli.Printf(cli.ColorNotice, "Downloading: %s\n", s.URL)
	if s.Destination == "" {
		s.Destination = target.Executable()
	}

	// Download based on format
	switch s.Format {
	case "file":
		return s.FromFile()
	case "zip":
		return s.FromZip()
	case "tar.gz":
		return s.FromTarGz()
	case "tar.xz":
		return s.FromTarXz()
	case "7z":
		return s.From7z()
	case "dmg":
		return s.FromDMG()
	}

	return nil
}

// Safe download to destination file to avoid collision
func (s *Source) SafeDownload(url string, destination string) error {

	// Check if destination file exists
	// When not exists, just download the new file
	exist, err := fs.FileExist(destination)
	if err != nil {
		return err
	} else if !exist {
		return fs.DownloadFile(url, destination, false)
	}

	// Perform safe download operation with the following process
	// - Remove existing .tmp
	// - Download new file to .tmp
	// - Remove existing .old file
	// - Rename existing file to .old
	// - Rename .tmp to final destination
	tmpDestination := fmt.Sprintf("%s.tmp", destination)
	oldDestination := fmt.Sprintf("%s.old", destination)

	err = fs.RemoveFile(tmpDestination)
	if err != nil {
		return err
	}

	err = fs.DownloadFile(url, tmpDestination, false)
	if err != nil {
		return err
	}

	err = fs.RemoveFile(oldDestination)
	if err != nil {
		return err
	}

	err = fs.MoveFile(destination, oldDestination)
	if err != nil {
		return err
	}

	err = fs.MoveFile(tmpDestination, destination)
	if err != nil {
		return err
	}

	return nil
}

// Download source from direct file
func (s *Source) FromFile() error {

	// Download the file to destination
	err := s.SafeDownload(s.URL, s.Destination)
	if err != nil {
		return err
	}

	return nil
}

// Download source from .zip
func (s *Source) FromZip() error {

	// Download file
	archiveFile := strings.TrimSuffix(s.Destination, filepath.Ext(s.Destination))
	archiveFile = fmt.Sprintf("%s.zip", archiveFile)
	err := s.SafeDownload(s.URL, archiveFile)
	if err != nil {
		return err
	}

	// Extract content to parent folder
	parentFolder := filepath.Dir(s.Destination)
	targetFile := strings.TrimPrefix(s.Destination, parentFolder)
	err = fs.ExtractZip(archiveFile, parentFolder, targetFile)
	if err != nil {
		return err
	}

	// Remove archive file
	err = fs.RemoveFile(archiveFile)
	if err != nil {
		return err
	}

	return nil
}

// Download source from .tar.gz
func (s *Source) FromTarGz() error {

	// Download file
	archiveFile := strings.TrimSuffix(s.Destination, filepath.Ext(s.Destination))
	archiveFile = fmt.Sprintf("%s.tar.gz", archiveFile)
	err := s.SafeDownload(s.URL, archiveFile)
	if err != nil {
		return err
	}

	// Extract content to parent folder
	parentFolder := filepath.Dir(s.Destination)
	targetFile := strings.TrimPrefix(s.Destination, parentFolder)
	err = fs.ExtractTarGz(archiveFile, parentFolder, targetFile)
	if err != nil {
		return err
	}

	// Remove archive file
	err = fs.RemoveFile(archiveFile)
	if err != nil {
		return err
	}

	return nil
}

// Download source from .tar.xz
func (s *Source) FromTarXz() error {

	// Download file
	archiveFile := strings.TrimSuffix(s.Destination, filepath.Ext(s.Destination))
	archiveFile = fmt.Sprintf("%s.tar.xz", archiveFile)
	err := s.SafeDownload(s.URL, archiveFile)
	if err != nil {
		return err
	}

	// Print warning message
	cli.Printf(cli.ColorWarn, "WARNING: Unable to extract from .tar.xz file.\n")
	cli.Printf(cli.ColorWarn, "Please manually extract the program.\n")
	cli.Printf(cli.ColorWarn, "Archive file: %s\n", archiveFile)
	cli.Printf(cli.ColorWarn, "Expected executable: %s\n", s.Destination)

	return nil
}

// Download source from .7z
func (s *Source) From7z() error {

	// Download file
	archiveFile := strings.TrimSuffix(s.Destination, filepath.Ext(s.Destination))
	archiveFile = fmt.Sprintf("%s.7z", archiveFile)
	err := s.SafeDownload(s.URL, archiveFile)
	if err != nil {
		return err
	}

	// Print warning message
	cli.Printf(cli.ColorWarn, "WARNING: Unable to extract from .7z file.\n")
	cli.Printf(cli.ColorWarn, "Please manually extract the program.\n")
	cli.Printf(cli.ColorWarn, "Archive file: %s\n", archiveFile)
	cli.Printf(cli.ColorWarn, "Expected executable: %s\n", s.Destination)

	return nil
}

// Download source from .dmg
func (s *Source) FromDMG() error {

	// Download file
	dmgFile := strings.TrimSuffix(s.Destination, filepath.Ext(s.Destination))
	dmgFile = fmt.Sprintf("%s.dmg", dmgFile)
	err := s.SafeDownload(s.URL, dmgFile)
	if err != nil {
		return err
	}

	// Print warning message
	cli.Printf(cli.ColorWarn, "WARNING: Unable to extract from .dmg file.\n")
	cli.Printf(cli.ColorWarn, "Please manually extract the program.\n")
	cli.Printf(cli.ColorWarn, "DMG file: %s\n", dmgFile)
	cli.Printf(cli.ColorWarn, "Expected executable: %s\n", s.Destination)

	return nil
}

// Find expected file format based on filename
func FindFormat(name string) string {

	if strings.HasSuffix(name, ".zip") {
		return "zip"
	} else if strings.HasSuffix(name, ".tar.gz") {
		return "tar.gz"
	} else if strings.HasSuffix(name, ".tar.xz") {
		return "tar.xz"
	} else if strings.HasSuffix(name, ".7z") {
		return "7z"
	} else if strings.HasSuffix(name, ".dmg") {
		return "dmg"
	}

	return "file"
}
