package main

import (
	"fmt"
	"log"

	"github.com/Daniel-Fontes98/pokedexcli/internal/pokeapi"
)

func callbackMap(c *Config) error{
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}