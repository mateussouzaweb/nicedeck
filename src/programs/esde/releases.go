package esde

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mateussouzaweb/nicedeck/src/packaging"
)

// Get download URL from the latest release available
func GetDownloadURL(releaseType string) (string, error) {

	endpoint := "https://gitlab.com/es-de/emulationstation-de/-/raw/master/latest_release.json"
	res, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer func() {
		errors.Join(err, res.Body.Close())
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var releases struct {
		Stable struct {
			Version  string `json:"version"`
			Packages []struct {
				Name     string `json:"name"`
				Filename string `json:"filename"`
				URL      string `json:"url"`
				Md5      string `json:"md5"`
			} `json:"packages"`
		} `json:"stable"`
	}

	err = json.Unmarshal(body, &releases)
	if err != nil {
		return "", err
	}

	for _, release := range releases.Stable.Packages {
		if release.Name == releaseType {
			return release.URL, nil
		}
	}

	return "", fmt.Errorf("could not retrieve latest release")
}

// Return packaging source from release
func Release(releaseType string, format string) *packaging.Source {
	return &packaging.Source{
		Format: format,
		Resolver: func() (string, error) {
			return GetDownloadURL(releaseType)
		},
	}
}
