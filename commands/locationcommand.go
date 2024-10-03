package commands

import (
	"fmt"
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
	location, err := GetLocationData(pageURL)
	if err != nil {
		return MapLocation{}, err
	}

	for _, location := range location.Results {
		fmt.Println(location.Name)
	}

	return location, nil

}
