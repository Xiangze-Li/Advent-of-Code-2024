package puzzle08

import (
	"bytes"

	"github.com/Xiangze-Li/advent-2024/internal"
)

type p struct {
	anntennas map[byte][][2]int
	x, y      int
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	p.x, p.y = len(lines), len(lines[0])
	p.anntennas = map[byte][][2]int{}
	for i := range lines {
		for j, c := range lines[i] {
			if c != '.' {
				p.anntennas[c] = append(p.anntennas[c], [2]int{i, j})
			}
		}
	}
}

func (p *p) Solve1() any {
	vis := map[[2]int]bool{}

	for _, kind := range p.anntennas {
		for i, cordI := range kind {
			for j, cordJ := range kind {
				if i == j {
					continue
				}
				x, y := 2*cordI[0]-cordJ[0], 2*cordI[1]-cordJ[1]
				if 0 <= x && x < p.x && 0 <= y && y < p.y {
					vis[[2]int{x, y}] = true
				}
			}
		}
	}

	return len(vis)
}

func (p *p) Solve2() any {
	vis := map[[2]int]bool{}

	for _, kind := range p.anntennas {
		for i, cordI := range kind {
			for j, cordJ := range kind {
				if i == j {
					continue
				}
				dx, dy := cordI[0]-cordJ[0], cordI[1]-cordJ[1]
				x, y := cordI[0], cordI[1]
				for 0 <= x && x < p.x && 0 <= y && y < p.y {
					vis[[2]int{x, y}] = true
					x += dx
					y += dy
				}
			}
		}
	}

	return len(vis)
}

func init() {
	internal.Register(8, &p{})
}
