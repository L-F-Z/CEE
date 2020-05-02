package probability

import (
	"github.com/L-F-Z/cee/graph"
)

// GenProbability generate a probability struct based on the graph
func GenProbability(g *graph.Graph) *Probability {
	r := NewProbability(g)
	r.SetProduct(true)
	to := *g.To()
	topo, _ := g.TopoSort()
	for _, n := range topo {
		p := NewProbability(g)
		p.AddVariable(n)
		for k := range to[n] {
			p.AddCond(k)
		}
		r.AddChildren(p)
	}
	return r
}
