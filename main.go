package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	newScan := bufio.NewScanner(os.Stdin)

	for {
		prompt("Pokedex >")

		if !newScan.Scan() {
			commandExit()
		}

		str := newScan.Text()

		cleanedInput := cleanInput(str)
		command := cleanedInput[0]

		if comm, ok := supportedCommands[command]; ok {
			comm.callback()
			if command == "help" {
				fmt.Printf("\n%s", helpText(supportedCommands))
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func prompt(PS1 string) {
	fmt.Fprint(os.Stdout, PS1+" ")
}
