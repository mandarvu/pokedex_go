package commands

import (
	"fmt"

	ijson "github.com/mandarvu/pokedex_go/internal/json"
)

func CommandExplore(c *Config, arguments []string) error {
	areaToExplore := arguments[0]
	areaURL := pokeAPIBaseURL + locationAreaEndpoint + areaToExplore + "/"

	body, err := ijson.GetPokemonInArea(areaURL)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", arguments[0])
	fmt.Printf("Found Pokemon:\n")

	printList(ijson.GetAreaPokemonList(body))

	return nil
}
