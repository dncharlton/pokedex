package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Print Prompt
		fmt.Print("\npoke_cli > ")
		scanner.Scan()

		// Grab and clean input, skip if empty
		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 { continue }

		// do some task
		// fmt.Println("echoing: ", text)
		command, ok := getCommands()[cleaned[0]]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		if !ok {
			fmt.Println(command.name, ": Command Not Found")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil { fmt.Println(err) }
	}
}

func cleanInput(str string) []string{
	return strings.Fields(strings.ToLower(str))
}

type cliCommand struct {
	name 				string
  description string
	callback 		func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "seeking help here",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "exit",
			callback: callbackQuit,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 locations",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locations",
			callback: callbackMapB,
		},
		"explore": {
			name: "explore {location_area}",
			description: "Lists the pokemon encounters in a location",
			callback: callbackExplore,
		},
		"catch": {
			name: "catch {pokemon_name}",
			description: "Catches a pokemon",
			callback: callbackCatch,
		},
		"inspect": {
			name: "inspect {pokemon_name}",
			description: "Inspects a pokemon",
			callback: callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists caught pokemon",
			callback: callbackPokedex,
		},
	}
}