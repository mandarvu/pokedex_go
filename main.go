package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	newScan := bufio.NewScanner(os.Stdin)

	for {
		shellPrompt("Pokedex >")
		str, err := inputReader(*newScan)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		cleanInput := cleanInput(str)

		fmt.Printf("Your command was: %s\n", cleanInput[0])
	}
}

	}
}

func shellPrompt(PS1 string) {
	fmt.Fprint(os.Stdout, PS1+" ")
}
