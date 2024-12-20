package puzzle20

import (
	"bytes"
	"slices"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	dist map[[2]int]int
}

func (p *p) Init(data []byte) {
	grid := bytes.Split(data, []byte{'\n'})
	x, y := len(grid), len(grid[0])
	idx := slices.Index(data, 'S')
	ori := [2]int{idx / (y + 1), idx % (y + 1)}
	dist := map[[2]int]int{ori: 0}
	q := [][2]int{ori}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, d := range util.Delta4 {
			nx, ny := cur[0]+d[0], cur[1]+d[1]
			next := [2]int{nx, ny}
			if nx < 0 || nx >= x || ny < 0 || ny >= y {
				continue
			}
			if grid[nx][ny] == '#' {
				continue
			}
			if _, ok := dist[next]; ok {
				continue
			}
			dist[next] = dist[cur] + 1
			q = append(q, [2]int{nx, ny})
		}
	}
	p.dist = dist
}

func (p *p) Solve1() any {
	cheats := map[[4]int]int{}
	check := func(start [2]int, iEnd, jEnd, distStart int) {
		if distEnd, ok := p.dist[[2]int{iEnd, jEnd}]; ok && distEnd > distStart+2 {
			cheats[[4]int{start[0], start[1], iEnd, jEnd}] = distEnd - distStart - 2
		}
	}

	for start, distStart := range p.dist {
		for i := 0; i <= 2; i++ {
			j := 2 - i
			check(start, start[0]+i, start[1]+j, distStart)
			check(start, start[0]+i, start[1]-j, distStart)
			check(start, start[0]-i, start[1]+j, distStart)
			check(start, start[0]-i, start[1]-j, distStart)
		}
	}

	count := 0
	for _, save := range cheats {
		if save >= 100 {
			count++
		}
	}
	return count
}

func (p *p) Solve2() any {
	cheats := map[[4]int]int{}
	check := func(start [2]int, iEnd, jEnd, distStart, cheatDist int) {
		if distEnd, ok := p.dist[[2]int{iEnd, jEnd}]; ok && distEnd > distStart+cheatDist {
			cheats[[4]int{start[0], start[1], iEnd, jEnd}] = distEnd - distStart - cheatDist
		}
	}

	for start, distStart := range p.dist {
		for cheatDist := 2; cheatDist <= 20; cheatDist++ {
			for i := 0; i <= cheatDist; i++ {
				j := cheatDist - i
				check(start, start[0]+i, start[1]+j, distStart, cheatDist)
				check(start, start[0]+i, start[1]-j, distStart, cheatDist)
				check(start, start[0]-i, start[1]+j, distStart, cheatDist)
				check(start, start[0]-i, start[1]-j, distStart, cheatDist)
			}
		}
	}

	count := 0
	for _, save := range cheats {
		if save >= 100 {
			count++
		}
	}
	return count
}

func init() {
	internal.Register(20, &p{})
}
