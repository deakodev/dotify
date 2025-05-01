package dot

import "fmt"

type Registry struct {
	graphs   map[ID]*Graph
	activeID ID
}

func (r *Registry) Add(strict bool, t, id string) {
	if _, exists := r.graphs[ID(id)]; exists {
		return
	}
	r.graphs[ID(id)] = makeGraph(strict, GraphType(t), ID(id))
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
