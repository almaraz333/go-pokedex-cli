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
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nPokedex > ")

		scanner.Scan()
		commandSplit := strings.Split(scanner.Text(), " ")

		command := commandSplit[0]
		var params string

		if len(commandSplit) > 1 {
			params = commandSplit[1]
		}

		if val, ok := commands[command]; ok {
			val.Callback(&config, &cache, params)
		}
	}
}
