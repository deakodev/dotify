package graph

import (
	"fmt"
	"strings"
)

// Node represents a node in the graph with properties like shape and label.
type Node struct {
	label string
	// style string
	// shape string
	// color string
}

// Edge represents an edge in the graph, connecting two nodes.
// n1 -> n2
type Edge struct {
	// label  string
	// style  string
	// color  string
	// weight int
	From *Node
	To   *Node
}

// Graph represents the entire directed graph, with nodes and edges.
type Graph struct {
	Nodes map[string]*Node
	Edges []Edge
}

func (g *Graph) AddNode(label string) *Node {
	g.Nodes[label] = &Node{label: label}
	return g.Nodes[label]
}

type GraphKey string

type Registry struct {
	Graphs    map[GraphKey]*Graph
	ActiveKey GraphKey
}

func (r *Registry) Active() (*Graph, error) {
	graph, exists := r.Graphs[r.ActiveKey]
	if !exists {
		return nil, fmt.Errorf("graph %s does not exist", r.ActiveKey)
	}
	return graph, nil
}

func (r *Registry) Find(key GraphKey) (*Graph, error) {
	graph, exists := r.Graphs[key]
	if !exists {
		return nil, fmt.Errorf("graph %s does not exist", key)
	}
	return graph, nil
}

func (g *Graph) ToDot() string {
	var sb strings.Builder

	sb.WriteString("digraph {\n")

	for _, edge := range g.Edges {
		fromLabel := edge.From.label
		toLabel := edge.To.label

		sb.WriteString(fmt.Sprintf("    \"%s\" -> \"%s\";\n", fromLabel, toLabel))
	}

	sb.WriteString("}\n")

	return sb.String()
}

func (g *Graph) FromDot(dot string) error {
	stmtList := strings.Split(dot, "\n")

	for _, stmt := range stmtList {
		stmt = strings.TrimSuffix(strings.TrimSpace(stmt), ";")

		// switch stmt {
		// 	case
		// }

		if stmt == "" || strings.HasPrefix(stmt, "digraph") || stmt == "{" || stmt == "}" {
			continue
		}

		// parts := strings.Split(stmt, "->")
		// if len(parts) != 2 {
		// 	return fmt.Errorf("invalid line: %s", stmt)
		// }

		// fromLabel := strings.Trim(parts[0], " \"")
		// toLabel := strings.Trim(parts[1], " \"")

		// fromNode, ok := nodeMap[fromLabel]
		// if !ok {
		// 	graph.AddNode(fromLabel)
		// 	nodeMap[fromLabel] = fromNode
		// }

		// toNode, ok := nodeMap[toLabel]
		// if !ok {
		// 	graph.AddNode(toLabel)
		// 	nodeMap[toLabel] = toNode
		// }

		// graph.Edges = append(graph.Edges, Edge{
		// 	From: fromNode,
		// 	To:   toNode,
		// })
	}

	return nil
}
