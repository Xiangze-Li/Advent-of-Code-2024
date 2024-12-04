package puzzle04

import (
	"bytes"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	grids [][]byte
	i, j  int
}

func (p *p) Init(data []byte) {
	p.grids = bytes.Split(data, []byte{'\n'})
	p.i, p.j = len(p.grids), len(p.grids[0])
}

func (p *p) Solve1() any {
	count := 0
	clip := func(i, j int) bool {
		return 0 <= i && i < p.i && 0 <= j && j < p.j
	}

	for i := range p.i {
		for j := range p.j {
			if p.grids[i][j] != 'X' {
				continue
			}
			for _, diff := range util.Delta8 {
				if !clip(i+diff[0]*3, j+diff[1]*3) {
					continue
				}
				if p.grids[i+diff[0]][j+diff[1]] != 'M' ||
					p.grids[i+diff[0]*2][j+diff[1]*2] != 'A' ||
					p.grids[i+diff[0]*3][j+diff[1]*3] != 'S' {
					continue
				}
				count++
			}
		}
	}

	return count
}

func (p *p) Solve2() any {
	count := 0

	for i := 1; i < p.i-1; i++ {
		for j := 1; j < p.j-1; j++ {
			if p.grids[i][j] != 'A' {
				continue
			}
			mas := 0
			for dir, diff := range util.Diagonal4 {
				op := util.Delta8[util.Opposite[dir]]
				if p.grids[i+diff[0]][j+diff[1]] == 'M' &&
					p.grids[i+op[0]][j+op[1]] == 'S' {
					mas++
				}
			}
			if mas == 2 {
				count++
			}
		}
	}

	return count
}

func init() {
	internal.Register(4, &p{})
}
