package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ekou123/pokedex/command_functionality"
	
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//resp, err := http.Get()
	

	for {
		
		fmt.Print("pokedex > ")
		scanner.Scan()
		userChoice := scanner.Text()

		switch  strings.ToLower(userChoice) {
		case "help":
			command_functionality.CommandHelp()
		case "exit":
			command_functionality.CommandExit()
		}




	}
}