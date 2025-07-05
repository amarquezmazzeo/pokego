package main

import (
	"fmt"
	"os"
)

func commandExit(config ConfigCommand) (ConfigCommand, error) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return config, nil
}
