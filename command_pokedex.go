package main

import "fmt"

func commandPokedex(config *configCommand, args []string) error {
	if len(config.pokedex) == 0 {
		println("Pokedex is empty!")
		return nil
	}

	for pokemon := range config.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}
