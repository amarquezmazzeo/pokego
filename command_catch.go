package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"

	"github.com/amarquezmazzeo/pokego/internal/pokeapi"
)

func commandCatch(config *configCommand, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon_name>")
	}
	pokemonName := &args[0]
	if _, ok := config.pokedex[*pokemonName]; ok {
		fmt.Printf("%s already caught!\n", *pokemonName)
		return nil
	}

	name := url.PathEscape(*pokemonName)
	URL := "https://pokeapi.co/api/v2/location-area" + name
	pointerURL := &URL

	pokemonResp, inCache := config.cache.Get(*pointerURL)
	if !inCache {
		apiResp, err := pokeapi.GetPokemonStats(pokemonName)
		if err != nil {
			return err
		}

		pokemonResp, err = json.Marshal(apiResp)
		if err != nil {
			return err
		}
		config.cache.Add(*pointerURL, pokemonResp)
	}

	var jsonResponse pokeapi.PokemonResponse

	err := json.Unmarshal(pokemonResp, &jsonResponse)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *pokemonName)
	chance := rand.Intn(int(jsonResponse.BaseExperience))
	if chance < 50 {
		config.pokedex[*pokemonName] = jsonResponse
		fmt.Printf("%s was caught!\n", *pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", *pokemonName)
	}

	return nil
}
