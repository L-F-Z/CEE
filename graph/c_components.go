package graph

import (
	"github.com/L-F-Z/cee/utils"
)

// CCompenent returns all c-compenents of a graph
func (g *Graph) CCompenent() [][]int64 {
	result := make([][]int64, 0)
	unexplored := make(map[int64]bool, 0)
	for k := range g.nodes {
		unexplored[k] = true
	}
	s := utils.NewStack()
	for k := range unexplored {
		delete(unexplored, k)
		comp := make([]int64, 0)
		comp = append(comp, k)
		s.Push(k)
		for s.Len() != 0 {
			from := s.Pop().(int64)
			for to := range unexplored {
				if val, exists := g.bidirect[from][to]; exists && val {
					comp = append(comp, to)
					delete(unexplored, to)
					s.Push(to)
				}
			}
		}
		result = append(result, comp)
	}
	return result
}
