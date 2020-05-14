package probability

import (
	"bytes"
	"fmt"
	"github.com/L-F-Z/cee/graph"
)

// Probability type
type Probability struct {
	do       []int64        // do variables
	variable []int64        // conditioned variables
	cond     []int64        // conditioning variables
	sum      bool           // is sum
	sumset   []int64        // variables to be summed over
	product  bool           // is product
	children []*Probability // factors of product
	fraction bool           // is fraction
	num      []*Probability // numerator
	den      []*Probability // denominator
	graph    *graph.Graph   // graph
}

func (p Probability) String() string {
	var buf bytes.Buffer
	g := p.graph
	if p.sum {
		buf.WriteString(" \\sum_{")
		var first bool = true
		for _, n := range p.sumset {
			if first {
				first = false
			} else {
				buf.WriteString(", ")
			}
			buf.WriteString(g.NodeName(n))
		}
		buf.WriteString("}{[")
	}
	if p.product {
		for _, prob := range p.children {
			buf.WriteString(fmt.Sprintf("%v", *prob))
		}
	}
	if p.fraction {
		buf.WriteString(" \\frac{")
		for _, prob := range p.num {
			buf.WriteString(fmt.Sprintf("%v", *prob))
		}
		buf.WriteString("}{")
		for _, prob := range p.den {
			buf.WriteString(fmt.Sprintf("%v", *prob))
		}
		buf.WriteString("}")
	}
	if len(p.variable) != 0 {
		buf.WriteString("P(")
		var first bool = true
		for _, n := range p.variable {
			if first {
				first = false
			} else {
				buf.WriteString(", ")
			}
			buf.WriteString(g.NodeName(n))
		}
	}
	if len(p.cond) != 0 {
		buf.WriteString("|")
		var first bool = true
		for _, n := range p.cond {
			if first {
				first = false
			} else {
				buf.WriteString(", ")
			}
			buf.WriteString(g.NodeName(n))
		}
	}
	if len(p.variable) != 0 {
		buf.WriteString(")")
	}
	if p.sum {
		buf.WriteString("]}")
	}
	return buf.String()
}

// NewProbability create a probability struct
func NewProbability(g *graph.Graph) *Probability {
	var p Probability
	p.do = make([]int64, 0)
	p.variable = make([]int64, 0)
	p.cond = make([]int64, 0)
	p.sum = false
	p.sumset = make([]int64, 0)
	p.product = false
	p.children = make([]*Probability, 0)
	p.fraction = false
	p.num = make([]*Probability, 0)
	p.den = make([]*Probability, 0)
	p.graph = g
	return &p
}

// Do getter
func (p *Probability) Do() []int64 {
	return p.do
}

// AddDo setter
func (p *Probability) AddDo(n int64) {
	if !contains(p.do, n) {
		p.do = append(p.do, n)
	}
}

// AddDoSlice setter
func (p *Probability) AddDoSlice(n []int64) {
	m := make(map[int64]bool)
	for _, val := range p.do {
		m[val] = true
	}
	for _, val := range n {
		if _, exists := m[val]; !exists {
			p.do = append(p.do, val)
		}
	}
}

// SetDo setter
func (p *Probability) SetDo(n []int64) {
	p.do = nil
	p.do = make([]int64, len(n))
	copy(p.do, n)
}

// Variable getter
func (p *Probability) Variable() []int64 {
	return p.variable
}

// AddVariable setter
func (p *Probability) AddVariable(n int64) {
	if !contains(p.variable, n) {
		p.variable = append(p.variable, n)
	}
}

// AddVariableSlice setter
func (p *Probability) AddVariableSlice(n []int64) {
	m := make(map[int64]bool)
	for _, val := range p.variable {
		m[val] = true
	}
	for _, val := range n {
		if _, exists := m[val]; !exists {
			p.variable = append(p.variable, val)
		}
	}
}

// SetVariable setter
func (p *Probability) SetVariable(n []int64) {
	p.variable = nil
	p.variable = make([]int64, len(n))
	copy(p.variable, n)
}

// Cond getter
func (p *Probability) Cond() []int64 {
	return p.cond
}

// AddCond setter
func (p *Probability) AddCond(n int64) {
	if !contains(p.cond, n) {
		p.cond = append(p.cond, n)
	}
}

// AddCondSlice setter
func (p *Probability) AddCondSlice(n []int64) {
	m := make(map[int64]bool)
	for _, val := range p.cond {
		m[val] = true
	}
	for _, val := range n {
		if _, exists := m[val]; !exists {
			p.cond = append(p.cond, val)
		}
	}
}

// SetCond setter
func (p *Probability) SetCond(n []int64) {
	p.cond = nil
	p.cond = make([]int64, len(n))
	copy(p.cond, n)
}

// Sum getter
func (p *Probability) Sum() bool {
	return p.sum
}

// SetSum setter
func (p *Probability) SetSum(b bool) {
	p.sum = b
}

// Sumset getter
func (p *Probability) Sumset() []int64 {
	return p.sumset
}

// AddSumset setter
func (p *Probability) AddSumset(n int64) {
	if !contains(p.sumset, n) {
		p.sumset = append(p.sumset, n)
	}
	if len(p.sumset) > 0 {
		p.sum = true
	}
}

// AddSumsetSlice setter
func (p *Probability) AddSumsetSlice(n []int64) {
	m := make(map[int64]bool)
	for _, val := range p.sumset {
		m[val] = true
	}
	for _, val := range n {
		if _, exists := m[val]; !exists {
			p.sumset = append(p.sumset, val)
		}
	}
	if len(p.sumset) > 0 {
		p.sum = true
	}
}

// SetSumset setter
func (p *Probability) SetSumset(n []int64) {
	p.sumset = nil
	p.sumset = make([]int64, len(n))
	copy(p.sumset, n)
	if len(p.sumset) > 0 {
		p.sum = true
	}
}

// Product getter
func (p *Probability) Product() bool {
	return p.product
}

// SetProduct setter
func (p *Probability) SetProduct(b bool) {
	p.product = b
}

// AddChildren setter
func (p *Probability) AddChildren(prob *Probability) {
	p.children = append(p.children, prob)
	if len(p.children) > 0 {
		p.product = true
	}
}

// Fraction getter
func (p *Probability) Fraction() bool {
	return p.fraction
}

// SetFraction setter
func (p *Probability) SetFraction(b bool) {
	p.fraction = b
}

// AddNum setter
func (p *Probability) AddNum(prob *Probability) {
	p.num = append(p.num, prob)
	if len(p.num) > 0 {
		p.fraction = true
	}
}

// AddDen setter
func (p *Probability) AddDen(prob *Probability) {
	p.den = append(p.den, prob)
	if len(p.den) > 0 {
		p.fraction = true
	}
}

// Graph getter
func (p *Probability) Graph() *graph.Graph {
	return p.graph
}

// contains check is an element is contained in a slice
func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
