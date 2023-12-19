package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/almaraz333/go-pokedex-cli/internal"
)

func startRepl() {
	config := internal.Config{NextUrl: "https://pokeapi.co/api/v2/location-area"}

	cache := internal.NewCache(10000)

	commands := map[string]internal.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    internal.HelpCallback,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    internal.ExitCallback,
		},
		"map": {
			Name:        "map",
			Description: "displays the names of 20 location areas in the Pokemon world",
			Callback:    internal.MapCallback,
		},
		"mapb": {
			Name:        "mapb",
			Description: "displays the previous 20 locations",
			Callback:    internal.MapBackCallback,
		},
		"explore": {
			Name:        "Explore",
			Description: "Explore the current area",
			Callback:    internal.ExploreCallbak,
		},
		"catch": {
			Name:        "Catch",
			Description: "Catch a pokemon in the current area",
			Callback:    internal.CatchCallback,
		},
		"inspect": {
			Name:        "Inspect",
			Description: "Get info about a pokemon in your pokedex",
			Callback:    internal.InspectCallback,
		},
		"pokedex": {
			Name:        "Pokedex",
			Description: "List the pokemon in your pokedex",
			Callback:    internal.PokedexCallback,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	params := &internal.PokeParams{}

	pokedex := map[string]internal.PokemonRes{}

	for {
		fmt.Print("\nPokedex > ")

		scanner.Scan()
		commandSplit := strings.Split(scanner.Text(), " ")

		command := commandSplit[0]

		if command == "explore" {
			params.Location = commandSplit[1]
		} else if command == "catch" {
			params.PokemonToCatch = commandSplit[1]
		}

		if val, ok := commands[command]; ok {
			val.Callback(&config, &cache, params, pokedex)
		}
	}
}
