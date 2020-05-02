package graph

import ()

// Subgraph generated a subgraph by the deletion of edges
func (g *Graph) Subgraph(InDel, OutDel []int64) *Graph {
	r := g.CopyGraph()
	for _, n := range InDel {
		for k := range r.to[n] {
			delete(r.to[n], k)
			delete(r.from[k], n)
			if len(r.from[k]) == 0 {
				delete(r.from, k)
			}
		}
		delete(r.to, n)
	}
	for _, n := range InDel {
		for k := range r.bidirect[n] {
			delete(r.bidirect[n], k)
			delete(r.bidirect[k], n)
			if len(r.bidirect[k]) == 0 {
				delete(r.bidirect, k)
			}
		}
		delete(r.bidirect, n)
	}
	for _, n := range OutDel {
		for k := range r.from[n] {
			delete(r.from[n], k)
			delete(r.to[k], n)
			if len(r.to[k]) == 0 {
				delete(r.to, k)
			}
		}
		delete(r.from, n)
	}
	for _, n := range OutDel {
		for k := range r.bidirect[n] {
			delete(r.bidirect[n], k)
			delete(r.bidirect[k], n)
			if len(r.bidirect[k]) == 0 {
				delete(r.bidirect, k)
			}
		}
		delete(r.bidirect, n)
	}
	return r
}

// InducedSubgraph generates a induced subgraph of g by the given subset of nodes
func (g *Graph) InducedSubgraph(n []int64) *Graph {
	r := g.CopyGraph()
	nodes := make(map[int64]bool)
	for _, val := range n {
		nodes[val] = true
	}
	for k, val := range r.names {
		if _, exists := nodes[val]; !exists {
			delete(r.names, k)
		}
	}
	for k := range r.nodes {
		if _, exists := nodes[k]; !exists {
			delete(r.nodes, k)
		}
	}
	for k, val := range r.from {
		if _, exists := nodes[k]; !exists {
			delete(r.from, k)
		} else {
			for k1 := range val {
				if _, exists := nodes[k1]; !exists {
					delete(val, k1)
				}
			}
			if len(val) == 0 {
				delete(r.from, k)
			}
		}
	}
	for k, val := range r.to {
		if _, exists := nodes[k]; !exists {
			delete(r.to, k)
		} else {
			for k1 := range val {
				if _, exists := nodes[k1]; !exists {
					delete(val, k1)
				}
			}
			if len(val) == 0 {
				delete(r.to, k)
			}
		}
	}
	for k, val := range r.bidirect {
		if _, exists := nodes[k]; !exists {
			delete(r.bidirect, k)
		} else {
			for k1 := range val {
				if _, exists := nodes[k1]; !exists {
					delete(val, k1)
				}
			}
			if len(val) == 0 {
				delete(r.bidirect, k)
			}
		}
	}
	return r
}
