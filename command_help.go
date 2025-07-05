package main

import "fmt"

func commandHelp(config *configCommand) error {
	helpText := `Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Lists next 20 locations
mapb: Lists 20 prior locations`
	fmt.Println(helpText)
	return nil
}
