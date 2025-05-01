package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/deakodev/go-remedy/internal/dot"
)

// type payload struct {
// 	Graph  string `json:"graph"`
// 	Layout string `json:"layout"`
// 	Format string `json:"format"`
// }

func Run(registry *dot.Registry) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nTangle > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Printf("scanner failed: %v\n", err)
			return
		}

		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		cmd, args := parseInput(input)
		args, err := cmds[cmd].parseFlags(args)
		if err != nil {
			fmt.Printf("Error parsing flags: %v\n", err)
			continue
		}

		cmds[cmd].handler(registry, args...)
	}
}

func parseInput(text string) (cmd string, args []string) {
	text = strings.TrimSpace(text)

	fields := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' // space as delimiter
	})

	if len(fields) > 0 {
		cmd = fields[0]
		args = fields[1:]
	}

	return cmd, args
}
