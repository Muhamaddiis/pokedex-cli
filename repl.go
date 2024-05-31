package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		availablecommands := getCommads()

		command, ok := availablecommands[commandName]
		if !ok {
			fmt.Println("Invalid Command😕")
			continue
		}
		err := command.callback(cfg)
		if err != nil{
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommads() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help Menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Lists some Location area",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous page of location areas",
			callback:    callbackMapBack,
		},
		"exit": {
			name:        "exit",
			description: "exits the program",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
