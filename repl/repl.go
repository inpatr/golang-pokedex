package repl

import (
	"strings"
	"os"
	"errors"
	"fmt"
	"github.com/inpatr/golang-pokedex/requests"
)

type Config struct {
	next	*string `json:"next"`
	previous	*string `json:"previous"`
}

type cliClommand struct {
	Name	string
	Description	string
	Callback	func(*Config) error
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
		"map": {
			Name:	"map",
			Description:	"Displays next 20 locations on each call",
			Callback: CommandMap,
		},
	}
}

func CleanInput(text string) []string {
	return strings.Fields(text)
}

func CommandExit(conf *Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("Could not exit!")
}

func CommandHelp(conf *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range CommandRegistry {
		fmt.Println(command.Name + ": " + command.Description)
	}
	return nil
}

func CommandMap(conf *Config) error {
	var url string
	if conf.next != nil && len(*conf.next) > 0 {
		url = *conf.next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := requests.GetLocationAreas(url)
	if err != nil {
		return err
	}

	if res.Next != nil {
		conf.next = res.Next
	}
	if res.Previous != nil {
		conf.previous = res.Previous
	}
	
	var locationAreaList []string

	for i := 0; i < len(res.Results); i++ {
		locationAreaList = append(locationAreaList, res.Results[i].Name)
	}

	for _, locationArea := range locationAreaList {
		fmt.Println(locationArea)
	}

	return nil
}
