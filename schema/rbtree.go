package schema

import (
	"fmt"

	"github.com/deakodev/go-dotify/dot"
	"github.com/emirpasic/gods/trees/redblacktree"
)

type RBTreeSchema redblacktree.Tree

func (r RBTreeSchema) Type() string {
	return "RBTree"
}

func (r RBTreeSchema) Dotify() *dot.Graph {
	graph := dot.Make(false, "digraph", "RBTree")
	var nodeCount int
	nodeMap := make(map[*redblacktree.Node]string)

	var build func(node *redblacktree.Node) string
	build = func(node *redblacktree.Node) string {
		if node == nil {
			return ""
		}
		nodeID := fmt.Sprintf("n%d", nodeCount)
		nodeMap[node] = nodeID
		nodeCount++

		graph.AddNode(dot.ID(nodeID), dot.AttrList{
			{IdKey: "label", IdValue: dot.ID(fmt.Sprintf("\"%v\"", node.Key))},
		})

		if node.Left != nil {
			leftID := build(node.Left)
			graph.AddEdge(dot.ID(nodeID), dot.ID(leftID), dot.Directed, nil)
		}
		if node.Right != nil {
			rightID := build(node.Right)
			graph.AddEdge(dot.ID(nodeID), dot.ID(rightID), dot.Directed, nil)
		}

		return nodeID
	}

	build(r.Root)
	return graph
}
