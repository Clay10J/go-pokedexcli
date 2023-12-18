package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location name provided. please enter a location along with the explore command")
	}
	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.ListPokemonInLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
