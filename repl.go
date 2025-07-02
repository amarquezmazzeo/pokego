package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		inputWordList := cleanInput(scanner.Text())
		if len(inputWordList) == 0 {
			continue
		}

		command := inputWordList[0]
		if c, ok := getCommands()[command]; ok {
			err := c.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Invalid command %s\n", command)
		}
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
	callback    func() error
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
	}
}
