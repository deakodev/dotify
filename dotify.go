package dotify

import (
	"fmt"

	"github.com/deakodev/go-dotify/dot"
)

// USER DEFINED INTERFACE
// Dotifiable is an interface that requires implementing the Dotify method
// which returns a dot.Graph representation of the data structure.
type Dotifiable interface {
	Dotify() *dot.Graph
	Type() string
}

var Registry = make(map[string]func(any) (Dotifiable, error))

// Register function allows users to register their own types with a loader function.
// The loader function takes raw data and returns a Dotifiable object or an error.
func Register(name string, loader func(any) (Dotifiable, error)) {
	Registry[name] = loader
}

// 'This' takes a registered type name and raw data, looks up the corresponding
// loader function in the Registry, and uses it to create a Dotifiable object.
// It then calls the Dotify method on that object to get a dot.Graph representation.
func This(typeName string, rawData any) (*dot.Graph, error) {
	loader, ok := Registry[typeName]
	if !ok {
		return nil, fmt.Errorf("type %s not registered", typeName)
	}

	dotifiable, err := loader(rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to load %s: %w", typeName, err)
	}

	return dotifiable.Dotify(), nil
}

func RenderPNG(graph *dot.Graph, outPath string) error {
	return dot.Render(graph, outPath, "png")
}

func RenderSVG(graph *dot.Graph, outPath string) error {
	return dot.Render(graph, outPath, "svg")
}
