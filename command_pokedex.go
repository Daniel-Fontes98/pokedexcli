package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *Config, args ...string) error{
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you haven't caught any pokemon")
	}

	fmt.Println("Your pokedex:")
	for _ ,pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}