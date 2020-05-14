package graph

import (
	"github.com/L-F-Z/cee/utils"
)

// GetNeighbours returns all connected nodes of n
func (g *Graph) GetNeighbours(n int64) []int64 {
	connected := make([]int64, 0)
	for k, exists := range g.from[n] {
		if exists {
			connected = append(connected, k)
		}
	}
	for k, exists := range g.to[n] {
		if exists {
			connected = append(connected, k)
		}
	}
	return connected
}

// IsNeighbour checks if node n is a neighbour of node me
func (g *Graph) IsNeighbour(n, me int64) bool {
	if n == me {
		return false
	}
	for k, exists := range g.from[me] {
		if n == k && exists {
			return true
		}
	}
	for k, exists := range g.to[me] {
		if n == k && exists {
			return true
		}
	}
	return false
}

// GetDescendants returns all descendant nodes of n
func (g *Graph) GetDescendants(n int64) []int64 {
	descendants := make([]int64, 0)
	s := utils.NewStack()
	s.Push(n)
	for s.Len() != 0 {
		cur := s.Pop().(int64)
		for k, exists := range g.from[cur] {
			if exists {
				descendants = append(descendants, k)
				s.Push(k)
			}
		}
	}
	return utils.SetUnique(descendants)
}

// IsDescendant checks if node d is a descendant of node me
func (g *Graph) IsDescendant(d, me int64) bool {
	if d == me {
		return false
	}
	s := utils.NewStack()
	s.Push(me)
	for s.Len() != 0 {
		cur := s.Pop().(int64)
		for k, exists := range g.from[cur] {
			if exists {
				if d == k {
					return true
				}
				s.Push(k)
			}
		}
	}
	return false
}

// GetAncestors returns all ancestor nodes of n
func (g *Graph) GetAncestors(n int64) []int64 {
	ancestors := make([]int64, 0)
	s := utils.NewStack()
	s.Push(n)
	for s.Len() != 0 {
		cur := s.Pop().(int64)
		for k, exists := range g.to[cur] {
			if exists {
				ancestors = append(ancestors, k)
				s.Push(k)
			}
		}
	}
	return utils.SetUnique(ancestors)
}

// IsAncestor checks if node a is an ancestor of node me
func (g *Graph) IsAncestor(a, me int64) bool {
	if a == me {
		return false
	}
	s := utils.NewStack()
	s.Push(me)
	for s.Len() != 0 {
		cur := s.Pop().(int64)
		for k, exists := range g.to[cur] {
			if exists {
				if a == k {
					return true
				}
				s.Push(k)
			}

		}
	}
	return false
}

// GetParents returns all parent nodes of n
func (g *Graph) GetParents(n int64) []int64 {
	parents := make([]int64, 0)
	for k, exists := range g.to[n] {
		if exists {
			parents = append(parents, k)
		}
	}
	return parents
}

// IsParent checks if node p is a parent of node me
func (g *Graph) IsParent(p, me int64) bool {
	if p == me {
		return false
	}
	for k, exists := range g.to[me] {
		if p == k && exists {
			return true
		}
	}
	return false
}

// GetChildren returns all child nodes of n
func (g *Graph) GetChildren(n int64) []int64 {
	children := make([]int64, 0)
	for k, exists := range g.from[n] {
		if exists {
			children = append(children, k)
		}
	}
	return children
}

// IsChild checks if node c is a child of node me
func (g *Graph) IsChild(c, me int64) bool {
	if c == me {
		return false
	}
	for k, exists := range g.from[me] {
		if c == k && exists {
			return true
		}
	}
	return false
}

// An returns a set of n ancesters and n
func (g *Graph) An(n []int64) []int64 {
	ancestors := make([]int64, 0)
	s := utils.NewStack()
	for _, me := range n {
		ancestors = append(ancestors, me)
		s.Push(me)
		for s.Len() != 0 {
			cur := s.Pop().(int64)
			for k, exists := range g.to[cur] {
				if exists {
					ancestors = append(ancestors, k)
					s.Push(k)
				}
			}
		}
	}
	return utils.SetUnique(ancestors)
}
