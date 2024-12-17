package puzzle17

import (
	"bytes"
	"slices"
	"strconv"
	"strings"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	prog    []uint64
	a, b, c uint64
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(data, []byte("\n"))
	p.prog = util.ArrayBytesToUint64(bytes.Split(lines[4][9:], []byte(",")))
	p.a = util.Must(strconv.ParseUint(string(lines[0][12:]), 10, 64))
	p.b = util.Must(strconv.ParseUint(string(lines[1][12:]), 10, 64))
	p.c = util.Must(strconv.ParseUint(string(lines[2][12:]), 10, 64))
}

func (p *p) Solve1() any {
	out, count := compiledRun(p.a)
	outs := []string{}
	for range count {
		outs = append(outs, strconv.FormatUint(out&0b111, 10))
		out >>= 3
	}
	slices.Reverse(outs)
	return strings.Join(outs, ",")
}

func (p *p) Solve2() any {
	prog := slices.Clone(p.prog)
	slices.Reverse(prog)

	guesses := []uint64{0}
	for _, out := range prog {
		next := []uint64{}
		for _, a := range guesses {
			next = append(next, fromOutput(a, out)...)
		}
		guesses = next
	}
	slices.Sort(guesses)
	return guesses[0]
}

func init() {
	internal.Register(17, &p{})
}
