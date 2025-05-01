package dot

import (
	"fmt"
	"strings"
)

type ID string

type AttrTarget string

const (
	NoneTarget  AttrTarget = "none"
	GraphTarget AttrTarget = "graph"
	NodeTarget  AttrTarget = "node"
	EdgeTarget  AttrTarget = "edge"
)

type GraphType string

const (
	DigraphYes GraphType = "digraph"
	DigraphNo  GraphType = "graph"
)

type EdgeOp string

const (
	Directed   EdgeOp = "->"
	Undirected EdgeOp = "--"
)

type Graph struct {
	stmtList []Stmt
	id       ID
	t        GraphType
	strict   bool
}

type NodeStmt struct {
	attrList AttrList
	id       ID
}

type EdgeStmt struct {
	attrList AttrList
	idFrom   ID
	idTo     ID
	op       EdgeOp
}

type AttrKeyValue struct {
	idKey   ID
	idValue ID
}

type AttrList []AttrKeyValue

type AttrStmt struct {
	attrList AttrList
	target   AttrTarget
}

type Stmt interface {
	String() string
}

func (a AttrKeyValue) String() string {
	return fmt.Sprintf("%s=%s", string(a.idKey), string(a.idValue))
}

func (l AttrList) String() string {
	var sb strings.Builder
	for _, attrPair := range l {
		sb.WriteString(fmt.Sprintf("%s;", attrPair.String()))
	}
	return sb.String()
}

func (s AttrStmt) String() string {
	attrs := s.attrList.String()
	if s.target == NoneTarget {
		return attrs
	}

	attrs = strings.TrimSuffix(attrs, ";")
	return fmt.Sprintf("%s [%s]", string(s.target), attrs)
}

func (s NodeStmt) String() string {
	return fmt.Sprintf("%s %s", string(s.id), s.attrList.String())
}

func (s EdgeStmt) String() string {
	return fmt.Sprintf("%s %s %s %s", string(s.idFrom), string(s.op), string(s.idTo), s.attrList.String())
}

func (g Graph) String() string {
	var sb strings.Builder

	if g.strict {
		sb.WriteString("strict ")
	}

	sb.WriteString(fmt.Sprintf("%s %s {\n", string(g.t), string(g.id)))

	for _, stmt := range g.stmtList {
		sb.WriteString(fmt.Sprintf("%s\n", stmt.String()))
	}

	sb.WriteString("}")

	return sb.String()
}

func (g *Graph) Print() {
	lines := strings.Split(g.String(), "\n")
	for i, line := range lines {
		fmt.Printf("%d |    %s\n", i, line)
	}
}

func makeGraph(strict bool, t GraphType, id ID) *Graph {
	return &Graph{
		id:       id,
		t:        t,
		strict:   strict,
		stmtList: []Stmt{},
	}
}

func (g *Graph) addStmt(stmt Stmt) {
	g.stmtList = append(g.stmtList, stmt)
}

func (g *Graph) AddNode(id ID, attrs AttrList) {
	g.addStmt(NodeStmt{id: id, attrList: attrs})
}

func (g *Graph) AddEdge(from, to ID, op EdgeOp, attrs AttrList) {
	g.addStmt(EdgeStmt{idFrom: from, idTo: to, op: op, attrList: attrs})
}

func (g *Graph) AddAttr(target AttrTarget, attrs AttrList) {
	g.addStmt(AttrStmt{target: target, attrList: attrs})
}
