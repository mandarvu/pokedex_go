// Package json (used as ijson to signify internal nature of this package)
package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

	body, err := GetResponseFromURL(url)
	if err != nil {
		return LocationAreaResult{}, err
	}

	err = json.Unmarshal([]byte(body), &laResult)
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

func GetResponseFromURL(u string) (string, error) {
	res, err := http.Get(u)
	if err != nil {
		return "", fmt.Errorf("could Not GET response from URL %s: %v", u, err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return "", fmt.Errorf("response failed with status code: %d and\nbody: %sn", res.StatusCode, body)
	}

	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	return string(body), nil
}
