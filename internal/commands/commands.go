// Package commands
package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		description: "Returns a location list",
		Callback:    CommandMap,
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
		body, err := getResponseFromURL(locAreaListEndpoint)
		if err != nil {
			return err
		}

		fmt.Println(body)
		json.Unmarshal([]byte(body), c)
	} else {
		body, err := getResponseFromURL(c.NextURL)
		if err != nil {
			return err
		}

		fmt.Println(body)
		json.Unmarshal([]byte(body), c)
	}
	return nil
}

func getResponseFromURL(u string) (string, error) {
	res, err := http.Get(u)
	if err != nil {
		return "", fmt.Errorf("could Not GET response from URL %s: %v", u, err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return "", fmt.Errorf("response failed with status code: %d and\nbody: %sn", res.StatusCode, body)
	}

	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return string(body), nil
}
