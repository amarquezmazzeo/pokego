package main

import "fmt"

func commandHelp() error {
	helpText := `Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`
	fmt.Println(helpText)
	return nil
}
