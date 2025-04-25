package main

import (
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	callback    func() error
}

var commandMap map[string]command

func commandMapInit() {
	commandMap = map[string]command{
		"exit": {
			name:        "exit",
			description: "Exit remedy.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
	}
}

func commandExit() error {
	fmt.Printf("\nExiting remedy...\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("\nUsage:\n")
	for _, c := range commandMap {
		fmt.Printf("[%s] -> %s\n", c.name, c.description)
	}
	return nil
}
