package main

import "fmt"

func commandInspect(config *configCommand, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon_name>")
	}
	pokemonName := args[0]
	pokemon, ok := config.pokedex[pokemonName]
	if !ok {
		fmt.Printf("you have not caught %s!\n", pokemonName)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Weight: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	println("Stats:")
	for _, val := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", *val.Stat.Name, val.BaseStat)
	}
	println("Types:")
	for _, val := range pokemon.Types {
		fmt.Printf("  - %s\n", *val.Type.Name)
	}
	return nil
}
