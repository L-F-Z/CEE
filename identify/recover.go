package identify

import (
	"fmt"
	"github.com/L-F-Z/cee/graph"
	"github.com/L-F-Z/cee/probability"
	"github.com/L-F-Z/cee/utils"
)

func recover(y []int64, x []int64, v []int64, s int64, p *probability.Probability, g *graph.Graph) *probability.Probability {
	tmp := utils.SetMinus(v, x)
	tmpG := g.InducedSubgraph(tmp)
	fmt.Println("##", tmpG)
	D := tmpG.An(y)
	tmpD := g.InducedSubgraph(D)
	Dset := tmpD.CCompenent()
	sum := utils.SetMinus(D, y)
	fmt.Println("Dset", Dset)
	if len(Dset) == 1 { // Dset has at least one element, because An(y) is not empty
		if len(sum) == 0 {
			return rc(Dset[0], v, s, p, g)
		}
		r := probability.NewProbability(g)
		r.SetSumset(sum)
		r.AddChildren(rc(Dset[0], v, s, p, g))
		return r
	}
	r := probability.NewProbability(g)
	if len(sum) != 0 {
		r.SetSumset(sum)
	}
	for _, Di := range Dset {
		r.AddChildren(rc(Di, v, s, p, g))
	}
	return r
}
