package puzzle06

import (
	"bytes"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	block  [][]bool
	x, y   int
	x0, y0 int
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	p.x, p.y = len(lines), len(lines[0])
	p.block = util.SliceND[bool](p.x, p.y).([][]bool)
	for i, line := range lines {
		for j, c := range line {
			p.block[i][j] = c == '#'
			if c == '^' {
				p.x0, p.y0 = i, j
			}
		}
	}
}

func (p *p) solve1() any {
	x, y := p.x0, p.y0
	dir := util.N
	vis := map[[2]int]bool{{x, y}: true}

	for {
		delta := util.Delta4[dir]
		xx, yy := x+delta[0], y+delta[1]
		if xx < 0 || p.x <= xx || yy < 0 || p.y <= yy {
			break
		}

		if p.block[xx][yy] {
			dir = util.Right90[dir]
			continue
		}

		x, y = xx, yy
		vis[[2]int{x, y}] = true
	}

	return vis
}

func (p *p) Solve1() any {
	return len(p.solve1().(map[[2]int]bool))
}

func (p *p) Solve2() any {
	count := 0

	for cord := range p.solve1().(map[[2]int]bool) {
		x1, y1 := cord[0], cord[1]

		x, y := p.x0, p.y0
		dir := util.N
		vis := map[[2]int]util.Direction{{x, y}: util.N}

		for {
			delta := util.Delta4[dir]
			xx, yy := x+delta[0], y+delta[1]
			if xx < 0 || p.x <= xx || yy < 0 || p.y <= yy {
				break
			}

			if p.block[xx][yy] || (xx == x1 && yy == y1) {
				dir = util.Right90[dir]
			} else {
				x, y = xx, yy
			}

			if vis[[2]int{x, y}]&dir != 0 {
				count++
				break
			}

			vis[[2]int{x, y}] |= dir
		}
	}

	return count
}

func init() {
	internal.Register(6, &p{})
}
