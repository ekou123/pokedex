package command_functionality

import (
    "os"
    "fmt"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func()
}

func GetCliCommands() map[string]CliCommand{
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
            Name: "map",
            Description: "Shows next 20 area locations",
            Callback: CommandMap,
        },
	}

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