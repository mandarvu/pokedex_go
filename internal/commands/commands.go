// Package commands
package commands

import (
	"fmt"
	"os"
)

const pokeAPIBaseURL = "https://pokeapi.co/api/v2/"

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
}

type Command struct {
	name        string
	description string
	Callback    func(*Config) error
}

func CommandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	return nil
}

func printList(list []string) {
	for _, item := range list {
		fmt.Println(item)
	}
}
