package shortcuts

import (
	"fmt"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/fs"
)

// Shortcut struct
type Shortcut struct {
	ID             string   `json:"id"`
	Platform       string   `json:"platform"`
	Program        string   `json:"program"`
	Layer          string   `json:"layer"`
	Type           string   `json:"type"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	StartDirectory string   `json:"startDirectory"`
	Executable     string   `json:"executable"`
	LaunchOptions  string   `json:"launchOptions"`
	ShortcutPath   string   `json:"shortcutPath"`
	RelativePath   string   `json:"relativePath"`
	ImagesPath     string   `json:"imagesPath"`
	IconPath       string   `json:"iconPath"`
	IconURL        string   `json:"iconUrl"`
	LogoPath       string   `json:"logoPath"`
	LogoURL        string   `json:"logoUrl"`
	CoverPath      string   `json:"coverPath"`
	CoverURL       string   `json:"coverUrl"`
	BannerPath     string   `json:"bannerPath"`
	BannerURL      string   `json:"bannerUrl"`
	HeroPath       string   `json:"heroPath"`
	HeroURL        string   `json:"heroUrl"`
	Tags           []string `json:"tags"`
}

// Process images of the shortcut
func (s *Shortcut) ProcessImages(overwriteExisting bool) error {

	toRemove := []string{}

	// Logo: ${ID}_logo.png
	logoPng := fmt.Sprintf("%s/%v_logo.png", s.ImagesPath, s.ID)
	logoPng = fs.NormalizePath(logoPng)

	if s.LogoURL != "" {
		s.LogoPath = logoPng
	} else {
		s.LogoPath = ""
		toRemove = append(toRemove, logoPng)
	}

	// Icon: ${ID}_icon.ico || ${ID}_icon.png
	iconPng := fmt.Sprintf("%s/%v_icon.png", s.ImagesPath, s.ID)
	iconPng = fs.NormalizePath(iconPng)

	iconIco := fmt.Sprintf("%s/%v_icon.ico", s.ImagesPath, s.ID)
	iconIco = fs.NormalizePath(iconIco)

	if strings.HasSuffix(s.IconURL, ".png") {
		s.IconPath = iconPng
		toRemove = append(toRemove, iconIco)
	} else if s.IconURL != "" {
		s.IconPath = iconIco
		toRemove = append(toRemove, iconPng)
	} else {
		s.IconPath = ""
		toRemove = append(toRemove, iconPng)
		toRemove = append(toRemove, iconIco)
	}

	// Cover: ${ID}_cover.png || ${ID}_cover.jpg
	coverPng := fmt.Sprintf("%s/%v_cover.png", s.ImagesPath, s.ID)
	coverPng = fs.NormalizePath(coverPng)

	coverJpg := fmt.Sprintf("%s/%v_cover.jpg", s.ImagesPath, s.ID)
	coverJpg = fs.NormalizePath(coverJpg)

	if strings.HasSuffix(s.CoverURL, ".png") {
		s.CoverPath = coverPng
		toRemove = append(toRemove, coverJpg)
	} else if s.CoverURL != "" {
		s.CoverPath = coverJpg
		toRemove = append(toRemove, coverPng)
	} else {
		s.CoverPath = ""
		toRemove = append(toRemove, coverPng)
		toRemove = append(toRemove, coverJpg)
	}

	// Banner: ${ID}_banner.png || ${ID}_banner.jpg
	bannerPng := fmt.Sprintf("%s/%v_banner.png", s.ImagesPath, s.ID)
	bannerPng = fs.NormalizePath(bannerPng)

	bannerJpg := fmt.Sprintf("%s/%v_banner.jpg", s.ImagesPath, s.ID)
	bannerJpg = fs.NormalizePath(bannerJpg)

	if strings.HasSuffix(s.BannerURL, ".png") {
		s.BannerPath = bannerPng
		toRemove = append(toRemove, bannerJpg)
	} else if s.BannerURL != "" {
		s.BannerPath = bannerJpg
		toRemove = append(toRemove, bannerPng)
	} else {
		s.BannerPath = ""
		toRemove = append(toRemove, bannerPng)
		toRemove = append(toRemove, bannerJpg)
	}

	// Hero: ${ID}_hero.png || ${ID}_hero.jpg
	heroPng := fmt.Sprintf("%s/%v_hero.png", s.ImagesPath, s.ID)
	heroPng = fs.NormalizePath(heroPng)

	heroJpg := fmt.Sprintf("%s/%v_hero.jpg", s.ImagesPath, s.ID)
	heroJpg = fs.NormalizePath(heroJpg)

	if strings.HasSuffix(s.HeroURL, ".png") {
		s.HeroPath = heroPng
		toRemove = append(toRemove, heroJpg)
	} else if s.HeroURL != "" {
		s.HeroPath = heroJpg
		toRemove = append(toRemove, heroPng)
	} else {
		s.HeroPath = ""
		toRemove = append(toRemove, heroPng)
		toRemove = append(toRemove, heroJpg)
	}

	// Remove duplicated or unnecessary images
	for _, file := range toRemove {
		err := fs.RemoveFile(file)
		if err != nil {
			return err
		}
	}

	// Create list of images to download
	images := map[string]string{}
	images[s.IconURL] = s.IconPath
	images[s.LogoURL] = s.LogoPath
	images[s.CoverURL] = s.CoverPath
	images[s.BannerURL] = s.BannerPath
	images[s.HeroURL] = s.HeroPath

	// Download available shortcut images
	for url, destinationFile := range images {
		if url == "" || destinationFile == "" {
			continue
		}
		err := fs.DownloadFile(url, destinationFile, overwriteExisting)
		if err != nil {
			return err
		}
	}

	return nil
}

// Perform action when creating the shortcut
func (s *Shortcut) OnCreate() error {

	// Ensure to have a valid shortcut ID
	if s.ID == "" || s.ID == "0" {
		s.ID = GenerateID(s.Name, s.Executable)
	}

	err := s.ProcessImages(true)
	if err != nil {
		return err
	}

	return nil
}

// Perform action when updating the shortcut
func (s *Shortcut) OnUpdate() error {

	err := s.ProcessImages(true)
	if err != nil {
		return err
	}

	return nil
}

// Perform action when merging the shortcut
func (s *Shortcut) OnMerge() error {

	err := s.ProcessImages(true)
	if err != nil {
		return err
	}

	return nil
}

// Perform action when removing the shortcut
func (s *Shortcut) OnRemove() error {

	// Remove all images of the shortcut
	images := []string{
		s.IconPath,
		s.LogoPath,
		s.CoverPath,
		s.BannerPath,
		s.HeroPath,
	}

	for _, file := range images {
		if file == "" {
			continue
		}
		err := fs.RemoveFile(file)
		if err != nil {
			return err
		}
	}

	return nil
}
