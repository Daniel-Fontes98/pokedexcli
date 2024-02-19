package main

import (
	"time"

	"github.com/Daniel-Fontes98/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.Client
	NextLocationAreaUrl *string
	PreviousLocationAreaUrl *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	config := Config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&config)

}

