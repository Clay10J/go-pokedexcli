package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pkmn := range cfg.pokedex {
		fmt.Printf("\t- %s\n", pkmn.Name)
	}

	return nil
}
