package main

import (
	"time"

	"github.com/Daniel-Fontes98/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.Client
	NextLocationAreaUrl *string
	PreviousLocationAreaUrl *string
}

func main() {
	config := Config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&config)

}

