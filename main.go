package main

import (
	"bufio"
	"github.com/ekou123"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World")

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
