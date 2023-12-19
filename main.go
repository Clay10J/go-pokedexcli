package main

import (
	"time"

	"github.com/c10j/go-pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
	pokedex              map[string]pokeapi.PokemonResponse
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokedex:       make(map[string]pokeapi.PokemonResponse),
	}

	startRepl(&cfg)
}
