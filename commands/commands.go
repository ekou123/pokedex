package commands

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func CommandList() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays next 20 locations",
			Callback:    CommandMap,
		},
	}
}

func CommandExit(cfg *Config) error {
	os.Exit(1)
	return nil
}

func CommandHelp(cfg *Config) error {
	commands := CommandList()
	for _, command := range commands {

		fmt.Printf("%s: \t %s \n", command.Name, command.Description)
	}
	return nil
}

func CommandMap(cfg *Config) error {

	locationsResp, err := ListLocations(cfg.NextURL)

	if err != nil {
		return err
	}

	cfg.NextURL = locationsResp.Next
	cfg.PreviousURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandMapB(cfg *Config) {
	if cfg.PreviousURL == nil {
		*cfg.currentURL = baseURL + "location-area?offset=20&limit=20"
		currentLocations, _ := ListLocations(cfg.currentURL)
		cfg.NextURL = currentLocations.Next
	}

	cfg.currentURL = cfg.PreviousURL

	location, _ := ListLocations(cfg.PreviousURL)

	cfg.NextURL = location.Next
	cfg.PreviousURL = location.Previous

}
