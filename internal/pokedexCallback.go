package internal

import (
	"fmt"
)

func PokedexCallback(config *Config, cache *Cache, params *PokeParams, pokedex map[string]PokemonRes) {
	for _, pokemon := range pokedex {
		fmt.Println(" - ", pokemon.Name)
	}

	return
}
