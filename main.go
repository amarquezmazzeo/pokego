package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	cliCommandMap := map[string]cliCommand{
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
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputWordList := cleanInput(scanner.Text())
		command := inputWordList[0]
		if c, ok := cliCommandMap[command]; ok {
			c.callback()
		} else {
			fmt.Printf("Invalid command %s\n", command)
		}
		//fmt.Printf("Your command was: %s\n", inputWL[0])
		fmt.Print("Pokedex > ")
	}
}

func cleanInput(text string) []string {
	textLower := strings.ToLower(text)
	cleanText := strings.Fields(textLower)
	return cleanText
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	helpText := `Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`
	fmt.Println(helpText)
	return nil
}
