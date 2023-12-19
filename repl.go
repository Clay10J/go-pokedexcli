package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		txt := scanner.Text()

		cleaned := cleanInput(txt)
		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		availableCommands := getCommands()

		cmd, ok := availableCommands[command]
		if !ok {
			fmt.Println("That command is not available. Use 'help' to see available commands.")
			fmt.Println("")
			continue
		}

		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(s string) []string {
	lowered := strings.ToLower(s)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Display the pokemon in a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    commandCatch,
		},
	}
}
