package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/inpatr/golang-pokedex/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		var input string
		for scanner.Scan() {
			input = scanner.Text()
			cleanedInput := repl.CleanInput(input)
			command, ok := repl.CommandRegistry[cleanedInput[0]] 
			if !ok {
				fmt.Printf("Unknown command")
			} else {
				config := &repl.Config{}
				err := command.Callback(config)
					if err != nil {
						fmt.Printf("%v", err)
					}
			}
			break
		}
		continue
	}
}
