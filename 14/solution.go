package puzzle14

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	p [][2]int
	v [][2]int
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(data, []byte("\n"))
	p.p = make([][2]int, len(lines))
	p.v = make([][2]int, len(lines))
	for i, line := range lines {
		util.Must(fmt.Fscanf(bytes.NewReader(line),
			"p=%d,%d v=%d,%d", &p.p[i][0], &p.p[i][1], &p.v[i][0], &p.v[i][1]))
	}
}

const x, y = 101, 103

func (p *p) Solve1() any {
	pos := slices.Clone(p.p)

	for range 100 {
		for i := range pos {
			pos[i][0] = (pos[i][0] + p.v[i][0] + x) % x
			pos[i][1] = (pos[i][1] + p.v[i][1] + y) % y
		}
	}

	q := [4]int{}
	for _, pp := range pos {
		switch {
		case pp[0] > x/2 && pp[1] > y/2:
			q[0]++
		case pp[0] > x/2 && pp[1] < y/2:
			q[1]++
		case pp[0] < x/2 && pp[1] > y/2:
			q[2]++
		case pp[0] < x/2 && pp[1] < y/2:
			q[3]++
		}
	}
	return q[0] * q[1] * q[2] * q[3]
}

func (p *p) Solve2() any {
	pos := slices.Clone(p.p)
	draw := func() {
		img := util.SliceND[bool](y, x).([][]bool)
		for _, pp := range pos {
			img[pp[1]][pp[0]] = true
		}
		var b strings.Builder
		b.Grow((x + 1) * y)
		for j := range y {
			for i := range x {
				if img[j][i] {
					b.WriteByte('#')
				} else {
					b.WriteByte('.')
				}
			}
			b.WriteByte('\n')
		}
		fmt.Print(b.String())
	}

	for step := 1; true; step++ {
		for i := range pos {
			pos[i][0] = (pos[i][0] + p.v[i][0] + x) % x
			pos[i][1] = (pos[i][1] + p.v[i][1] + y) % y
		}
		vis := make(map[[2]int]bool, len(pos))
		for _, pp := range pos {
			vis[pp] = true
		}
		if len(vis) == len(pos) {
			draw()
			return step
		}
	}
	return 0
}

func init() {
	internal.Register(14, &p{})
}
