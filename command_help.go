package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	for _, info := range getCommands() {
		str := fmt.Sprintf("%s: %s", info.name, info.description)
		fmt.Println(str)
	}

	fmt.Println()
	return nil
}
