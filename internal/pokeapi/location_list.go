package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// TODO: Wrap errors with context and return an error for non-200 HTTP responses.
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

// TODO: Consider a context-aware request (timeouts/cancel)
// instead of http.Get (http.NewRequestWithContext)
func ExploreLocation(locationArea *string) (ExploreResponse, error) {
	if locationArea == nil || *locationArea == "" {
		return ExploreResponse{}, fmt.Errorf("locationArea must not be empty")
	}

	name := url.PathEscape(*locationArea)
	URL := baseURL + "/location-area/" + name

	res, err := http.Get(URL)
	if err != nil {
		return ExploreResponse{}, fmt.Errorf("GET %s: %w", URL, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ExploreResponse{}, fmt.Errorf("GET %s: unexpected status: %d", URL, res.StatusCode)
	}

	var exploreResp ExploreResponse

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&exploreResp); err != nil {
		return ExploreResponse{}, fmt.Errorf("decode explore response: %w", err)
	}
	return exploreResp, nil
}

func GetPokemonStats(pokemonName *string) (PokemonResponse, error) {
	if pokemonName == nil || *pokemonName == "" {
		return PokemonResponse{}, fmt.Errorf("pokemonName must not be empty")
	}

	name := url.PathEscape(*pokemonName)
	URL := baseURL + "/pokemon/" + name

	res, err := http.Get(URL)
	if err != nil {
		return PokemonResponse{}, fmt.Errorf("GET %s: %w", URL, err)
	}

	if res.StatusCode != http.StatusOK {
		return PokemonResponse{}, fmt.Errorf("GET %s: unexpected status: %d", URL, res.StatusCode)
	}

	var pokemonResponse PokemonResponse

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&pokemonResponse); err != nil {
		return PokemonResponse{}, fmt.Errorf("decode pokemon response: %w", err)
	}

	return pokemonResponse, nil
}
