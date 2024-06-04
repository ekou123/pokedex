package REPL

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	NextURL     string
	PreviousURL string
}

type AreaInfo struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var InitialURL = "https://pokeapi.co/api/v2/"

func GetAPIData(url string) (AreaInfo, error) {
	fmt.Println("Deez")

	//fullURL := initialURL + url

	fullURL := "https://pokeapi.co/api/v2/location-area/"
	//client := http.Client{Timeout: time.Duration(1) * time.Second}

	 resp, err := http.Get(fullURL)
	 if err != nil {
		return AreaInfo{}, err
	 }

	 body, err := io.ReadAll(resp.Body)
	 if err != nil {
		return AreaInfo{}, err
	 }

	 programmableBody := AreaInfo{}
	 // Unmarshal

	 json.Unmarshal(body, &programmableBody)

	 return programmableBody, nil
}
