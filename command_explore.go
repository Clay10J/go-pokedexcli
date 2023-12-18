package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if cfg.name == nil {
		return errors.New("no location name provided. please enter a location along with the explore command")
	}

	resp, err := cfg.pokeapiClient.ListPokemonInLocationArea(cfg.name)
	if err != nil {
		return err
	}

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
