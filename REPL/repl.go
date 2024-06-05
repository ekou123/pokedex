package REPL

import (
	"encoding/json"
	"io"
	"net/http"
)

type Config struct {
	NextURL     *string
	PreviousURL *string
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





func GetAPIData(url *string) (AreaInfo, error) {

	//client := http.Client{Timeout: time.Duration(1) * time.Second}

	 resp, err := http.Get(*url)
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


