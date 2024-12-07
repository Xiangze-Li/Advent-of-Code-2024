package puzzle07

import (
	"bytes"
	"math"
	"strconv"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	targets  []int64
	operands [][]int64
	n        int
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	p.n = len(lines)
	p.targets = make([]int64, 0, p.n)
	p.operands = make([][]int64, 0, p.n)
	for _, line := range lines {
		sp := bytes.SplitN(line, []byte(": "), 2)
		p.targets = append(p.targets, util.Must(strconv.ParseInt(string(sp[0]), 10, 64)))
		p.operands = append(p.operands, util.ArrayBytesToInt64(bytes.Split(sp[1], []byte{' '})))
	}
}

type state struct {
	target int64
	step   int
}

func (p *p) Solve1() any {
	sum := int64(0)

NextEquation:
	for i := range p.n {
		target := p.targets[i]
		operands := p.operands[i]

		q := []state{{target, 0}}

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]

			next := operands[len(operands)-1-cur.step]

			if cur.step == len(operands)-1 {
				if cur.target == next {
					sum += target
					continue NextEquation
				}
				continue
			}

			if cur.target%next == 0 {
				q = append(q, state{cur.target / next, cur.step + 1})
			}

			if cur.target-next > 0 {
				q = append(q, state{cur.target - next, cur.step + 1})
			}
		}
	}

	return sum
}

func ceilingToPowTen(n int64) int64 {
	log10 := math.Log10(float64(n))
	pow := int(math.Ceil(log10))
	if int(log10) == pow {
		// already a power of 10
		pow++
	}
	return int64(math.Pow10(pow))
}

func (p *p) Solve2() any {
	sum := int64(0)

NextEquation:
	for i := range p.n {
		target := p.targets[i]
		operands := p.operands[i]

		q := []state{{target, 0}}

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]

			next := operands[len(operands)-1-cur.step]

			if cur.step == len(operands)-1 {
				if cur.target == next {
					sum += target
					continue NextEquation
				}
				continue
			}

			if cur.target%next == 0 {
				q = append(q, state{cur.target / next, cur.step + 1})
			}

			if cur.target-next > 0 {
				q = append(q, state{cur.target - next, cur.step + 1})
			}

			if pow := ceilingToPowTen(next); cur.target%pow == next {
				q = append(q, state{cur.target / pow, cur.step + 1})
			}
		}
	}

	return sum
}

func init() {
	internal.Register(7, &p{})
}
