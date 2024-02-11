package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	for _, cmd := range getCommands() {
		fmt.Println(cmd.name, ": ", cmd.description)
	}
	return nil
}