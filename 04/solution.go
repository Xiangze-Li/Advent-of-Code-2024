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

func (p *p) clip(i, j, pad int) bool {
	return pad <= i && i+pad < p.i && pad <= j && j+pad < p.j
}

func (p *p) Init(data []byte) {
	p.grids = bytes.Split(data, []byte{'\n'})
	p.i, p.j = len(p.grids), len(p.grids[0])
}

func (p *p) Solve1() any {
	const target = "XMAS"
	count := 0

	for i, row := range p.grids {
		for j, cell := range row {
			if cell != 'X' || !p.clip(i, j, 3) {
				continue
			}
			for _, diff := range util.Delta8 {
				ii, jj := i, j
				for k := range 3 {
					ii, jj = ii+diff[0], jj+diff[1]
					if p.grids[ii][jj] != target[k+1] {
						break
					}
					if k == 2 {
						count++
					}
				}
			}
		}
	}

	return count
}

func (p *p) Solve2() any {
	count := 0

	for i, row := range p.grids {
		for j, cell := range row {
			if cell != 'A' || !p.clip(i, j, 1) {
				continue
			}
			mas := 0
			for dir, diff := range util.Diagonal4 {
				if p.grids[i+diff[0]][j+diff[1]] == 'M' {
					op := util.Delta8[util.Opposite[dir]]
					if p.grids[i+op[0]][j+op[1]] == 'S' {
						mas++
					}
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
