package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func CatchCallback(config *Config, cache *Cache, params *PokeParams, pokedex map[string]PokemonRes) {
	if params.Location == "" || params.PokemonToCatch == "" {
		fmt.Println("Cannot catch without a valid location or pokemon to catch")
		return
	}

	cachedLocationValue, locationCached := cache.Get(params.Location, cache)

	if !locationCached {
		fmt.Println("No location to explore")
		return
	}

	exploreRes := ExploreRes{}

	unmarshError := json.Unmarshal(cachedLocationValue, &exploreRes)

	if unmarshError != nil {
		log.Fatal(unmarshError)
	}

	found := false

	for _, val := range exploreRes.PokemonEncounters {
		if val.Pokemon.Name == params.PokemonToCatch {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("This pokemon was not found...")
		return
	}

	fmt.Printf("Throwing a pokeball at %v\n", params.PokemonToCatch)

	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + params.PokemonToCatch)

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

	pokemonRes := PokemonRes{}

	unmarshErrorPokemon := json.Unmarshal(body, &pokemonRes)

	if unmarshErrorPokemon != nil {
		log.Fatal(unmarshError)
	}

	randomNumber := rand.Intn(400)
	if randomNumber > pokemonRes.BaseExperience {
		fmt.Printf("%v was caught!\n", params.PokemonToCatch)
		pokedex[params.PokemonToCatch] = pokemonRes
	} else {
		fmt.Printf("%v escaped!\n", params.PokemonToCatch)
	}

	return
}
