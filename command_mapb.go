package main

import (
	"fmt"
)

func callbackMapb(cfg *Config, args ...string) error{
	if cfg.PreviousLocationAreaUrl == nil {
		return fmt.Errorf("there are no previous pages")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.PreviousLocationAreaUrl)
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