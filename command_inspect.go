package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
  if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("you have not caught this pokemon")
	}

	fmt.Printf("Name:                  %s\n", pokemon.Name)
	fmt.Printf("ID:                    %v\n", pokemon.ID)
	fmt.Printf("Height:                %v\n", pokemon.Height)
	fmt.Printf("Weight:                %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %-20s%v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, types := range pokemon.Types {
		fmt.Printf(" - %s\n", types.Type.Name)
	}

	return nil
}
