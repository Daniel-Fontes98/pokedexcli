package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *Config, args ...string) error{
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	resp, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNumber := rand.Intn(resp.BaseExperience)
	if randNumber > threshold {
		return fmt.Errorf("%s escaped", pokemonName)
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("You may now inpect it with the inspect command")
	cfg.caughtPokemon[pokemonName] = resp
	

	return nil
}