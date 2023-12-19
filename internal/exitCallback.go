package internal

import (
	"os"
)

func ExitCallback(config *Config, cache *Cache, params *PokeParams, pokedex map[string]PokemonRes) {
	os.Exit(1)
}
