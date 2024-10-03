package commands

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config)
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

func CommandExit(cfg *Config) {
	os.Exit(1)
}

func CommandHelp(cfg *Config) {
	commands := CommandList()
	for _, command := range commands {

		fmt.Printf("%s: \t %s \n", command.Name, command.Description)
	}
}

func CommandMap(cfg *Config) {
	if cfg.NextURL == nil {
		fmt.Println("Cannot continue further soz")
	}

	location := ListLocations(cfg.NextURL)
	cfg.NextURL = location.Next
	cfg.PreviousURL = cfg.NextURL
}

func CommandMapB(cfg *Config) {
	if cfg.PreviousURL == nil {
		*cfg.currentURL = baseURL + "location-area?offset=20&limit=20"
		cfg.NextURL = ListLocations(cfg.currentURL).Next
		return
	}

	cfg.currentURL = cfg.PreviousURL

	location := ListLocations(cfg.PreviousURL)

	cfg.NextURL = location.Next
	cfg.PreviousURL = location.Previous

}
