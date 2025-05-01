package repl

import (
	"flag"
	"fmt"

	"github.com/deakodev/go-remedy/internal/dot"
)

var (
	peekFlagSet   = flag.NewFlagSet("peek", flag.ContinueOnError)
	graphFlagSet  = flag.NewFlagSet("graph", flag.ContinueOnError)
	importFlagSet = flag.NewFlagSet("import", flag.ContinueOnError)
	exportFlagSet = flag.NewFlagSet("export", flag.ContinueOnError)
	addFlagSet    = flag.NewFlagSet("add", flag.ContinueOnError)
)

func init() {
	// peek
	peekFlagSet.String("g", "", "Prints graph. Defaults to active graph, if no graph is specified.")
	peekFlagSet.String("r", "", "Prints the graph registry.")

	peekFlagSet.Usage = func() {
		println("Usage: peek [flags]")
		peekFlagSet.PrintDefaults()
	}

	// graph
	graphFlagSet.String("g", "", "Graph to render. Defaults to active graph.")
	graphFlagSet.String("output", "graph.png", "Output image file name.")
	graphFlagSet.String("format", "png", "Output format (png, svg, etc.).")

	graphFlagSet.Usage = func() {
		fmt.Println("Usage: graph [flags]")
		graphFlagSet.PrintDefaults()
	}

	// import
	importFlagSet.String("f", "", "File to import.")
	importFlagSet.String("file", "", "File to import (alias).")
	importFlagSet.String("format", "dot", "Format of the graph file (dot, json, etc.).")

	importFlagSet.Usage = func() {
		fmt.Println("Usage: import [flags]")
		importFlagSet.PrintDefaults()
	}

	// export
	exportFlagSet.String("f", "", "File to export to.")
	exportFlagSet.String("file", "", "File to export to (alias).")
	exportFlagSet.String("format", "dot", "Export format (dot, json, etc.).")

	exportFlagSet.Usage = func() {
		fmt.Println("Usage: export [flags]")
		exportFlagSet.PrintDefaults()
	}

	// add
	addFlagSet.String("type", "", "What to add: graph, node, or edge.")
	addFlagSet.String("name", "", "Name of the graph/node/edge.")
	addFlagSet.String("from", "", "Source node (for edge).")
	addFlagSet.String("to", "", "Destination node (for edge).")
	addFlagSet.Bool("s", false, "Whether to use strict attribute for graph.")

	addFlagSet.Usage = func() {
		fmt.Println("Usage: add [flags]")
		addFlagSet.PrintDefaults()
	}
}

// type flag struct {
// 	description string
// 	handler     func(registry *dot.Registry, arg string) error
// }

// func (f flag) NewFlagSet(flags map[string]flag) {
// 	panic("unimplemented")
// }

// var peekFlags = map[string]flag{
// 	"-r": {
// 		description: "Prints the graph registry.",
// 		handler:     flagPeekRegistry,
// 	},
// 	"-g": {
// 		description: "Prints graph. Defaults to active graph, if no graph is specified.",
// 		handler:     flagPeekGraph,
// 	},
// }

func flagPeekRegistry(registry *dot.Registry, _ string) error {
	registry.Print()
	return nil
}

func flagPeekGraph(registry *dot.Registry, graphId string) error {
	if graphId == "" {
		activeGraph, err := registry.Active()
		if err != nil {
			return err
		}
		activeGraph.Print()
		return nil
	}

	graph, err := registry.Find(graphId)
	if err != nil {
		return err
	}
	graph.Print()
	return nil
}
