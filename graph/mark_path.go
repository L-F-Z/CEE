package graph

import ()

// MarkPath mark the path with e,c,f,h
func (g *Graph) MarkPath(p *Path) {
	nodes := p.nodes
	len := len(nodes)
	p.types = make([]byte, len)
	p.types[0] = 'e'
	p.types[len-1] = 'e'
	// fork - f   chain - h   collider - c   end - e
	last := g.from[nodes[0]][nodes[1]]
	// if n1->n2--n3 last = true
	// if n1<-n2--n3 last = false
	for i := 1; i < len-1; i++ {
		current := g.from[nodes[i]][nodes[i+1]]
		if last && !current { // n1->n2<-n3
			p.types[i] = 'c'
		} else if !last && current { // n1<-n2->n3
			p.types[i] = 'f'
		} else {
			p.types[i] = 'h'
		}
		last = current
	}
}