package main

import "fmt"

func commandHelp(config *configCommand, args []string) error {
	helpText := `Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Lists next 20 locations
mapb: Lists 20 prior locations
explore: Lists pokemon in given location
catch: Attempt to catch given pokemon
`
	fmt.Println(helpText)
	return nil
}
