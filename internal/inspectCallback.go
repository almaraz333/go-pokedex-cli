package internal

import (
	"fmt"
)

func InspectCallback(config *Config, cache *Cache, params *PokeParams, pokedex map[string]PokemonRes) {
	if val, exists := pokedex[params.PokemonToCatch]; !exists {
		fmt.Println("This pokemon is not in your pokedex...")
		return
	} else {
		fmt.Printf("Height: %v\n", val.Height)
		fmt.Printf("Weight: %v\n", val.Weight)
		fmt.Println("Stats:")
		for _, elem := range val.Stats {
			fmt.Printf(" -%v: %v\n", elem.Stat.Name, elem.BaseStat)
		}
		fmt.Println("Types:")
		for _, elem := range val.Types {
			fmt.Printf(" - %v", elem.Type.Name)
		}

	}

	return
}
