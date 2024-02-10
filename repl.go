package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Print Prompt
		fmt.Print("\npoke_cli > ")
		scanner.Scan()

		// Grab and clean input, skip if empty
		text := cleanInput(scanner.Text())
		if len(text) == 0 { continue }

		// do some task
		// fmt.Println("echoing: ", text)
		commandName := text[0]
		command, ok := getCommands()[commandName]

		if !ok {
			fmt.Println(command.name, ": Command Not Found")
			continue
		}

		command.callback()
	}
}

func cleanInput(str string) []string{
	return strings.Fields(strings.ToLower(str))
}

type cliCommand struct {
	name 				string
  description string
	callback 		func()
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
	}
}