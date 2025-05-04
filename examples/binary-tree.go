package main

import (
	"fmt"

	"github.com/deakodev/go-dotify"
	"github.com/deakodev/go-dotify/dot"
)

func init() {
	// (1) Create a custom loader function for your custom type (eg. BinaryTree defined below)
	btLoader := func(raw any) (dotify.Dotifiable, error) {
		tree, ok := raw.(*BinaryTree)
		if !ok {
			return nil, fmt.Errorf("expected *BinaryTree")
		}
		return tree, nil
	}

	// (2) Register the loader function with a name
	// This name will be used to identify the type when calling dotify.This
	// and should match the type name used in the Dotify method.
	dotify.Register("BinaryTree", btLoader)
}

func main() {
	// (3) Your program logic here
	// Let's create a simple binary tree from BinaryTree (or your custom data structure)
	tree := &BinaryTree{
		Root: &MyNode{
			Key: 1,
			Left: &MyNode{
				Key: 2,
				Left: &MyNode{
					Key: 4,
				},
				Right: &MyNode{
					Key: 5,
				},
			},
			Right: &MyNode{
				Key: 3,
				Left: &MyNode{
					Key: 6,
				},
				Right: &MyNode{
					Key: 7,
				},
			},
		},
	}

	// (4) Call dotify.This to get the dot.Graph representation of the tree
	// The first argument is the name of the type as registered in the Registry
	// The second argument is the actual data to be converted to dot notation (the tree in this case)
	dot, err := dotify.This("BinaryTree", tree)
	if err != nil {
		fmt.Println("Error visualizing tree:", err)
		return
	}

	// (5) Render the dot.Graph to a PNG file
	// The first argument is the dot.Graph object
	// The second argument is the relative output path
	err = dotify.RenderPNG(dot, "out")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}

// USER DEFINED DOTIFIABLE INTERFACE

// Below is an example of a custom data structure (BinaryTree) that implements the Dotifiable interface.
// It has a method Type() that returns the type name and a method Dotify()
// that generates a dot.Graph representation of the binary tree.

type BinaryTree struct {
	Root *MyNode
}

type MyNode struct {
	Key   int
	Left  *MyNode
	Right *MyNode
}

func (t *BinaryTree) Type() string {
	return "MyBinaryTree"
}

func (t *BinaryTree) Dotify() *dot.Graph {
	g := dot.Make(false, "digraph", "BinaryTree")
	var id int

	var visit func(n *MyNode) string
	visit = func(n *MyNode) string {
		if n == nil {
			return ""
		}
		nodeID := fmt.Sprintf("n%d", id)
		id++

		g.AddNode(dot.ID(nodeID), dot.AttrList{
			{IdKey: "label", IdValue: dot.ID(fmt.Sprintf("\"%d\"", n.Key))},
		})

		if n.Left != nil {
			leftID := visit(n.Left)
			g.AddEdge(dot.ID(nodeID), dot.ID(leftID), dot.Directed, nil)
		}
		if n.Right != nil {
			rightID := visit(n.Right)
			g.AddEdge(dot.ID(nodeID), dot.ID(rightID), dot.Directed, nil)
		}

		return nodeID
	}

	visit(t.Root)
	return g
}
