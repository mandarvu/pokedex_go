// Package commands
package commands

import (
	"fmt"
	"os"
)

const (
	pokeAPIBaseURL       = "https://pokeapi.co/api/v2/"
	locationAreaEndpoint = "location-area/"
)

type Config struct {
	NextURL string `json:"next"`
	PrevURL string `json:"previous"`
}

var SupportedCommands map[string]Command = map[string]Command{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		Callback:    CommandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		Callback:    CommandHelp,
	},
	"map": {
		name:        "map",
		description: "Returns a next location list",
		Callback:    CommandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Returns a previous location list",
		Callback:    CommandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Returns a list of pokemon in the given area",
		Callback:    CommandExplore,
	},
}

type Command struct {
	name        string
	description string
	Callback    func(*Config, []string) error
}

func CommandExit(c *Config, arguments []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(c *Config, arguments []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	return nil
}

func printList(list []string) {
	for _, item := range list {
		fmt.Println(item)
	}
}
