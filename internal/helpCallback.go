package internal

import (
	"fmt"
)

func HelpCallback(config *Config, cache *Cache, params *PokeParams, pokedex map[string]PokemonRes) {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
}
