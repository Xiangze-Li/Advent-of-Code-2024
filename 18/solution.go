package puzzle18

import (
	"bytes"
	"fmt"
	"maps"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	pos [][2]int
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(data, []byte{'\n'})
	p.pos = make([][2]int, len(lines))
	for i, line := range lines {
		util.Must(fmt.Fscanf(bytes.NewReader(line), "%d,%d", &p.pos[i][0], &p.pos[i][1]))
	}
}

func bfs(wall map[[2]int]bool) int {
	q := [][2]int{{0, 0}}
	vis := map[[2]int]int{{0, 0}: 0}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for _, d := range util.Delta4 {
			next := [2]int{cur[0] + d[0], cur[1] + d[1]}
			if next[0] == 70 && next[1] == 70 {
				return vis[cur] + 1
			}

			if next[0] < 0 || 70 < next[0] || next[1] < 0 || 70 < next[1] {
				continue
			}
			if wall[next] {
				continue
			}
			if _, ok := vis[next]; ok {
				continue
			}
			vis[next] = vis[cur] + 1
			q = append(q, next)
		}
	}

	return -1
}

func (p *p) Solve1() any {
	wall := make(map[[2]int]bool, 1024)
	for i := range 1024 {
		wall[p.pos[i]] = true
	}
	return bfs(wall)
}

func (p *p) Solve2() any {
	wall := make(map[[2]int]bool, len(p.pos))
	for i := range 1024 {
		wall[p.pos[i]] = true
	}
	l, r := 1023, len(p.pos)-1
	wallL := maps.Clone(wall)
	for r-l > 1 {
		mid := (l + r) / 2
		for i := l; i <= mid; i++ {
			wall[p.pos[i]] = true
		}
		if bfs(wall) == -1 {
			r = mid
			wall = maps.Clone(wallL)
		} else {
			l = mid
			wallL = maps.Clone(wall)
		}
	}
	return fmt.Sprintf("%d,%d", p.pos[r][0], p.pos[r][1])
}

func init() {
	internal.Register(18, &p{})
}
