/*
Package puzzle21

CREDIT: https://www.reddit.com/r/adventofcode/comments/1hj2odw/comment/m33uf55/
*/
package puzzle21

import (
	"bytes"
	"slices"
	"strconv"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	codes        [][]byte
	digit, direc map[byte][2]int
}

func (p *p) Init(data []byte) {
	p.codes = bytes.Split(data, []byte("\n"))
	p.digit = make(map[byte][2]int, 12)
	for i, c := range []byte("789456123 0A") {
		p.digit[c] = [2]int{i % 3, i / 3}
	}
	p.direc = make(map[byte][2]int, 6)
	for i, c := range []byte(" ^A<v>") {
		p.direc[c] = [2]int{i % 3, i / 3}
	}
}

func steps(kbd map[byte][2]int, code []byte, times int) map[[3]int]int {
	acc := map[[3]int]int{}
	x, y := kbd['A'][0], kbd['A'][1]
	xGap, yGap := kbd[' '][0], kbd[' '][1]
	for _, c := range code {
		xN, yN := kbd[c][0], kbd[c][1]
		gap := 0
		if (xN == xGap && y == yGap) || (x == xGap && yN == yGap) {
			gap = 1
		}
		acc[[3]int{xN - x, yN - y, gap}] += times
		x, y = xN, yN
	}
	return acc
}

func (p *p) next(step map[[3]int]int) map[[3]int]int {
	next := map[[3]int]int{}
	for k, v := range step {
		dx, dy := k[0], k[1]
		codeN := []byte{}
		if dy > 0 {
			codeN = append(codeN, bytes.Repeat([]byte{'v'}, dy)...)
		} else {
			codeN = append(codeN, bytes.Repeat([]byte{'^'}, -dy)...)
		}
		if dx < 0 {
			codeN = append(bytes.Repeat([]byte{'<'}, -dx), codeN...)
		} else {
			codeN = append(codeN, bytes.Repeat([]byte{'>'}, dx)...)
		}
		if k[2] != 0 {
			slices.Reverse(codeN)
		}
		codeN = append(codeN, 'A')
		nn := steps(p.direc, codeN, v)
		for k, v := range nn {
			next[k] += v
		}
	}
	return next
}

func (p *p) Solve1() any {
	acc := 0

	for _, code := range p.codes {
		step := steps(p.digit, code, 1)
		for range 3 {
			step = p.next(step)
		}
		c := util.Must(strconv.Atoi(string(code[:len(code)-1])))
		for _, v := range step {
			acc += v * c
		}
	}

	return acc
}

func (p *p) Solve2() any {
	acc := 0

	for _, code := range p.codes {
		step := steps(p.digit, code, 1)
		for range 26 {
			step = p.next(step)
		}
		c := util.Must(strconv.Atoi(string(code[:len(code)-1])))
		for _, v := range step {
			acc += v * c
		}
	}

	return acc
}

func init() {
	internal.Register(21, &p{})
}
