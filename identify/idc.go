package identify

import (
	"github.com/L-F-Z/cee/graph"
	"github.com/L-F-Z/cee/probability"
	"github.com/L-F-Z/cee/utils"
	"fmt"
)

func idc(y, x, z, v []int64, p *probability.Probability, g *graph.Graph, topo []int64) *probability.Probability{
	fmt.Println("IDC y=", y, " x=", x, " z=", z, " v=", v, " p=", p)
	// line 1
	subg := g.Subgraph(x, z)
	for _, nz := range z {
		sz := []int64{nz}
		tmp := utils.SetMinus(z, sz)
		tmp = utils.SetUnion(x, tmp)
		if subg.DSeparation(y, sz, tmp) {
			newX := utils.SetUnion(x, sz)
			newZ := utils.SetMinus(z, sz)
			return idc(y, newX, newZ, v, p, g, topo)
		}
	}
	newY := utils.SetUnion(y, z)
	return id(newY, x, v, p, g, topo)
}