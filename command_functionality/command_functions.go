package command_functionality



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
		"mapb": {
			Name: "mapb",
			Description: "Shows previous 20 area locations",
			Callback: CommandMap,
		},
	}

}



