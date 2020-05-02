package graph

import (
	"bytes"
	"fmt"
)

// Node defined a node with int64 ID and string name
type Node struct {
	id         int64
	name       string
	observable bool
}

// Path defines a path with node type
// fork - f   chain - h   collider - c    end - e
type Path struct {
	from, to int64
	nodes    []int64
	types    []byte
}

// Graph defines a directed acyclic graph
type Graph struct {
	names     map[string]int64
	nodes     map[int64]Node
	from      map[int64]map[int64]bool
	to        map[int64]map[int64]bool
	bidirect  map[int64]map[int64]bool
	maxNodeID int64
}

func (n Node) String() string {
	return fmt.Sprintf("[%v]%v", n.id, n.name)
}

func (p Path) String() string {
	return fmt.Sprintf("%v", p.nodes)
}

func (g Graph) String() string {
	var buf bytes.Buffer
	buf.WriteString("\n Nodes \n ")
	for n := range g.nodes {
		buf.WriteString(fmt.Sprintf("%v ", g.nodes[n]))
	}
	buf.WriteString("\n Edges \n ")
	for f := range g.from {
		for t, ok := range g.from[f] {
			if ok {
				buf.WriteString(fmt.Sprintf("%v->%v ", g.nodes[f], g.nodes[t]))
			}
		}
	}
	buf.WriteString("\n Bidirect Edges \n ")
	for f := range g.bidirect {
		for t, ok := range g.bidirect[f] {
			if ok && (f < t) {
				buf.WriteString(fmt.Sprintf("%v<->%v ", g.nodes[f], g.nodes[t]))
			}
		}
	}
	buf.WriteString("\n")
	return buf.String()
}

// NewGraph create a new directed acyclic graph
func NewGraph() *Graph {
	return &Graph{
		nodes:     make(map[int64]Node),
		from:      make(map[int64]map[int64]bool),
		to:        make(map[int64]map[int64]bool),
		bidirect:  make(map[int64]map[int64]bool),
		names:     make(map[string]int64),
		maxNodeID: 0,
	}
}

// CopyGraph copys a graph
func (g *Graph) CopyGraph() *Graph {
	c := NewGraph()
	for k, val := range g.nodes {
		c.nodes[k] = val
	}
	for k, val := range g.from {
		c.from[k] = make(map[int64]bool)
		for k1, val1 := range val {
			c.from[k][k1] = val1
		}
	}
	for k, val := range g.to {
		c.to[k] = make(map[int64]bool)
		for k1, val1 := range val {
			c.to[k][k1] = val1
		}
	}
	for k, val := range g.bidirect {
		c.bidirect[k] = make(map[int64]bool)
		for k1, val1 := range val {
			c.bidirect[k][k1] = val1
		}
	}
	for k, val := range g.names {
		c.names[k] = val
	}
	c.maxNodeID = g.maxNodeID
	return c
}

// AddNode add a new graph node and generate an ID for it
func (g *Graph) AddNode(name string, observable bool) *Node {
	if _, exists := g.names[name]; exists {
		panic(fmt.Sprintf("Find a redeclared node named [%s]", name))
	}
	g.maxNodeID++
	ID := g.maxNodeID
	node := Node{ID, name, observable}
	g.nodes[ID] = node
	g.names[name] = ID
	return &node
}

// AddEdge add a new direct edge to a graph
func (g *Graph) AddEdge(from, to string) {
	fromID, exists := g.names[from]
	if !exists {
		panic(fmt.Sprintf("Node [%s] doesn't exist when creating edge", from))
	}
	toID, exists := g.names[to]
	if !exists {
		panic(fmt.Sprintf("Node [%s] doesn't exist when creating edge", to))
	}
	if fromID == toID {
		panic(fmt.Sprintf("Can't create a loop %s->%s in DAG", from, to))
	}
	if g.from[fromID] == nil {
		g.from[fromID] = make(map[int64]bool)
	}
	g.from[fromID][toID] = true
	if g.to[toID] == nil {
		g.to[toID] = make(map[int64]bool)
	}
	g.to[toID][fromID] = true
}

// AddBidirect add a new bidirect edge to a graph
func (g *Graph) AddBidirect(from, to string) {
	fromID, exists := g.names[from]
	if !exists {
		panic(fmt.Sprintf("Node [%s] doesn't exist when creating edge", from))
	}
	toID, exists := g.names[to]
	if !exists {
		panic(fmt.Sprintf("Node [%s] doesn't exist when creating edge", to))
	}
	if fromID == toID {
		panic(fmt.Sprintf("Can't create a meaning-less bidirect %s->%s", from, to))
	}
	if g.bidirect[fromID] == nil {
		g.bidirect[fromID] = make(map[int64]bool)
	}
	g.bidirect[fromID][toID] = true
	if g.bidirect[toID] == nil {
		g.bidirect[toID] = make(map[int64]bool)
	}
	g.bidirect[toID][fromID] = true
}

// NodeID search a node by its name and returns its ID
func (g *Graph) NodeID(name string) int64 {
	ID, exists := g.names[name]
	if !exists {
		return 0
	}
	return ID
}

// NodeName search a node by its ID and returns its name
func (g *Graph) NodeName(ID int64) string {
	return g.nodes[ID].name
}

// NodeSlice return a slice of observable nodes
func (g *Graph) NodeSlice() []int64 {
	s := make([]int64, 0)
	for k := range g.nodes {
		s = append(s, k)
	}
	return s
}

// HasEdge check is there is an edge from A to B
func (g *Graph) HasEdge(A, B int64) bool {
	e, exists := g.from[A][B]
	return exists && e
}

// To getter
func (g *Graph) To() *map[int64]map[int64]bool {
	return &g.to
}