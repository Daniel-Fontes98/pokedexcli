package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type CliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

type Config struct {
	Next *string
	Previous *string
}

func newConfig() Config {
	return Config {
		Next: nil,
		Previous: nil,
	}
}

func getCommands() map[string]CliCommand{
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name: "map",
			description: `Displays the name of 20 location areas in the Pokemon world. 
			Each subsequent call to map displays the next 20 location`,
			callback: callbackMap,
		},
	}
}

func startRepl() {
	config := newConfig()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback(&config)
	}
}

func cleanInput(s string) []string {
	lowerCaseInput := strings.ToLower(s)
	words := strings.Fields(lowerCaseInput)
	return words
}