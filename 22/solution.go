package puzzle22

import (
	"bytes"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	n []uint64
}

func (p *p) Init(data []byte) {
	p.n = util.ArrayBytesToUint64(bytes.Split(data, []byte{'\n'}))
}

func step(n uint64) uint64 {
	const module = 16777216
	n ^= n << 6
	n %= module
	n ^= n >> 5
	n %= module
	n ^= n << 11
	n %= module
	return n
}

func (p *p) Solve1() any {
	acc := uint64(0)
	for _, n := range p.n {
		for range 2000 {
			n = step(n)
		}
		acc += n
	}
	return acc
}

func (p *p) Solve2() any {
	prices := make(map[[4]int]int, 2000*len(p.n))

	for _, n := range p.n {
		vis := map[[4]int]bool{}
		k := [4]int{}
		prevCount := int(n) % 10
		for x := range 3 {
			n = step(n)
			c := int(n) % 10
			k[x+1] = c - prevCount
			prevCount = c
		}

		for x := 3; x < 2000; x++ {
			n = step(n)
			c := int(n) % 10
			k[0], k[1], k[2], k[3] = k[1], k[2], k[3], c-prevCount
			prevCount = c
			if vis[k] {
				continue
			}
			vis[k] = true
			prices[k] += c
		}
	}

	maxi := 0
	for _, v := range prices {
		maxi = max(maxi, v)
	}
	return maxi
}

func init() {
	internal.Register(22, &p{})
}
