package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
  if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	locationAreaName := args[0]
	resp, err := cfg.pokeapiClient.GetPokemon(locationAreaName)

	if err != nil {
		return err
	}

	const threshold = 50
	randCatch := rand.Intn(resp.BaseExperience)

	if randCatch > threshold {
		return fmt.Errorf("failed to catch %s", resp.Name)
	}

	cfg.caughtPokemon[resp.Name] = resp
	// fmt.Printf("%s was caught\n", resp.Name)

	return nil
}
