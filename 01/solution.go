package puzzle01

import (
	"bytes"
	"slices"
	"strconv"
	"strings"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	l, r []uint64
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	p.l = make([]uint64, 0, len(lines))
	p.r = make([]uint64, 0, len(lines))
	for _, line := range lines {
		sp := strings.Fields(string(line))
		p.l = append(p.l, util.Must(strconv.ParseUint(sp[0], 10, 64)))
		p.r = append(p.r, util.Must(strconv.ParseUint(sp[1], 10, 64)))
	}
	slices.Sort(p.l)
	slices.Sort(p.r)
}

func (p *p) Solve1() any {
	diff := uint64(0)
	for i, vl := range p.l {
		vr := p.r[i]
		if vl > vr {
			diff += vl - vr
		} else {
			diff += vr - vl
		}
	}
	return diff
}

func (p *p) Solve2() any {
	score := uint64(0)
	for i := 0; i < len(p.l); {
		v := p.l[i]
		prevI := i
		i, _ = slices.BinarySearch(p.l, v+1)
		rLB, found := slices.BinarySearch(p.r, v)
		if !found {
			continue
		}
		rUB, _ := slices.BinarySearch(p.r, v+1)
		score += uint64(rUB-rLB) * v * uint64(i-prevI)
	}
	return score
}

func init() {
	internal.Register(1, &p{})
}
