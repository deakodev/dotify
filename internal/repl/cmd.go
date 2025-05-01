package repl

import (
	"flag"
	"fmt"
	"os"

	"github.com/deakodev/go-remedy/internal/dot"
)

type cmd struct {
	flagSet     *flag.FlagSet
	description string
	handler     func(registry *dot.Registry, args ...string) error
}

var cmds map[string]cmd

func init() {
	cmds = map[string]cmd{
		"exit": {
			description: "Exit tangle.",
			flagSet:     nil,
			handler:     cmdExit,
		},
		"help": {
			description: "Displays a help message.",
			flagSet:     nil,
			handler:     cmdHelp,
		},
		"peek": {
			description: "Prints the current registry.",
			flagSet:     peekFlagSet,
			handler:     cmdPeek,
		},
		"graph": {
			description: "Generate a graph image.",
			flagSet:     graphFlagSet,
			handler:     cmdGraph,
		},
		"import": {
			description: "Import a graph from a file.",
			flagSet:     importFlagSet,
			handler:     cmdUse,
		},
		"export": {
			description: "Export a graph to a file.",
			flagSet:     exportFlagSet,
			handler:     cmdWrite,
		},
		"add": {
			description: "Add a new graph, node, or edge.",
			flagSet:     addFlagSet,
			handler:     cmdAdd,
		},
	}
}

func (c cmd) parseFlags(args []string) (nonflagArgs []string, err error) {
	if c.flagSet != nil {
		if err := c.flagSet.Parse(args); err != nil {
			return nil, fmt.Errorf("error parsing flags: %v", err)
		}

		c.flagSet.Visit(func(f *flag.Flag) {
			fmt.Printf("Flag %s: %s\n", f.Name, f.Value.String())
		})

		return c.flagSet.Args(), nil
	}
	return nil, nil // some cmds dont have flags
}

func cmdExit(registry *dot.Registry, args ...string) error {
	fmt.Printf("\nExiting tangle...\n")
	os.Exit(0)
	return nil
}

func cmdHelp(registry *dot.Registry, args ...string) error {
	fmt.Printf("\nUsage:\n")
	for k, v := range cmds {
		fmt.Printf("[%s] -> %s\n", k, v.description)
	}
	return nil
}

func cmdPeek(registry *dot.Registry, args ...string) error {
	fs := cmds["peek"].flagSet

	graphFlag := fs.Lookup("g")
	if graphFlag != nil && graphFlag.Value.String() != "" {
		graph, err := registry.Find(graphFlag.Value.String())
		if err != nil {
			return fmt.Errorf("failed to find graph: %w", err)
		}
		fmt.Printf("Graph: %s\n", graph.String())

	} else if graphFlag != nil {
		activeGraph, err := registry.Active()
		if err != nil {
			return fmt.Errorf("failed to get active graph: %w", err)
		}
		fmt.Printf("Active Graph: %s\n", activeGraph.String())
	}

	registryFlag := fs.Lookup("r")
	if registryFlag != nil {
		registry.Print()
	}

	return nil
}

