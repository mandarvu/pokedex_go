package commands

import (
	"fmt"

	ijson "github.com/mandarvu/pokedex_go/internal/json"
)

func CommandInspect(c *Config, arguments []string) error {
	item, ok := Pokedex[arguments[0]]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	} else {
		fmt.Printf("Name: %s\n", item.Name)
		fmt.Printf("Height: %d\n", item.Height)
		fmt.Printf("Weight: %d\n", item.Weight)
		fmt.Println("Stats:")

		stats := getStats(item)

		fmt.Printf("  -hp: %d\n", stats["hp"])
		fmt.Printf("  -attack: %d\n", stats["attack"])
		fmt.Printf("  -defense: %d\n", stats["defense"])
		fmt.Printf("  -special-attack: %d\n", stats["special-attack"])
		fmt.Printf("  -special-defense: %d\n", stats["special-defense"])
		fmt.Printf("  -speed: %d\n", stats["speed"])

		types := getTypes(item)

		fmt.Printf("Type:\n")
		for _, j := range types {
			fmt.Printf("  - %s\n", j)
		}
	}

	return nil
}

func getStats(p ijson.Pokemon) map[string]int {
	m := map[string]int{}

	for _, j := range p.Stats {
		m[j.Stat.Name] = j.BaseStat
	}

	return m
}

func getTypes(p ijson.Pokemon) []string {
	m := []string{}

	for _, j := range p.Types {
		m = append(m, j.Type.Name)
	}

	return m
}
