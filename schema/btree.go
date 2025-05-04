package schema

import (
	"fmt"
	"strings"

	"github.com/deakodev/go-dotify/dot"
	"github.com/emirpasic/gods/trees/btree"
)

type BTreeSchema btree.Tree

func (b BTreeSchema) Type() string {
	return "BTree"
}

func (s BTreeSchema) Dotify() *dot.Graph {
	graph := dot.Make(false, "digraph", "BTree")
	var nodeCount int

	var build func(node *btree.Node) string
	build = func(node *btree.Node) string {
		if node == nil {
			return ""
		}
		nodeID := fmt.Sprintf("n%d", nodeCount)
		nodeCount++

		// Create a label for all entries in this node
		var keys []string
		for _, entry := range node.Entries {
			keys = append(keys, fmt.Sprintf("%v", entry.Key))
		}
		label := strings.Join(keys, "|")

		graph.AddNode(dot.ID(nodeID), dot.AttrList{
			{IdKey: "label", IdValue: dot.ID(fmt.Sprintf("\"%s\"", label))},
			{IdKey: "shape", IdValue: dot.ID("record")},
		})

		// Each node has len(Entries)+1 children
		for _, child := range node.Children {
			childID := build(child)
			graph.AddEdge(dot.ID(nodeID), dot.ID(childID), dot.Directed, nil)
		}

		return nodeID
	}

	build(s.Root)
	return graph
}
