package main

import "fmt"

func callbackHelp() {
	for _, cmd := range getCommands() {
		fmt.Println(cmd.name, ": ", cmd.description)
	}
}