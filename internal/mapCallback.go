package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func MapCallback(config *Config, cache *Cache, params string) {

	cacheData, exists := cache.Get(config.NextUrl, cache)

	pokeRes := PokeMapRes{}

	if exists {
		unmarshError := json.Unmarshal(cacheData, &pokeRes)
		if unmarshError == nil {
			for _, val := range pokeRes.Results {
				fmt.Println(val.Name)
			}
		}

		config.NextUrl = pokeRes.Next
		config.PrevUrl = pokeRes.Previous

		return
	}

	res, err := http.Get(config.NextUrl)

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

	cache.Add(config.NextUrl, body, cache)

	for _, val := range pokeRes.Results {
		fmt.Println(val.Name)
	}

	config.NextUrl = pokeRes.Next
	config.PrevUrl = pokeRes.Previous
}
