package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
  if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]
	resp, err := cfg.pokeapiClient.ListLocationEncounters(locationAreaName)

	if err != nil {
		return err
	}

	fmt.Printf("Location Encounters:\n\n")
	for _, pk := range resp.PokemonEncounters {
		fmt.Println(pk.Pokemon.Name)
	}

	return nil
}
