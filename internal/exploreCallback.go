package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ExploreCallbak(config *Config, cache *Cache, params *PokeParams, pokedex map[string]PokemonRes) {
	if params.Location == "" {
		fmt.Println("Please provide an area to explore...")
		return
	}

	exploreRes := ExploreRes{}

	fmt.Printf("Exploring %v ...\n", params.Location)

	cachedData, exists := cache.Get(params.Location, cache)

	if exists {
		unmarshError := json.Unmarshal(cachedData, &exploreRes)

		if unmarshError != nil {
			log.Fatal(unmarshError)
		}

		fmt.Println("Found Pokemon:")
		for _, val := range exploreRes.PokemonEncounters {
			fmt.Println(val.Pokemon.Name)
		}

		return
	}

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + params.Location)

	if err != nil {
		fmt.Println("Something went wrong...")
		return
	}

	body, readError := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if readError != nil {
		log.Fatal(err)
	}

	cache.Add(params.Location, body, cache)

	unmarshError := json.Unmarshal(body, &exploreRes)

	if unmarshError != nil {
		log.Fatal(unmarshError)
	}

	fmt.Println("Found Pokemon:")
	for _, val := range exploreRes.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}

}
