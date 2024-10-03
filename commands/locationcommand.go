package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MapLocation struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const baseURL = "https://pokeapi.co/api/v2/"

func ListLocations(pageURL *string) (MapLocation, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return MapLocation{}, err
	}

	defer res.Body.Close()

	var data []byte
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return MapLocation{}, err
	}

	var location MapLocation
	err = json.Unmarshal(data, &location)
	if err != nil {
		return MapLocation{}, fmt.Errorf("unable to unmarshal")
	}

	for _, location := range location.Results {
		fmt.Println(location.Name)
	}

	return location, nil

}
