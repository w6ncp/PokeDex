package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("no pokemon to list")
	}

	fmt.Println("Your PokeDex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println("  -", pokemon.Name)
	}
	return nil
}
