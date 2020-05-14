package identify

import (
	"fmt"
	"github.com/L-F-Z/cee/graph"
	"github.com/L-F-Z/cee/probability"
	"github.com/L-F-Z/cee/utils"
)

func ident(c, t []int64, q *probability.Probability, g *graph.Graph) *probability.Probability {
	fmt.Println("[ident]", c, t)
	gt := g.InducedSubgraph(t)
	a := gt.An(c)
	a = utils.SetIntersect(a, t)
	// line i)
	if utils.SameSet(a, c) {
		fmt.Println("[ident] i)")
		sumset := utils.SetMinus(t, c)
		r := q.CopyProbability()
		r.AddSumsetSlice(sumset)
		return r
	}
	// line ii)
	if utils.SameSet(a, t) {
		fmt.Println("[ident] ii)")
		panic(fmt.Sprintf("Can't Ident. %v", g))
	}
	// line iii)
	fmt.Println("[ident] iii)")
	ga := g.InducedSubgraph(a)
	cc := ga.CCompenent()
	t1 := utils.SliceContainSlice(cc, c)
	newT := utils.SetIntersect(t1, a)
	fmt.Println("#[ident] cc = ", cc)
	fmt.Println("#[ident] t1 = ", t1)
	fmt.Println("#[ident] newT = ", newT)
	qa := q.Q(t, a)
	newQ := qa.Q(a, t1)
	return ident(c, newT, newQ, g)
}
