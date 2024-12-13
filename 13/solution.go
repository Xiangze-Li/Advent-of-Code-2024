package puzzle13

import (
	"bytes"
	"regexp"
	"strconv"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	n      int
	deltaA [][2]int
	deltaB [][2]int
	prize  [][2]int
}

var re = regexp.MustCompile(`[XY][+=](\d+)`)

func (p *p) Init(data []byte) {
	items := bytes.Split(data, []byte("\n\n"))
	p.n = len(items)
	p.deltaA = make([][2]int, p.n)
	p.deltaB = make([][2]int, p.n)
	p.prize = make([][2]int, p.n)
	for i, item := range items {
		matches := re.FindAll(item, 6)
		p.deltaA[i][0] = util.Must(strconv.Atoi(string(matches[0][2:])))
		p.deltaA[i][1] = util.Must(strconv.Atoi(string(matches[1][2:])))
		p.deltaB[i][0] = util.Must(strconv.Atoi(string(matches[2][2:])))
		p.deltaB[i][1] = util.Must(strconv.Atoi(string(matches[3][2:])))
		p.prize[i][0] = util.Must(strconv.Atoi(string(matches[4][2:])))
		p.prize[i][1] = util.Must(strconv.Atoi(string(matches[5][2:])))
	}
}

func solveCase(ax, ay, bx, by, px, py int) int {
	a := (bx*py - by*px) / (ay*bx - ax*by)
	if a < 0 || a*(ay*bx-ax*by) != (bx*py-by*px) {
		return 0
	}
	b := (px - ax*a) / bx
	if b < 0 || b*bx != px-ax*a {
		return 0
	}
	return 3*a + b
}

func (p *p) Solve1() any {
	token := 0

	for i := range p.n {
		ax, ay := p.deltaA[i][0], p.deltaA[i][1]
		bx, by := p.deltaB[i][0], p.deltaB[i][1]
		px, py := p.prize[i][0], p.prize[i][1]

		token += solveCase(ax, ay, bx, by, px, py)
	}

	return token
}

func (p *p) Solve2() any {
	const bias = 10000000000000
	token := 0

	for i := range p.n {
		ax, ay := p.deltaA[i][0], p.deltaA[i][1]
		bx, by := p.deltaB[i][0], p.deltaB[i][1]
		px, py := p.prize[i][0]+bias, p.prize[i][1]+bias

		token += solveCase(ax, ay, bx, by, px, py)
	}

	return token
}

func init() {
	internal.Register(13, &p{})
}
