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
inspect: Retrieve stats of given pokemon, if in pokedex
`
	fmt.Println(helpText)
	return nil
}
