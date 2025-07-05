package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	config := &configCommand{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputWordList := cleanInput(scanner.Text())
		if len(inputWordList) == 0 {
			continue
		}

		command := inputWordList[0]

		c, ok := getCommands()[command]
		if !ok {
			fmt.Printf("Invalid command %s\n", command)
			continue
		}
		err := c.callback(config)

		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println(config.Next)
		//fmt.Printf("Your command was: %s\n", inputWL[0])
	}

}

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	cleanText := strings.Fields(textLower)
	return cleanText
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configCommand) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists 20 next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists prior 20 locations",
			callback:    commandMapb,
		},
	}
}
