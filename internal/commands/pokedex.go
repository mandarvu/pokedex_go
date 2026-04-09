package commands

import "fmt"

func CommandPokedex(c *Config, arguments []string) error {
	if len(Pokedex) == 0 {
		fmt.Println("You have not caught any pokemon yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, v := range Pokedex {
		fmt.Printf("  - %s\n", v.Name)
	}

	return nil
}
