package main

import (
	"time"

	"github.com/c10j/go-pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(&cfg)
}
