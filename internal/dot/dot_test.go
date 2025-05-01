package dot

import (
	"testing"
)

func TestAttrPairString(t *testing.T) {
	p := AttrKeyValue{ID("color"), ID("red")}
	expected := "color=red"
	if p.String() != expected {
		t.Errorf("expected %q, got %q", expected, p.String())
	}
}

func TestAttrPairStringWithQuotes(t *testing.T) {
	p := AttrKeyValue{ID("color"), ID(`"red"`)}
	expected := `color="red"`
	if p.String() != expected {
		t.Errorf("expected %q, got %q", expected, p.String())
	}
}

func TestAttrListString(t *testing.T) {
	l := AttrList{
		{ID("color"), ID("\"red\"")},
		{ID("shape"), ID("box")},
	}
	expected := `[color="red",shape=box]`
	if l.String() != expected {
		t.Errorf("expected %q, got %q", expected, l.String())
	}
}

func TestNodeStmtString(t *testing.T) {
	node := NodeStmt{
		id: "A",
		attrList: AttrList{
			{ID("label"), ID("Start")},
		},
	}
	expected := "A [label=Start]"
	if node.String() != expected {
		t.Errorf("expected %q, got %q", expected, node.String())
	}
}

func TestEdgeStmtString(t *testing.T) {
	edge := EdgeStmt{
		idFrom: "A",
		idTo:   "B",
		op:     Directed,
		attrList: AttrList{
			{ID("weight"), ID("2")},
		},
	}
	expected := "A -> B [weight=2]"
	if edge.String() != expected {
		t.Errorf("expected %q, got %q", expected, edge.String())
	}
}

func TestAttrStmtString(t *testing.T) {
	stmt := AttrStmt{
		target: GraphTarget,
		attrList: AttrList{
			{ID("rankdir"), ID("LR")},
		},
	}
	expected := "graph [rankdir=LR]"
	if stmt.String() != expected {
		t.Errorf("expected %q, got %q", expected, stmt.String())
	}
}

func TestGraphString(t *testing.T) {
	graph := Graph{
		id:     "G",
		t:      DigraphYes,
		strict: true,
		stmtList: []Stmt{
			AttrStmt{
				target: GraphTarget,
				attrList: AttrList{
					{ID("rankdir"), ID("LR")},
				},
			},
			NodeStmt{
				id: "A",
				attrList: AttrList{
					{ID("label"), ID("Start")},
				},
			},
			EdgeStmt{
				idFrom: "A",
				idTo:   "B",
				op:     Directed,
				attrList: AttrList{
					{ID("weight"), ID("1")},
				},
			},
		},
	}

	expected := `strict digraph G {
graph [rankdir=LR]
A [label=Start]
A -> B [weight=1]
}`

	if graph.String() != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, graph.String())
	}
}
