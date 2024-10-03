package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Config struct {
	PreviousURL *string
	currentURL  *string
	NextURL     *string
}

func GetExploreData(pageURL *string) (PokemonArea, error) {
	if pageURL == nil {
		return PokemonArea{}, errors.New("No URL provided")
	}

	resp, err := http.Get(*pageURL)
	if err != nil {
		return PokemonArea{}, err
	}

	defer resp.Body.Close()

	var pokemonArea PokemonArea
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonArea{}, err
	}

	err = json.Unmarshal(body, &pokemonArea)
	if err != nil {
		return PokemonArea{}, err
	}

	return pokemonArea, nil

}

func GetLocationData(pageURL *string) (MapLocation, error) {
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

	return location, nil
}

func CheckIfCanExplore(locationData MapLocation, userRequest string) (string, bool) {
	for _, location := range locationData.Results {
		if strings.ToLower(location.Name) == strings.ToLower(userRequest) {
			return location.URL, true
		}
	}
	return "", false
}
