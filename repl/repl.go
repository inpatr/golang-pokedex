package repl

import (
	"strings"
	"os"
	"errors"
	"fmt"
)

type cliClommand struct {
	Name	string
	Description	string
	Callback	func() error
}

var CommandRegistry map[string]cliClommand

func init() {
	CommandRegistry = map[string]cliClommand{
		"exit": {
			Name:	"exit",
			Description:	"Exit the Pokedex",
			Callback:	CommandExit,
		},
		"help": {
			Name:	"help",
			Description: "Displays a help message",
			Callback: CommandHelp,
		},
	}
}

func CleanInput(text string) []string {
	return strings.Fields(text)
}

func CommandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Could not exit!")
}

func CommandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range CommandRegistry {
		fmt.Println(command.Name + ": " + command.Description)
	}
	return nil
}
