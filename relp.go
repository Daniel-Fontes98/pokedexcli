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
	callback    func(*Config, ...string) error
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
			description: `Displays the name of the next 20 location areas in the Pokemon world.`,
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: `Displays the name of the previous 20 location areas in the Pokemon world.`,
			callback: callbackMapb,
		},
		"explore": {
			name: "explore <location_name>",
			description: `Displays the name of all pokemon in a given area`,
			callback: callbackExplore,
		},
	}
}

func startRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(config, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(s string) []string {
	lowerCaseInput := strings.ToLower(s)
	words := strings.Fields(lowerCaseInput)
	return words
}