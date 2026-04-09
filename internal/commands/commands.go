// Package commands
package commands

import (
	"fmt"
	"os"

	ijson "github.com/mandarvu/pokedex_go/internal/json"
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

func CommandMap(c *Config) error {
	locAreaListEndpoint := pokeAPIBaseURL + "location-area/"

	if c.NextURL == "" {
		body, err := ijson.GetLocationAreaData(locAreaListEndpoint)
		if err != nil {
			return err
		}

		c.NextURL = body.Next
		c.PrevURL = body.Previous

		printList(ijson.GetLocationAreaList(body))

	} else {
		body, err := ijson.GetLocationAreaData(c.NextURL)
		if err != nil {
			return err
		}

		c.NextURL = body.Next
		c.PrevURL = body.Previous
		printList(ijson.GetLocationAreaList(body))
	}
	return nil
}

func CommandMapb(c *Config) error {
	locAreaListEndpoint := pokeAPIBaseURL + "location-area/"

	if c.PrevURL == "" {
		body, err := ijson.GetLocationAreaData(locAreaListEndpoint)
		if err != nil {
			return err
		}

		c.NextURL = body.Next
		c.PrevURL = body.Previous

		printList(ijson.GetLocationAreaList(body))

	} else {
		body, err := ijson.GetLocationAreaData(c.PrevURL)
		if err != nil {
			return err
		}

		c.NextURL = body.Next
		c.PrevURL = body.Previous
		printList(ijson.GetLocationAreaList(body))
	}
	return nil
}

func printList(list []string) {
	for _, item := range list {
		fmt.Println(item)
	}
}
