package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replRun() {
	commandMapInit()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nRemedy > ")
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
