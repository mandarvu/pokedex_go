package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type config struct {
	NextURL string `json:"next"`
	PrevURL string `json:"previous"`
}

var supportedCommands map[string]command = map[string]command{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Returns a location list",
		callback:    commandMap,
	},
}

type command struct {
	name        string
	description string
	callback    func(*config) error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	return nil
}

func commandMap(c *config) error {
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
