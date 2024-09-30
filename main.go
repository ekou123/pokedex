package main

import (
	"bufio"
	"fmt"
	"github.com/ekou123/pokedex/commands"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	command := commands.CommandList()

	fmt.Println(command)

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Pokedex > ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		switch userInput {
		case strings.ToLower("help"):
		}

	}
}
