package graph

import (
	"github.com/L-F-Z/cee/utils"
)

// TopoSort sort the graph by topo order, return ok=false if not DAG
func (g *Graph) TopoSort() (sorted []int64, ok bool) {
	sorted = make([]int64, 0)
	q := utils.NewStack()
	var sum int64 = 0
	inDegree := make(map[int64]int64)
	for n := range g.nodes {
		inDegree[n] = 0
		for _, val := range g.to[n] {
			if val {
				inDegree[n]++
			}
		}
		if inDegree[n] == 0 {
			q.Push(n)
		}
	}
	for q.Len() != 0 {
		u := q.Peek().(int64)
		sorted = append(sorted, u)
		q.Pop()
		sum++
		for v, val := range g.from[u] {
			if val {
				inDegree[v]--
				if inDegree[v] == 0 {
					q.Push(v)
				}
			}
		}
	}
	return sorted, sum == int64(len(g.nodes))
}

// IsDAG judges if the graph is a Directed Acyclic Graph
func (g *Graph) IsDAG() bool {
	_, isDAG := g.TopoSort()
	return isDAG
}

// SubTopo returns a subset of topo sort
func SubTopo(sorted []int64, subset []int64) []int64 {
	newTopo := make([]int64, 0)
	m := make(map[int64]bool)
	for _, val := range subset {
		m[val] = true
	}
	for _, val := range sorted {
		if _, exists := m[val]; exists {
			newTopo = append(newTopo, val)
		}
	}
	return newTopo
}

// TopoAncestor return the ancestors of a node in the topo sort
func TopoAncestor(topo []int64, me int64) []int64 {
	an := make([]int64, 0)
	for _, e := range topo {
		if e == me {
			break
		}
		an = append(an, e)
	}
	return an
}

// TopoAncestorWithMe return the ancestors of a node and itself in the topo sort
func TopoAncestorWithMe(topo []int64, me int64) []int64 {
	an := make([]int64, 0)
	for _, e := range topo {
		an = append(an, e)
		if e == me {
			break
		}
	}
	return an
}
