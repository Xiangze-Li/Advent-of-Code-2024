package puzzle11

import (
	"bytes"
	"maps"
	"math"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	stones map[int64]uint64
}

func (p *p) Init(data []byte) {
	v := util.ArrayBytesToInt64(bytes.Split(data, []byte{' '}))
	p.stones = make(map[int64]uint64, len(v))
	for _, s := range v {
		p.stones[s]++
	}
}

func step(src map[int64]uint64) map[int64]uint64 {
	dst := make(map[int64]uint64, 2*len(src))
	for v, c := range src {
		d := util.CountDigits(v)
		switch {
		case v == 0:
			dst[1] += c
		case d%2 == 0:
			base := int64(math.Pow10(d / 2))
			dst[v/base] += c
			dst[v%base] += c
		default:
			dst[v*2024] += c
		}
	}
	return dst
}

func sum(s map[int64]uint64) uint64 {
	var acc uint64
	for _, v := range s {
		acc += v
	}
	return acc
}

func (p *p) Solve1() any {
	stones := maps.Clone(p.stones)
	for range 25 {
		stones = step(stones)
	}
	return sum(stones)
}

func (p *p) Solve2() any {
	stones := maps.Clone(p.stones)
	for range 75 {
		stones = step(stones)
	}
	return sum(stones)
}

func init() {
	internal.Register(11, &p{})
}
