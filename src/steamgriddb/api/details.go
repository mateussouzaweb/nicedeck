package api

// DetailsResult struct
type DetailsResult struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Types    []string `json:"types"`
	Verified bool     `json:"verified"`
}

// DetailsByIdResult struct
type DetailsByIdResult struct {
	Success bool          `json:"success"`
	Errors  []string      `json:"errors"`
	Data    DetailsResult `json:"data"`
}

// DetailsBySteamAppIdResult struct
type DetailsBySteamAppIdResult struct {
	Success bool          `json:"success"`
	Errors  []string      `json:"errors"`
	Data    DetailsResult `json:"data"`
}

// Retrieve game or application details by ID
func GetDetailsById(id string) (*DetailsByIdResult, error) {
	endpoint := baseUrl + "/games/id/" + id
	result := DetailsByIdResult{}
	err := Request("GET", endpoint, &result)

	return &result, err
}

// Retrieve game or application details by Steam AppID
func GetDetailsBySteamAppId(appId string) (*DetailsBySteamAppIdResult, error) {

	endpoint := baseUrl + "/games/steam/" + appId
	result := DetailsBySteamAppIdResult{}
	err := Request("GET", endpoint, &result)

	return &result, err
}
