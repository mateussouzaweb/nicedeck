package packaging

import (
	"fmt"
	"os"
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
	var err error
	switch s.Format {
	case "file":
		err = s.FromFile()
	case "zip", "tar.gz", "tar.xz", "7z":
		err = s.FromArchive()
	case "dmg":
		err = s.FromDMG()
	}
	if err != nil {
		return err
	}

	// Validate if download was success
	installed, err := target.Installed()
	if err != nil {
		return err
	} else if !installed {
		return fmt.Errorf("expected content not detected: %s", s.Destination)
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

// Download source from archive file and extract the content
func (s *Source) FromArchive() error {

	// Gather information
	parentFolder := filepath.Dir(s.Destination)
	extractFolder := filepath.Join(parentFolder, ".extract")
	targetFile := strings.TrimPrefix(s.Destination, parentFolder)
	targetName := strings.TrimSuffix(targetFile, filepath.Ext(targetFile))
	archiveName := fmt.Sprintf("%s.%s", targetName, s.Format)
	archiveFile := filepath.Join(parentFolder, archiveName)

	// Download file
	err := s.SafeDownload(s.URL, archiveFile)
	if err != nil {
		return err
	}

	// Create temporary extract folder
	err = os.MkdirAll(extractFolder, 0755)
	if err != nil {
		return err
	}

	// Extract content to temporary extract folder
	switch s.Format {
	case "zip":
		err = fs.ExtractZip(archiveFile, extractFolder)
	case "tar.gz":
		err = fs.ExtractTarGz(archiveFile, extractFolder)
	case "7z":
		err = fs.Extract7z(archiveFile, extractFolder)
	default:
		err = fmt.Errorf("manual extract required")
		cli.Printf(cli.ColorWarn, "Unable to extract from archive file.\n")
		cli.Printf(cli.ColorWarn, "Please manually extract the program.\n")
		cli.Printf(cli.ColorWarn, "Archive file: %s\n", archiveFile)
		cli.Printf(cli.ColorWarn, "Expected executable: %s\n", s.Destination)
	}
	if err != nil {
		return err
	}

	// Detects where the target file is located in the extract folder
	realPath, err := fs.FindPath(extractFolder, targetFile)
	if err != nil {
		return err
	}

	// Copy files from extract to final destination
	// Use copy to avoid losing files
	copyFrom := filepath.Dir(realPath)
	err = fs.CopyDirectory(copyFrom, parentFolder)
	if err != nil {
		return err
	}

	// Remove extract folder
	err = fs.RemoveDirectory(extractFolder)
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

// Download source from .dmg
func (s *Source) FromDMG() error {

	// Download file
	dmgFile := strings.TrimSuffix(s.Destination, filepath.Ext(s.Destination))
	dmgFile = fmt.Sprintf("%s.dmg", dmgFile)
	err := s.SafeDownload(s.URL, dmgFile)
	if err != nil {
		return err
	}

	// Create script to extract from DMG
	destinationFolder := filepath.Dir(s.Destination)
	appFile := filepath.Base(s.Destination)
	script := fmt.Sprintf(`
		export PAGER="cat";
		VOLUME=$(yes | hdiutil attach -nobrowse -readonly %s);
		VOLUME=$(echo "$VOLUME" | sed -n 's/^.*\(\/Volumes\/.*\)$/\1/p');
		cp -R "$VOLUME/%s" "%s/";
		hdiutil detach "$VOLUME";`,
		dmgFile,
		appFile,
		destinationFolder,
	)

	// Run extraction process
	command := cli.Command(script)
	err = cli.Run(command)
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
