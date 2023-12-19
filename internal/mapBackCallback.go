package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func MapBackCallback(config *Config, cache *Cache, params *PokeParams, pokedex map[string]PokemonRes) {
	if config.PrevUrl == nil {
		fmt.Println("Nothing to go back to")
		return
	}

	pokeRes := PokeMapRes{}

	cacheData, exists := cache.Get(*config.PrevUrl, cache)

	if exists {

		unmarshError := json.Unmarshal(cacheData, &pokeRes)

		if unmarshError == nil {
			for _, val := range pokeRes.Results {
				fmt.Println(val.Name)
			}
		}
	}

	res, err := http.Get(*config.PrevUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	unmarshError := json.Unmarshal(body, &pokeRes)

	if unmarshError != nil {
		fmt.Println(unmarshError)
	}

	for _, val := range pokeRes.Results {
		fmt.Println(val.Name)
	}

	config.NextUrl = pokeRes.Next
	config.PrevUrl = pokeRes.Previous
}
