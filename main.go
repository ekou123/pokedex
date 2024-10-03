package main

import (
	"bufio"
	"example.com/pokedex/commands"
	"fmt"
	"os"
	"strings"
)

func main() {

	cfg := &commands.Config{}

	startRepl(cfg)

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(cfg *commands.Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands.CommandList()[commandName]

		if exists {
			err := command.Callback(cfg, commandName)
			if err != nil {
				return
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
