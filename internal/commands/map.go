package commands

import ijson "github.com/mandarvu/pokedex_go/internal/json"

func CommandMap(c *Config, arguments []string) error {
	locAreaListEndpoint := pokeAPIBaseURL + locationAreaEndpoint

	if c.NextURL == "" {
		body, err := ijson.GetLocationAreaData(locAreaListEndpoint)
		if err != nil {
			return err
		}

		c.NextURL = body.Next
		c.PrevURL = body.Previous

		printList(ijson.GetLocationAreaList(body))

	} else {
		body, err := ijson.GetLocationAreaData(c.NextURL)
		if err != nil {
			return err
		}

		c.NextURL = body.Next
		c.PrevURL = body.Previous
		printList(ijson.GetLocationAreaList(body))
	}
	return nil
}

func CommandMapb(c *Config, arguments []string) error {
	locAreaListEndpoint := pokeAPIBaseURL + locationAreaEndpoint

	if c.PrevURL == "" {
		body, err := ijson.GetLocationAreaData(locAreaListEndpoint)
		if err != nil {
			return err
		}

		c.NextURL = body.Next
		c.PrevURL = body.Previous

		printList(ijson.GetLocationAreaList(body))

	} else {
		body, err := ijson.GetLocationAreaData(c.PrevURL)
		if err != nil {
			return err
		}

		c.NextURL = body.Next
		c.PrevURL = body.Previous
		printList(ijson.GetLocationAreaList(body))
	}
	return nil
}
