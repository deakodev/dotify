package dot

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type ID string

type AttrTarget int

const (
	AttrTargetNone AttrTarget = iota
	AttrTargetGraph
	AttrTargetNode
	AttrTargetEdge
)

var attrTargetName = map[AttrTarget]string{
	AttrTargetNone:  "",
	AttrTargetGraph: "graph",
	AttrTargetNode:  "node",
	AttrTargetEdge:  "edge",
}

func (at AttrTarget) String() string {
	return attrTargetName[at]
}

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
	IdKey   ID
	IdValue ID
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
	return fmt.Sprintf("%s=%s", string(a.IdKey), string(a.IdValue))
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
	if s.target == AttrTargetNone {
		return attrs
	}
	attrs = strings.TrimSuffix(attrs, ";")
	return fmt.Sprintf("%s [%s]", s.target.String(), attrs)
}

func (s NodeStmt) String() string {
	return fmt.Sprintf("%s [%s]", string(s.id), strings.TrimSuffix(s.attrList.String(), ";"))
}

func (s EdgeStmt) String() string {
	return fmt.Sprintf("%s %s %s [%s]", string(s.idFrom), string(s.op), string(s.idTo), strings.TrimSuffix(s.attrList.String(), ";"))
}

func (g Graph) String() string {
	var sb strings.Builder
	if g.strict {
		sb.WriteString("strict ")
	}
	sb.WriteString(fmt.Sprintf("%s %s {\n", string(g.t), string(g.id)))
	for _, stmt := range g.stmtList {
		sb.WriteString(fmt.Sprintf("    %s\n", stmt.String()))
	}
	sb.WriteString("}")
	return sb.String()
}

func (g *Graph) Name() ID {
	return g.id
}

func (g *Graph) Type() GraphType {
	return g.t
}

func (g *Graph) Strict() bool {
	return g.strict
}

func (g *Graph) Print() {
	lines := strings.Split(g.String(), "\n")
	for i, line := range lines {
		fmt.Printf("%d |    %s\n", i, line)
	}
}

func Make(strict bool, t GraphType, id ID) *Graph {
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

func ParseStmts(lines []string) (*Graph, error) {
	var graph *Graph
	var graphName ID
	var graphType GraphType
	var isStrict bool

	stmtRegex := regexp.MustCompile(`^\s*(strict\s+)?(graph|digraph)\s+([a-zA-Z0-9_-]+)\s*\{\s*$`)
	endRegex := regexp.MustCompile(`^\s*\}\s*$`)
	edgeRegex := regexp.MustCompile(`^\s*(\w+)\s*(--|->)\s*(\w+)\s*(\[(.*?)\])?;?\s*$`)
	nodeRegex := regexp.MustCompile(`^\s*(\w+)\s*\[(.+)\]\s*;?\s*$`)
	// attrRegex := regexp.MustCompile(`^\s*(graph|node|edge)\s*\[(.+)\]\s*;?\s*$`)
	// plainAttrRegex := regexp.MustCompile(`^\s*(graph|node|edge)\s+((\w+\s*=\s*[^;\s]+)(\s*;\s*\w+\s*=\s*[^;\s]+)*)\s*;?$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if stmtMatch := stmtRegex.FindStringSubmatch(line); stmtMatch != nil {
			isStrict = stmtMatch[1] != ""
			graphType = GraphType(stmtMatch[2])
			graphName = ID(stmtMatch[3])
			graph = Make(isStrict, graphType, graphName)
			continue
		}

		if endRegex.MatchString(line) {
			break
		}

		if edgeMatch := edgeRegex.FindStringSubmatch(line); edgeMatch != nil {
			fromID := ID(edgeMatch[1])
			op := EdgeOp(edgeMatch[2])
			toID := ID(edgeMatch[3])
			attrStr := edgeMatch[5]
			attrList := parseAttrList(attrStr)
			graph.AddEdge(fromID, toID, op, attrList)
			continue
		}

		if nodeMatch := nodeRegex.FindStringSubmatch(line); nodeMatch != nil {
			id := ID(nodeMatch[1])
			attrs := parseAttrList(nodeMatch[2])
			graph.AddNode(id, attrs)
			continue
		}

		// if attrMatch := attrRegex.FindStringSubmatch(line); attrMatch != nil {
		// 	target := AttrTarget(attrMatch[1])
		// 	attrs := parseAttrList(attrMatch[2])
		// 	graph.AddAttr(target, attrs)
		// 	continue
		// }

		// if plainAttrMatch := plainAttrRegex.FindStringSubmatch(line); plainAttrMatch != nil {
		// 	target := AttrTarget(plainAttrMatch[1])
		// 	attrStr := plainAttrMatch[2]
		// 	attrs := parseAttrList(attrStr)
		// 	graph.AddAttr(target, attrs)
		// 	continue
		// }

		return nil, fmt.Errorf("unrecognized line: %q", line)
	}

	if graph == nil {
		return nil, errors.New("no valid graph declaration found")
	}

	return graph, nil
}

func parseAttrList(s string) AttrList {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := regexp.MustCompile(`[;,]`).Split(s, -1)
	var list AttrList
	for _, part := range parts {
		kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(kv) == 2 {
			list = append(list, AttrKeyValue{
				IdKey:   ID(strings.TrimSpace(kv[0])),
				IdValue: ID(strings.TrimSpace(kv[1])),
			})
		}
	}
	return list
}
