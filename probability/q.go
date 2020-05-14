package probability

import (
	"github.com/L-F-Z/cee/utils"
)

// Q following the notation in Tian and Pearl 2002
func (p *Probability) Q(v, c []int64) *Probability {
	r := p.CopyProbability()
	add := utils.SetMinus(v, c)
	r.AddSumsetSlice(add)
	return r
}
