package main

import (
	"errors"
	"fmt"

	"github.com/c10j/go-pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided. please enter a pokemon along with the inspect command")
	}
	pokemonName := args[0]

	pokemon, exists := cfg.pokedex[pokemonName]
	if !exists {
		fmt.Printf("you have not caught %s yet\n", pokemonName)
		return nil
	}

	printPokemonInfo(pokemon)

	return nil
}

func printPokemonInfo(pokemon pokeapi.PokemonResponse) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("\t- %s\n", typ.Type.Name)
	}
}
