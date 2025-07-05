package main

import (
	"errors"
	"fmt"

	"github.com/amarquezmazzeo/pokego/internal/pokeapi"
)

func commandMap(config *configCommand) error {

	locationResp, err := pokeapi.ListLocations(config.nextURL)
	if err != nil {
		return err
	}

	// fmt.Println(URL)

	config.nextURL = locationResp.Next
	config.previousURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(config *configCommand) error {
	if config.previousURL == nil {
		return errors.New("you're in the first page")
	}

	locationResp, err := pokeapi.ListLocations(config.previousURL)
	if err != nil {
		return err
	}

	// fmt.Println(URL)

	config.nextURL = locationResp.Next
	config.previousURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
