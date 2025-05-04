package dot

import "fmt"

type Registry struct {
	graphs   map[ID]*Graph
	activeID ID
}

func (r *Registry) Make(strict bool, t, id string) error {
	if t != "graph" && t != "digraph" {
		return fmt.Errorf("invalid graph type: %q (expected 'graph' or 'digraph')", t)
	}

	if id == "" {
		return fmt.Errorf("graph name is required")
	}

	if _, exists := r.graphs[ID(id)]; exists {
		return fmt.Errorf("graph %q already exists", id)
	}

	r.graphs[ID(id)] = Make(strict, GraphType(t), ID(id))
	r.activeID = ID(id)
	return nil
}

func (r *Registry) Add(graph *Graph) {
	r.graphs[graph.Name()] = graph
	r.activeID = graph.Name()
}

func (r *Registry) Active() (*Graph, error) {
	graph, exists := r.graphs[r.activeID]
	if !exists {
		return nil, fmt.Errorf("graph %s does not exist", r.activeID)
	}
	return graph, nil
}

func (r *Registry) Find(id string) (*Graph, error) {
	graph, exists := r.graphs[ID(id)]
	if !exists {
		return nil, fmt.Errorf("graph %s does not exist", ID(id))
	}
	return graph, nil
}

func (r *Registry) Print() {
	fmt.Printf("Registry (%v):\n", len(r.graphs))
	for id, graph := range r.graphs {
		fmt.Printf("%s\n", id)
		fmt.Println(graph.String())
	}
}

func MakeRegistry() *Registry {
	return &Registry{
		graphs:   make(map[ID]*Graph),
		activeID: "",
	}
}
