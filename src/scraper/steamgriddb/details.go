package steamgriddb

// DetailsResult struct
type DetailsResult struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Types       []string `json:"types"`
	Verified    bool     `json:"verified"`
	ReleaseDate int64    `json:"release_date"`
}

// DetailsByIDResult struct
type DetailsByIDResult struct {
	Success bool          `json:"success"`
	Errors  []string      `json:"errors"`
	Data    DetailsResult `json:"data"`
}

// DetailsBySteamAppIDResult struct
type DetailsBySteamAppIDResult struct {
	Success bool          `json:"success"`
	Errors  []string      `json:"errors"`
	Data    DetailsResult `json:"data"`
}

// Retrieve game or application details by ID
func GetDetailsByID(id string) (*DetailsByIDResult, error) {
	endpoint := baseURL + "/games/id/" + id
	result := DetailsByIDResult{}
	err := Request("GET", endpoint, &result)

	return &result, err
}

// Retrieve game or application details by Steam AppID
func GetDetailsBySteamAppID(appID string) (*DetailsBySteamAppIDResult, error) {

	endpoint := baseURL + "/games/steam/" + appID
	result := DetailsBySteamAppIDResult{}
	err := Request("GET", endpoint, &result)

	return &result, err
}
