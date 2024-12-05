package puzzle05

import (
	"bytes"
	"slices"
	"strconv"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type pair struct {
	a, b int64
}

type p struct {
	rules   map[pair]bool
	updates [][]int64
}

func (p *p) cmp(l, r int64) int {
	if p.rules[pair{l, r}] {
		return -1
	}
	if p.rules[pair{r, l}] {
		return 1
	}
	return 0
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	idx := 0

	p.rules = make(map[pair]bool)
	for len(lines[idx]) != 0 {
		sp := bytes.Split(lines[idx], []byte{'|'})
		a := util.Must(strconv.ParseInt(string(sp[0]), 10, 64))
		b := util.Must(strconv.ParseInt(string(sp[1]), 10, 64))
		p.rules[pair{a, b}] = true
		idx++
	}
	idx++
	for idx < len(lines) {
		sp := bytes.Split(lines[idx], []byte{','})
		p.updates = append(p.updates, util.ArrayBytesToInt64(sp))
		idx++
	}
}

func (p *p) Solve1() any {
	sum := int64(0)

NextUpdate:
	for _, u := range p.updates {
		for i, l := range u {
			for _, r := range u[i+1:] {
				if p.rules[pair{r, l}] {
					continue NextUpdate
				}
			}
		}
		sum += u[len(u)/2]
	}

	return sum
}

func (p *p) Solve2() any {
	sum := int64(0)

NextUpdate:
	for _, u := range p.updates {
		for i, l := range u {
			for _, r := range u[i+1:] {
				if p.rules[pair{r, l}] {
					uu := slices.Clone(u)
					slices.SortFunc(uu, p.cmp)
					sum += uu[len(uu)/2]
					continue NextUpdate
				}
			}
		}
	}

	return sum
}

func init() {
	internal.Register(5, &p{})
}
