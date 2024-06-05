package command_functionality

import (
	"fmt"
	"os"

	"github.com/ekou123/pokedex/REPL"
)

var initialURL = "https://pokeapi.co/api/v2/location-area"

var config = REPL.Config {
	NextURL: &initialURL,
	PreviousURL: nil,
}

func CommandExit() {
	os.Exit(0)
}

func CommandHelp() {
    cliCommands := GetCliCommands()

    fmt.Println("\nAVAILABLE COMMANDS:")

    for _, cliCommand := range cliCommands {
        fmt.Println(cliCommand.Name + ": " + cliCommand.Description)
    }
    fmt.Println()
}


func CommandMap() {
	if *config.NextURL == "" {
		*config.PreviousURL = *config.NextURL
	}
	programmableBody, _ := REPL.GetAPIData(config.NextURL)
	
	config.NextURL = &programmableBody.Next
	config.PreviousURL = &programmableBody.Previous

	for _, name := range programmableBody.Results {
		fmt.Println(name.Name)

	}


}

func CommandMapB() {
	if *config.PreviousURL == "" {
		*config.NextURL = initialURL
	}
	

	fmt.Println(*config.PreviousURL)
	programmableBody, _ := REPL.GetAPIData(config.PreviousURL)
	
	config.NextURL = &programmableBody.Next
	config.PreviousURL = &programmableBody.Previous

	for _, name := range programmableBody.Results {
		fmt.Println(name.Name)

	}


}