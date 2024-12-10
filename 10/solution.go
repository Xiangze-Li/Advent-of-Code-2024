package puzzle10

import (
	"bytes"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	grids [][]byte
	x, y  int
}

func (p *p) Init(data []byte) {
	// 	data = []byte(`89010123
	// 78121874
	// 87430965
	// 96549874
	// 45678903
	// 32019012
	// 01329801
	// 10456732`)

	p.grids = bytes.Split(data, []byte("\n"))
	p.x, p.y = len(p.grids), len(p.grids[0])
}

func (p *p) bfs(i, j int, sum *int, skipVis bool) {
	queue := [][2]int{{i, j}}
	vis := map[[2]int]bool{}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if vis[cur] && skipVis {
			continue
		}
		vis[cur] = true

		curH := p.grids[cur[0]][cur[1]]
		if curH == '0' {
			*sum += 1
			continue
		}
		for _, delta := range util.Delta4 {
			ii, jj := cur[0]+delta[0], cur[1]+delta[1]
			if 0 <= ii && ii < p.x && 0 <= jj && jj < p.y &&
				p.grids[ii][jj] == curH-1 {
				queue = append(queue, [2]int{ii, jj})
			}
		}
	}
}

func (p *p) Solve1() any {
	sum := 0

	for i := range p.x {
		for j := range p.y {
			if p.grids[i][j] != '9' {
				continue
			}
			p.bfs(i, j, &sum, true)
		}
	}

	return sum
}

func (p *p) Solve2() any {
	sum := 0

	for i := range p.x {
		for j := range p.y {
			if p.grids[i][j] != '9' {
				continue
			}
			p.bfs(i, j, &sum, false)
		}
	}

	return sum
}

func init() {
	internal.Register(10, &p{})
}
