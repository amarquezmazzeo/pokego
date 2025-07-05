package main

import "fmt"

func commandHelp(config ConfigCommand) (ConfigCommand, error) {
	helpText := `Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Lists 20 locations`
	fmt.Println(helpText)
	return config, nil
}
