package commands

import (
	"errors"
	"io"
	"net/http"
)

type Config struct {
	PreviousURL *string
	currentURL  *string
	NextURL     *string
}

func getExploreData(pageURL *string) error {
	if pageURL == nil {
		return errors.New("No URL provided")
	}

	resp, err := http.Get(*pageURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var pokemonArea PokemonArea
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

}
