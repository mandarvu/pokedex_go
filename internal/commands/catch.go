package commands

import (
	"fmt"
	"math"
	"math/rand"

	ijson "github.com/mandarvu/pokedex_go/internal/json"
)

func CommandCatch(c *Config, arguments []string) error {
	pokemonToCatch := pokeAPIBaseURL + pokemonEndpoint + arguments[0]

	body, err := ijson.GetPokemonData(pokemonToCatch)
	if err != nil {
		return err
	}

	r := rand.NormFloat64()

	fmt.Printf("Throwing a Pokeball at %s...\n", arguments[0])

	if math.Abs(r*100+(float64(body.BaseExperience)/69)) > catchThreshold {
		fmt.Println(arguments[0], " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		Pokedex[body.Name] = body
	} else {
		fmt.Println(arguments[0], " escaped!")
	}
	return nil
}
