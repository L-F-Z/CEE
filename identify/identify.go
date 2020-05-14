package identify

import (
	"github.com/L-F-Z/cee/graph"
	"github.com/L-F-Z/cee/probability"
)

// Identify main algo
func Identify(y, x, z []int64, g *graph.Graph) *probability.Probability {
	topo, _ := g.TopoSort()
	v := g.NodeSlice()
	p := probability.GenProbability(g)
	if s := g.NodeID("$S$"); s != 0 { // need to recover causal effects from selection bias
		return recover(y, x, v, s, p, g)
	}
	if len(z) == 0 {
		return id(y, x, v, p, g, topo)
	}
	return idc(y, x, z, v, p, g, topo)
}
