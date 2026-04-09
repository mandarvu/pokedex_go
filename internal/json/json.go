// Package json (used as ijson to signify internal nature of this package)
package json

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	cache "github.com/mandarvu/pokedex_go/internal/pokecache"
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

var CachedData cache.Cache = cache.NewCache(5 * time.Millisecond)

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

func GetResponseFromURL(u string, cachedData *cache.Cache) ([]byte, error) {
	if data, ok := cachedData.Get(u); ok {
		return data, nil
	}

	res, err := http.Get(u)
	if err != nil {
		return []byte{}, fmt.Errorf("could Not GET response from URL %s: %v", u, err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("response failed with status code: %d and\nbody: %sn", res.StatusCode, body)
	}

	if err != nil {
		return []byte{}, fmt.Errorf("%v", err)
	}

	cachedData.Add(u, body)

	return body, nil
}
