package puzzle25

import (
	"bytes"
	"strings"

	"github.com/Xiangze-Li/advent-2024/internal"
)

type p struct {
	keys  [][5]int
	locks [][5]int
}

//nolint:gocognit // .
func (p *p) Init(data []byte) {
	sp := bytes.Split(data, []byte("\n\n"))
	for _, s := range sp {
		height := [5]int{}
		lines := strings.Split(string(s), "\n")
		if lines[0] == "....." {
			for j := range 5 {
				for i := 1; i <= 6; i++ {
					if lines[i][j] == '#' {
						height[j] = 6 - i
						break
					}
				}
			}
			p.keys = append(p.keys, height)
		} else {
			for j := range 5 {
				for i := 5; i >= 0; i-- {
					if lines[i][j] == '#' {
						height[j] = i
						break
					}
				}
			}
			p.locks = append(p.locks, height)
		}
	}
}

func (p *p) Solve1() any {
	count := 0
	for _, k := range p.keys {
		for _, l := range p.locks {
			ok := true
			for x := range 5 {
				if k[x]+l[x] > 5 {
					ok = false
					break
				}
			}
			if ok {
				count++
			}
		}
	}
	return count
}

func (p *p) Solve2() any {
	return "Happy Holidays!"
}

func init() {
	internal.Register(25, &p{})
}
