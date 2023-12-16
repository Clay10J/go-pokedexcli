package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	for _, info := range getCommands() {
		str := fmt.Sprintf("%s: %s", info.name, info.description)
		fmt.Println(str)
	}

	fmt.Println()
	return nil
}
