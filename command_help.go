package main

import (
	"fmt"
)
func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to pokedex Help Menu!(●'◡'●)")
	fmt.Println("Here are the available Commands")
	availablecommands := getCommads()
	for _, cmd := range availablecommands{
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}

	return nil
}