package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	newScan := bufio.NewScanner(os.Stdin)

	for {
		shellPrompt("Pokedex>")
		str, err := inputReader(*newScan)
		if err != nil {
			fmt.Printf("error reading the input: %v", err)
		}
		fmt.Println(str)
	}
}

func inputReader(scanner bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		return "", fmt.Errorf("could not read user input")
	}
	return scanner.Text(), nil
}

func shellPrompt(PS1 string) {
	fmt.Fprint(os.Stdout, PS1+" ")
}
