package probability

import ()

// CopyProbability copy a probability struct
func (p *Probability) CopyProbability() *Probability {
	var r Probability
	r.do = make([]int64, len(p.do))
	copy(r.do, p.do)
	r.variable = make([]int64, len(p.variable))
	copy(r.variable, p.variable)
	r.cond = make([]int64, len(p.cond))
	copy(r.cond, p.cond)
	r.sum = p.sum
	r.sumset = make([]int64, len(p.sumset))
	copy(r.sumset, p.sumset)
	r.product = p.product
	r.children = make([]*Probability, len(p.children))
	copy(r.children, p.children)
	r.fraction = p.fraction
	r.num = make([]*Probability, len(p.num))
	copy(r.num, p.num)
	r.den = make([]*Probability, len(p.den))
	copy(r.den, p.den)
	r.graph = p.graph
	return &r
}
