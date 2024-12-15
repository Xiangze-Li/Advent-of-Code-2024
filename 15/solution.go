package puzzle15

import (
	"bytes"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	grids [][]byte
	moves []byte
}

func (p *p) Init(data []byte) {
	sp := bytes.SplitN(data, []byte("\n\n"), 2)
	p.grids = bytes.Split(sp[0], []byte("\n"))
	p.moves = bytes.ReplaceAll(sp[1], []byte("\n"), nil)
}

func (p *p) Solve1() any {
	g := warehouse1{}
	i, j := g.Init(p.grids)

	for _, m := range p.moves {
		dir := util.ConvertFrom[m]
		delta := util.Delta4[dir]
		if g.Push(i+delta[0], j+delta[1], dir) {
			i += delta[0]
			j += delta[1]
		}
	}

	acc := 0
	for i := range g.x {
		for j := range g.y {
			if g.gg[i][j] == 'O' {
				acc += 100*i + j
			}
		}
	}
	return acc
}

func (p *p) Solve2() any {
	g := warehouse2{}
	i, j := g.Init(p.grids)

	for _, m := range p.moves {
		dir := util.ConvertFrom[m]
		delta := util.Delta4[dir]
		if g.Push(dir, i+delta[0], j+delta[1]) {
			i += delta[0]
			j += delta[1]
		}
	}

	acc := 0
	for i := range g.x {
		for j := range g.y {
			if g.gg[i][j] == '[' {
				acc += 100*i + j
			}
		}
	}
	return acc
}

func init() {
	internal.Register(15, &p{})
}
