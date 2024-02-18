package main

import (
	"fmt"
)

func callbackMap(cfg *Config, args ... string) error{
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.NextLocationAreaUrl)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}

	cfg.NextLocationAreaUrl = resp.Next
	cfg.PreviousLocationAreaUrl = resp.Previous

	return nil
}