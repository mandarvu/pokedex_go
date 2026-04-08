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
			commandExit(&conf)
		}

		str := newScan.Text()

		if len(str) == 0 {
			continue
		}

		cleanedInput := cleanInput(str)
		command := cleanedInput[0]

		if comm, ok := supportedCommands[command]; ok {
			comm.callback(&conf)
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

func helpText(supportedCommands map[string]command) string {
	hstr := []byte{}

	for k, v := range supportedCommands {
		s := fmt.Sprintf("%s: %s\n", k, v.description)
		hstr = fmt.Append(hstr, s)
	}

	return string(hstr)
}
