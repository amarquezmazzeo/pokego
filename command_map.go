package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/amarquezmazzeo/pokego/internal/pokeapi"
)

func commandMap(config *configCommand) error {
	if config.nextURL == nil {
		defaultURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
		config.nextURL = &defaultURL
	}

	locationResp, inCache := config.cache.Get(*config.nextURL)
	if !inCache {
		// fmt.Println("Cache Not Hit :(")
		apiResp, err := pokeapi.ListLocations(config.nextURL)
		if err != nil {
			return err
		}

		locationResp, err = json.Marshal(apiResp)
		if err != nil {
			return err
		}
		config.cache.Add(*config.nextURL, locationResp)
	} // else {
	// 	fmt.Println("Cache Hit!")
	// }

	var jsonResponse pokeapi.LocationResponse

	err := json.Unmarshal(locationResp, &jsonResponse)
	if err != nil {
		return err
	}

	// fmt.Println(URL)

	config.nextURL = jsonResponse.Next
	config.previousURL = jsonResponse.Previous

	for _, location := range jsonResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(config *configCommand) error {
	if config.previousURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, inCache := config.cache.Get(*config.previousURL)
	if !inCache {
		// fmt.Println("Cache Not Hit :(")
		apiResp, err := pokeapi.ListLocations(config.previousURL)
		if err != nil {
			return err
		}

		locationResp, err = json.Marshal(apiResp)
		if err != nil {
			return err
		}
		config.cache.Add(*config.previousURL, locationResp)
	} // else {
	// 	fmt.Println("Cache Hit!")
	// }

	var jsonResponse pokeapi.LocationResponse

	err := json.Unmarshal(locationResp, &jsonResponse)
	if err != nil {
		return err
	}

	// fmt.Println(URL)

	config.nextURL = jsonResponse.Next
	config.previousURL = jsonResponse.Previous

	for _, location := range jsonResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
