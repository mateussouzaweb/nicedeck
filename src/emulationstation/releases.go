package emulationstation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Get latest release package available
func GetLatestRelease() (string, error) {

	endpoint := "https://gitlab.com/es-de/emulationstation-de/-/raw/master/latest_release.json"
	res, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
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
		if release.Name == "LinuxAppImage" {
			return release.URL, nil
		}
	}

	return "", fmt.Errorf("could not retrieve latest release")
}
