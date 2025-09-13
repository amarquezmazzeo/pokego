package pokeapi

import (
	"encoding/json"
	"net/http"
)

func ListLocations(pageURL *string) (LocationResponse, error) {
	URL := baseURL + "/location-area"
	// fmt.Println(config)
	if pageURL != nil {
		URL = *pageURL
	}

	res, err := http.Get(URL)
	if err != nil {
		return LocationResponse{}, err
	}
	defer res.Body.Close()

	var locationResp LocationResponse

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&locationResp); err != nil {
		return LocationResponse{}, err
	}

	return locationResp, nil
}
