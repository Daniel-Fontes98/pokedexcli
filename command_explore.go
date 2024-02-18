package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *Config, args ...string) error{
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]

	fmt.Printf("Exploring %s\n", locationAreaName)

	resp, err := cfg.pokeapiClient.ListPokemonInArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}