package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Get asset direct download URL from the latest release available
func GetAssetURL(repository string, search string) (string, error) {

	repository = strings.ReplaceAll(repository, "https://github.com/", "")
	repository = strings.Trim(repository, "/")

	// Request latest releases
	endpoint := fmt.Sprintf("https://api.github.com/repos/%s/releases", repository)
	res, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Decode response into struct
	var releases []struct {
		Name       string `json:"name"`
		PreRelease bool   `json:"prerelease"`
		Assets     []struct {
			Name        string `json:"name"`
			ContentType string `json:"content_type"`
			URL         string `json:"url"`
			DownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}

	err = json.Unmarshal(body, &releases)
	if err != nil {
		return "", err
	}

	// Create rule from search
	search = strings.ReplaceAll(search, "*", "(.+)")
	searchRegex := regexp.MustCompile(search)

	// Check for matching asset
	for _, release := range releases {
		for _, asset := range release.Assets {
			if searchRegex.MatchString(asset.Name) {
				return asset.DownloadURL, nil
			}
		}
	}

	return "", fmt.Errorf("could not retrieve latest release asset")
}

// Find expected format based on search term
func FindFormat(search string) string {

	if strings.HasSuffix(search, ".zip") {
		return "zip"
	} else if strings.HasSuffix(search, ".tar.gz") {
		return "tar.gz"
	} else if strings.HasSuffix(search, ".tar.xz") {
		return "tar.xz"
	} else if strings.HasSuffix(search, ".7z") {
		return "7z"
	} else if strings.HasSuffix(search, ".dmg") {
		return "dmg"
	}

	return "file"
}

// Return packaging source from release
func Release(repository string, search string) *packaging.Source {
	return &packaging.Source{
		Format: FindFormat(search),
		Resolver: func() (string, error) {
			return GetAssetURL(repository, search)
		},
	}
}
