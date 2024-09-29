package main

import (
	"bufio"
	"fmt"
	"github.com/ekou123/pokedex/commands"
	"os"
)

func main() {
	fmt.Println("Hello World")

	for {
		reader := bufio.NewReader(os.Stdin)

		userInput, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println(userInput)

		fmt.Println(commands.Test())

	}
}
