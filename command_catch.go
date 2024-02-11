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
	pokemonName := args[0]
	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	var threshold = int(float64(resp.BaseExperience)*0.4)
	randCatch := rand.Intn(resp.BaseExperience)

	if randCatch > threshold {
		return fmt.Errorf("failed to catch %s", resp.Name)
	}

	cfg.caughtPokemon[resp.Name] = resp
	fmt.Printf("%s was caught\n", resp.Name)

	return nil
}
