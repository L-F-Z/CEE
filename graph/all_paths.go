package graph

import (
	"github.com/L-F-Z/cee/utils"
	"fmt"
)

// nodeSet defines a node set with pointer
type nodeSet struct {
	nodes   []int64
	pointer int64
}

func (n nodeSet) String() string {
	return fmt.Sprintf("[%v]%v", n.pointer, n.nodes)
}

func neighbours(g *Graph, n int64) nodeSet {
	var result nodeSet
	result.pointer = 0
	result.nodes = g.GetNeighbours(n)
	return result
}

// GetAllPaths function generate all paths from one source node to multiple destinations
func (g *Graph) GetAllPaths(src int64, dst []int64) []Path {
	result := make([]Path, 0)
	visited := make([]int64, 0)
	visited = append(visited, src)
	steps := utils.NewStack()
	steps.Push(neighbours(g, src))
	for steps.Len() != 0 {
		children := steps.Pop().(nodeSet)
		if children.pointer == int64(len(children.nodes)) {
			visited = visited[:len(visited)-1] // delete last element
			continue
		}
		child := children.nodes[children.pointer]
		children.pointer++
		steps.Push(children)
		if contains(visited, child) {
			continue
		}
		if contains(dst, child) {
			var path Path
			path.from = src
			path.to = child
			path.nodes = make([]int64, len(visited))
			copy(path.nodes, visited)
			path.nodes = append(path.nodes, child)
			result = append(result, path)
		}
		visited = append(visited, child)
		if !allTargets(visited, dst) {
			steps.Push(neighbours(g, child))
		} else {
			visited = visited[:len(visited)-1] // delete last element
		}
	}
	return result
}

// contains check is an element is contained in a slice
func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// allTargets check is all dst elements have been contained in visited
func allTargets(visited, dst []int64) bool {
	for _, d := range dst {
		if !contains(visited, d) {
			return false
		}
	}
	return true
}

/*
Inspired by
R. Sedgewick, "Algorithms in C, Part 5: Graph Algorithms",
Addison Wesley Professional, 3rd ed., 2001.
*/
