package steamgriddb

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

const baseURL = "https://www.steamgriddb.com/api/v2"
const authorization = "Bearer 68e3c101bac17f05cafc31b437a012e5"

// Make request on SteamGridDB API
func Request(method string, endpoint string, result any) error {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	cli.Debug("Requesting %s\n", endpoint)

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		errors.Join(err, res.Body.Close())
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, result)

	return nil
}
