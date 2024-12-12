package puzzle12

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
	p.grids = bytes.Split(data, []byte("\n"))
	p.x, p.y = len(p.grids), len(p.grids[0])
}

func (p *p) bfs1(i, j int, vis *[][]bool) int {
	region := util.SliceND[bool](p.x, p.y).([][]bool)
	mark := p.grids[i][j]

	(*vis)[i][j] = true
	region[i][j] = true
	queue := [][2]int{{i, j}}

	area := 1
	perimeter := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, d := range util.Delta4 {
			x, y := cur[0]+d[0], cur[1]+d[1]
			if x < 0 || x >= p.x || y < 0 || y >= p.y {
				perimeter++
				continue
			}
			if p.grids[x][y] != mark {
				perimeter++
				continue
			}
			if region[x][y] {
				continue
			}
			(*vis)[x][y] = true
			region[x][y] = true
			area++
			queue = append(queue, [2]int{x, y})
		}
	}

	return area * perimeter
}

func (p *p) Solve1() any {
	acc := int(0)
	vis := util.SliceND[bool](p.x, p.y).([][]bool)

	for i := range p.grids {
		for j := range p.grids[i] {
			if vis[i][j] {
				continue
			}
			acc += p.bfs1(i, j, &vis)
		}
	}

	return acc
}

func (p *p) bfs2(i, j int, vis *[][]bool) int {
	region := util.SliceND[bool](p.x, p.y).([][]bool)
	mark := p.grids[i][j]

	(*vis)[i][j] = true
	region[i][j] = true
	queue := [][2]int{{i, j}}

	minX, maxX := i, i
	minY, maxY := j, j
	area := 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, d := range util.Delta4 {
			x, y := cur[0]+d[0], cur[1]+d[1]
			if x < 0 || x >= p.x || y < 0 || y >= p.y {
				continue
			}
			if p.grids[x][y] != mark {
				continue
			}
			if region[x][y] {
				continue
			}
			(*vis)[x][y] = true
			region[x][y] = true
			area++
			queue = append(queue, [2]int{x, y})
			minX, maxX = min(minX, x), max(maxX, x)
			minY, maxY = min(minY, y), max(maxY, y)
		}
	}

	return area * getSides(region, minX, maxX, minY, maxY)
}

//nolint:gocognit
func getSides(region [][]bool, minX, maxX, minY, maxY int) int {
	state := func(x, y int) bool {
		if x < minX || x > maxX || y < minY || y > maxY {
			return false
		}
		return region[x][y]
	}

	perimeter := 0
	for i := minX; i <= maxX; i++ {
		st := state(i, minY-1)
		for j := minY; j <= maxY+1; j++ {
			if st != state(i, j) {
				if st != state(i-1, j-1) || st == state(i-1, j) {
					perimeter++
				}
				if st != state(i+1, j-1) || st == state(i+1, j) {
					perimeter++
				}
				st = !st
			}
		}
	}

	for j := minY; j <= maxY; j++ {
		st := state(minX-1, j)
		for i := minX; i <= maxX+1; i++ {
			if st != state(i, j) {
				if st != state(i-1, j-1) || st == state(i, j-1) {
					perimeter++
				}
				if st != state(i-1, j+1) || st == state(i, j+1) {
					perimeter++
				}
				st = !st
			}
		}
	}

	return perimeter / 2
}

func (p *p) Solve2() any {
	acc := int(0)
	vis := util.SliceND[bool](p.x, p.y).([][]bool)

	for i := range p.grids {
		for j := range p.grids[i] {
			if vis[i][j] {
				continue
			}
			acc += p.bfs2(i, j, &vis)
		}
	}

	return acc
}

func init() {
	internal.Register(12, &p{})
}
