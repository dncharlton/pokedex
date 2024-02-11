package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		return err
	}

	fmt.Printf("Location Areas:\n\n")
	for _, location := range resp.Locations {
		fmt.Printf("> %v\n", location.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Prev

	return nil
}

func callbackMapB(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)

	if err != nil {
		return err
	}
	
	fmt.Printf("Location Areas:\n\n")
	for _, location := range resp.Locations {
		fmt.Printf("> %v\n", location.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Prev

	return nil
}