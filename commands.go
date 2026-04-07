package main

import (
	"fmt"
	"os"
)

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
}

type command struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	return nil
}

func helpText(supportedCommands map[string]command) string {
	hstr := []byte{}

	for k, v := range supportedCommands {
		s := fmt.Sprintf("%s: %s\n", k, v.description)
		hstr = fmt.Append(hstr, s)
	}

	return string(hstr)
}
