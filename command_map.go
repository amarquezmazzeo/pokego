package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func commandMap(config ConfigCommand) (ConfigCommand, error) {
	URL := config.Next
	// fmt.Println(config)

	if config.isBack {
		if config.Previous == "" {
			err := errors.New("you're on the first page")
			return config, err
		}
		URL = config.Previous
	}

	if len(URL) == 0 {
		URL = "https://pokeapi.co/api/v2/location-area"
	}

	// fmt.Println(URL)

	res, err := http.Get(URL)
	if err != nil {
		return config, err
	}

	var locationResponse LocationResponse

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&locationResponse); err != nil {
		return config, err
	}
	for _, element := range locationResponse.Results {
		fmt.Println(element.Name)
	}

	config.Next = locationResponse.Next
	config.Previous = locationResponse.Previous
	config.isBack = false

	return config, nil
}

type LocationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
