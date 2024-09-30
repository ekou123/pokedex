package commands

import "fmt"

type CliCommand struct {
	Name        string
	Description string
	Callback    func()
}

func CommandList() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
	}
}

func CommandHelp() {
	commands := CommandList()

	for _, command := range commands {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println(command.Name + ": \t" + command.Description)
	}
}
