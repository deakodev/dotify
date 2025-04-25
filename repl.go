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
		fmt.Printf("Command: %s\n", cleanText[0])
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
