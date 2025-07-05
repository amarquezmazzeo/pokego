package pokeapi

import (
	"encoding/json"
	"net/http"
)

func ListLocations(pageURL *string) (locationResponse, error) {
	URL := baseURL + "/location-area"
	// fmt.Println(config)
	if pageURL != nil {
		URL = *pageURL
	}

	res, err := http.Get(URL)
	if err != nil {
		return locationResponse{}, err
	}
	defer res.Body.Close()

	var locationResp locationResponse

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&locationResp); err != nil {
		return locationResponse{}, err
	}

	return locationResp, nil
}
