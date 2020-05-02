package graph

import ()

// DoGraph opeartes "do" on a set of nodes
func (g *Graph) DoGraph(do []int64) *Graph {
	r := g.CopyGraph()
	for _, n := range do {
		for k := range r.to[n] {
			delete(r.to[n], k)
			delete(r.from[k], n)
			if len(r.from[k]) == 0 {
				delete(r.from, k)
			}
		}
		delete(r.to, n)
	}
	for _, n := range do {
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
