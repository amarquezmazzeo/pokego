package main

import (
	"encoding/json"
	"fmt"

	"github.com/amarquezmazzeo/pokego/internal/pokeapi"
)

func commandExplore(config *configCommand, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <area_name>")
	}
	area := args[0]

	exploreResp, inCache := config.cache.Get(area)
	if !inCache {
		// fmt.Println("Cache Not Hit :(")
		apiResp, err := pokeapi.ExploreLocation(&area)
		if err != nil {
			return err
		}
		// time.Sleep(500 * time.Millisecond)

		exploreResp, err = json.Marshal(apiResp)
		if err != nil {
			return err
		}
		config.cache.Add(area, exploreResp)
	} // else {
	// 	fmt.Println("Cache Hit!")
	// }

	var jsonResponse pokeapi.ExploreResponse

	err := json.Unmarshal(exploreResp, &jsonResponse)
	if err != nil {
		return err
	}

	// fmt.Println(URL)

	for _, pokemonEncounters := range jsonResponse.PokemonEncounters {
		fmt.Println(pokemonEncounters.Pokemon.Name)
	}

	return nil
}
