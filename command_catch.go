package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided. please enter a pokemon along with the catch command")
	}
	pokemonName := args[0]

	resp, err := cfg.pokeapiClient.FetchPokemon(pokemonName)
	if err != nil {
		return err
	}

	pokemonName = resp.Name

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	isCaught := attemptCatch(resp.BaseExperience)

	if isCaught {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = resp
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

// Highest base experience is 608 (Blissey) based on https://www.thegamer.com/pokemon-exp-when-defeated-experience-xp-base-yield-ranked/#eternamax-eternatus
// Therefore, using 700 for catch probability calculation
const maxBaseExp = 700

func attemptCatch(baseExperience int) bool {
	isCaught := false

	catchProbability := int(100 * (1 - (float64(baseExperience) / float64(maxBaseExp))))
	catchVal := rand.Intn(100)
	if catchVal <= catchProbability-1 {
		isCaught = true
	}

	return isCaught
}
