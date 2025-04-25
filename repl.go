package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replRun() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Remedy > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Printf("scanner failed: %v\n", err)
			return
		}

		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		cleanText := cleanInput(text)
		command := cleanText[0]
		fmt.Printf("Command: %s\n", command)

		handler := commandMap[command].callback
		if handler == nil {
			fmt.Printf("command handler unknown.")
			return
		}

		err := handler()
		if err != nil {
			fmt.Printf("command handler error: %v", err)
			return
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Printf("Exiting remedy...")
	os.Exit(0)
	return nil
}

type command struct {
	name        string
	description string
	callback    func() error
}

var commandMap = map[string]command{
	"exit": {
		name:        "exit",
		description: "Exit remedy.",
		callback:    commandExit,
	},
}
