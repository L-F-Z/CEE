package identify

import (
	"fmt"
	"github.com/L-F-Z/cee/graph"
	"github.com/L-F-Z/cee/probability"
	"github.com/L-F-Z/cee/utils"
)

func rc(d []int64, v []int64, s int64, p *probability.Probability, g *graph.Graph) *probability.Probability {
	fmt.Println("#[rc] d = ", d, "v = ", v)
	// line 1
	AnD := g.An(d)
	AnS := g.An([]int64{s})
	tmp := utils.SetUnion(AnD, AnS)
	reduced := utils.SetMinus(v, tmp)
	if len(reduced) != 0 {
		fmt.Println("[rc] line 1")
		newP := p.CopyProbability()
		newP.AddSumsetSlice(reduced)
		newG := g.InducedSubgraph(tmp)
		newV := newG.NodeSlice()
		return rc(d, newV, s, newP, newG)
	}
	// line 2
	ccom := g.CCompenent()
	c := utils.DeleteSliceBySet(ccom, AnS)
	fmt.Println("#[rc] c = ", c)
	// line 3
	if len(c) == 0 {
		fmt.Println("[rc] line 3")
		panic(fmt.Sprintf("Can't Recover. %v", g))
	}
	// line 4
	ci := utils.SliceContainSlice(c, d)
	if len(ci) != 0 {
		fmt.Println("[rc] line 4")
		q := p.Q(v, ci)
		return ident(d, ci, q, g)
	}
	// line 5
	fmt.Println("[rc] line 5")
	den := probability.NewProbability(g)
	for _, cc := range c {
		child := p.Q(v, cc)
		den.AddChildren(child)
	}
	newP := probability.NewProbability(g)
	newP.AddNum(p)
	newP.AddDen(den)
	tmp = utils.CombineSlice(c)
	tmp = utils.SetMinus(v, tmp)
	newG := g.InducedSubgraph(tmp)
	newV := newG.NodeSlice()
	return rc(d, newV, s, newP, newG)
}
