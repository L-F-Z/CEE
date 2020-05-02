package graph

import ()

// DSeparation judges if x and y are independece on z
func (g *Graph) DSeparation(x []int64, y []int64, z []int64) bool {
	m := make(map[int64]bool)
	for _, e := range z {
		m[e] = true
	}
	for _, nx := range x {
		paths := g.GetAllPaths(nx, y)
		forpath:
		for _, path := range paths{
			g.MarkPath(&path)
			for idx, c := range path.types {
				if c == 'h' || c == 'f' {
					if m[path.nodes[idx]] { // this path is d-separated
						continue forpath
					}
				}
			}
			fornode:
			for idx, c := range path.types {
				if c == 'c' {
					node := path.nodes[idx]
					if m[node] { // this node cannot d-separate the path
						continue fornode
					}
					de := g.GetDescendants(node)
					for _, n := range de {
						if m[n] { // this node cannot d-separate the path
							continue fornode
						}
					}
					// this path is d-separated
					continue forpath
				}
			}
			// this path is not d-separated
			return false
		}
	}
	// all paths are d-separated
	return true
}