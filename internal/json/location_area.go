package json

import "encoding/json"

type LocationAreaResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreaData(url string) (LocationAreaResult, error) {
	laResult := LocationAreaResult{}

	body, err := GetResponseFromURL(url, &CachedData)
	if err != nil {
		return LocationAreaResult{}, err
	}

	err = json.Unmarshal(body, &laResult)
	if err != nil {
		return LocationAreaResult{}, err
	}

	return laResult, nil
}

func GetLocationAreaList(l LocationAreaResult) []string {
	outputString := []string{}

	for _, item := range l.Results {
		outputString = append(outputString, item.Name)
	}

	return outputString
}
