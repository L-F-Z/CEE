package identify

import (
	"github.com/L-F-Z/cee/graph"
	"github.com/L-F-Z/cee/probability"
	"github.com/L-F-Z/cee/utils"
	"fmt"
)

func id(y []int64, x []int64, v []int64, p *probability.Probability, g *graph.Graph, topo []int64) *probability.Probability{
	fmt.Println("ID y=", y, " x=", x, " v=", v, " p=", p)
	// line 1
	if len(x) == 0 {
		fmt.Println("Line 1 ")
		if p.Fraction() || p.Product() {
			tmp := utils.SetMinus(v, y)
			p.AddSumsetSlice(tmp)
		} else {
			p.SetVariable(y)
		}
		return p
	}
	// line 2
	an := g.An(y)
	if !utils.SameSet(v, an) {
		fmt.Println("Line 2")
		if p.Fraction() || p.Product() {
			tmp := utils.SetMinus(v, an)
			p.AddSumsetSlice(tmp)
		} else {
			p.SetVariable(an)
		}
		newG := g.InducedSubgraph(an)
		newV := newG.NodeSlice()
		newT := graph.SubTopo(topo, newV)
		newX := utils.SetIntersect(x, an)
		return id(y, newX, newV, p, newG, newT)
	}
	// line 3
	Gx := g.DoGraph(x)
	an = Gx.An(y)
	w := utils.SetMinus(v, x)
	w = utils.SetMinus(w, an)
	if len(w) != 0 {
		fmt.Println("Line 3")
		newX := utils.SetUnion(x, w)
		return id(y, newX, v, p, g, topo)
	}
	// line 4
	tmp := utils.SetMinus(v, x)
	newG := g.InducedSubgraph(tmp)
	c := newG.CCompenent()
	if len(c) > 1 {
		fmt.Println("Line 4")
		r := probability.NewProbability(g)
		tmp := utils.SetUnion(x, y)
		tmp = utils.SetMinus(v, tmp)
		r.SetSumset(tmp)
		for _, si := range c {
			newX := utils.SetMinus(v, si)
			newP := p.CopyProbability()
			prod := id(si, newX, v, newP, g, topo)
			r.AddChildren(prod)
		}
		return r
	}
	// line 5
	s := c[0]
	c = g.CCompenent()
	if utils.SameSet(v, c[0]) {
		panic(fmt.Sprintf("Can't Identify. %v \n %v", g, s))
	}
	// line 6
	if utils.SliceInSlice(c, s) {
		fmt.Println("Line 6 ")
		list := make([]*probability.Probability, 0)
		var pChild *probability.Probability
		for _, val := range s {
			cond := graph.TopoAncestor(topo, val)
			if p.Product() {
				pNum := p.CopyProbability()
				tmp := utils.SetMinus(v, graph.TopoAncestorWithMe(topo, val))
				pNum.AddSumsetSlice(tmp)
				if len(cond) > 0 {
					pDen := p.CopyProbability()
					tmp := utils.SetMinus(v, cond)
					pDen.AddSumsetSlice(tmp)
					pChild = probability.NewProbability(p.Graph())
					pChild.AddNum(pNum)
					pChild.AddDen(pDen)
				} else {
					pChild = pNum.CopyProbability()
				}
			} else {
				pChild = p.CopyProbability()
				pChild.SetVariable([]int64{val})
				pChild.SetCond(cond)
			}
			list = append(list, pChild)
		}
		r := probability.NewProbability(p.Graph())
		if len(s) > 1{
			tmp := utils.SetMinus(s, y)
			r.SetSumset(tmp)
			for _, e := range list {
				r.AddChildren(e)
			}
			return r
		}
		if (pChild.Product() || pChild.Fraction()) {
			tmp := utils.SetMinus(s, y)
			pChild.AddSumsetSlice(tmp)
		} else {
			tmp := utils.SetMinus(s, y)
			tmp = utils.SetUnion(pChild.Sumset(), tmp)
			tmp = utils.SetMinus(pChild.Variable(), tmp)
			pChild.SetVariable(tmp)
		}
		r = pChild.CopyProbability()
		return r
	}
	// line 7
	fmt.Println("Line 7")
	s = utils.SliceContainSlice(c, s)
	list := make([]*probability.Probability, 0)
	var pChild *probability.Probability
	for _, val := range s {
		cond := graph.TopoAncestor(topo, val)
		if p.Product() {
			pNum := p.CopyProbability()
			tmp := utils.SetMinus(v, graph.TopoAncestorWithMe(topo, val))
			pNum.AddSumsetSlice(tmp)
			if len(cond) > 0 {
				pDen := p.CopyProbability()
				tmp := utils.SetMinus(v, cond)
				pDen.AddSumsetSlice(tmp)
				pChild = probability.NewProbability(p.Graph())
				pChild.AddNum(pNum)
				pChild.AddDen(pDen)
			} else {
				pChild = pNum.CopyProbability()
			}
		} else {
			pChild = p.CopyProbability()
			pChild.SetVariable([]int64{val})
			pChild.SetCond(cond)
		}
		list = append(list, pChild)
	}
	newG = g.InducedSubgraph(s)
	newV := newG.NodeSlice()
	newT := graph.SubTopo(topo, newV)
	newX := utils.SetIntersect(x, s)
	var newP *probability.Probability
	if len(s) == 1 {
		newP = pChild.CopyProbability()
	} else {
		newP = probability.NewProbability(p.Graph())
		for _, e := range list {
			newP.AddChildren(e)
		}
	}
	return id(y, newX, newV, newP, newG, newT)
}