func cmdAdd(registry *dot.Registry, args ...string) error {
	// fs := cmds["add"].flagSet

	// if len(params) == 0 {
	// 	return fmt.Errorf("no input provided")
	// }

	// // activeGraph, err := registry.Active()
	// // if err != nil {
	// // 	return fmt.Errorf("failed to get active graph: %w", err)
	// // }

	// switch params[0] {
	// case "-g":
	// 	if len(params) < 3 {
	// 		return fmt.Errorf("missing add graph args")
	// 	}
	// 	graphStrict := params[1] == "strict"
	// 	graphType := params[1]
	// 	graphID := params[2]
	// 	if graphStrict {
	// 		graphType = params[2]
	// 		graphID = params[3]
	// 	}
	// 	registry.Add(graphStrict, graphType, graphID)
	// 	return nil
	// case "-a":
	// 	if len(params) < 2 {
	// 		return fmt.Errorf("missing add node args")
	// 	}
	// 	// target := dot.AttrTarget(params[1])
	// 	// attrs := dot.AttrList(params[2:])
	// 	// activeGraph.AddAttr(target, attrs)
	// case "-n":
	// 	panic("not implemented")
	// case "-e":
	// 	panic("not implemented")
	// default:
	// 	return fmt.Errorf("unknown command: %s", params[0])
	// }

	return nil

	// g, err := registry.Active()
	// if err != nil {
	// 	return err
	// }

	// var lastNode *graph.Node

	// for i := 0; i < len(params); i++ {
	// 	token := params[i]
	// 	switch token {
	// 	case "->":
	// 		// Create an edge between lastNode and nextNode
	// 		if lastNode == nil || i+1 >= len(params) {
	// 			return fmt.Errorf("invalid syntax near '->'")
	// 		}
	// 		nextToken := params[i+1]

	// 		// Create the next node
	// 		nextNode := g.AddNode(nextToken)

	// 		// Create the edge
	// 		g.Edges = append(g.Edges, graph.Edge{
	// 			From: lastNode,
	// 			To:   nextNode,
	// 		})

	// 		lastNode = nextNode
	// 		i++ // skip the next token, already processed
	// 	default:
	// 		lastNode = g.AddNode(token)
	// 	}
	// }

	// fmt.Println("Graph content:\n", *g)
}

func cmdUse(registry *dot.Registry, args ...string) error {
	// g, err := registry.Active()
	// if err != nil {
	// 	return err
	// }

	// if len(params) == 0 {
	// 	return fmt.Errorf("no file path provided")
	// }

	// filePath := params[0]

	// data, err := os.ReadFile(filePath)
	// if err != nil {
	// 	return fmt.Errorf("failed to read file: %w", err)
	// }

	// dotGraph := string(data)
	// fmt.Println("Graph content:\n", dotGraph)

	// err = g.FromDot(dotGraph)
	// if err != nil {
	// 	return fmt.Errorf("failed to parse dot: %w", err)
	// }

	// fmt.Printf("graph: %v\n", *g)
	return nil
}

func cmdWrite(registry *dot.Registry, args ...string) error {
	// g, err := registry.Active()
	// if err != nil {
	// 	return err
	// }

	// if len(params) == 0 {
	// 	return fmt.Errorf("no file path provided")
	// }

	// filePath := params[0]

	// dotContent := g.ToDot() // Turn the graph into dot text

	// err = os.WriteFile(filePath, []byte(dotContent), 0644)
	// if err != nil {
	// 	return fmt.Errorf("failed to write file: %w", err)
	// }

	// fmt.Println("Graph written to", filePath)
	return nil
}

func cmdGraph(registry *dot.Registry, args ...string) error {
	// g, err := registry.Active()
	// if err != nil {
	// 	return err
	// }

	// payload := payload{g.ToDot(), params[0], params[1]}
	// fmt.Println("payload: ", payload)
	// // Convert the payload to JSON
	// payloadBytes, err := json.Marshal(payload)
	// if err != nil {
	// 	return fmt.Errorf("failed to marshal payload: %w", err)
	// }

	// // Make the POST request to QuickChart API
	// res, err := http.Post("https://quickchart.io/graphviz", "application/json", bytes.NewBuffer(payloadBytes))
	// if err != nil {
	// 	return fmt.Errorf("get failed: %w", err)
	// }
	// defer res.Body.Close()

	// // Check if the response status is OK
	// if res.StatusCode != http.StatusOK {
	// 	return fmt.Errorf("bad response: %s", res.Status)
	// }

	// // Save the response body (image) as a file
	// fileName := fmt.Sprintf("%s.%s", registry.ActiveKey, payload.Format)
	// filePath := path.Join("out", fileName)
	// outFile, err := os.Create(filePath)
	// if err != nil {
	// 	return fmt.Errorf("failed to create file: %w", err)
	// }
	// defer outFile.Close()

	// // Copy the response body to the file
	// _, err = io.Copy(outFile, res.Body)
	// if err != nil {
	// 	return fmt.Errorf("failed to write to file: %w", err)
	// }

	// fmt.Println("Diagram saved as graph.png")
	return nil
}
