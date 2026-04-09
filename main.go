package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mandarvu/pokedex_go/internal/commands"
	"github.com/mandarvu/pokedex_go/internal/input"
)

func main() {
	newScan := bufio.NewScanner(os.Stdin)

	conf := commands.Config{
		NextURL: "",
		PrevURL: "",
	}

	for {
		prompt("Pokedex >")

		if !newScan.Scan() {
			commands.CommandExit(&conf)
		}

		str := newScan.Text()

		if len(str) == 0 {
			continue
		}

		cleanedInput := input.CleanInput(str)
		command := cleanedInput[0]

		if comm, ok := commands.SupportedCommands[command]; ok {
			comm.Callback(&conf)
			if command == "help" {
				fmt.Printf("\n%s", commands.HelpText(commands.SupportedCommands))
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func prompt(PS1 string) {
	fmt.Fprint(os.Stdout, PS1+" ")
}
