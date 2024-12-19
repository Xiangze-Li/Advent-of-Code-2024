package puzzle19

import (
	"bytes"
	"strings"

	"github.com/Xiangze-Li/advent-2024/internal"
)

type p struct {
	towels  []string
	designs []string
	mem     map[string]int
}

func (p *p) Init(data []byte) {
	sp := bytes.SplitN(data, []byte("\n\n"), 2)
	p.towels = strings.Split(string(sp[0]), ", ")
	p.designs = strings.Split(string(sp[1]), "\n")
	p.mem = make(map[string]int)
}

func (p *p) build(design string) int {
	if v, ok := p.mem[design]; ok {
		return v
	}

	count := 0
	for _, towel := range p.towels {
		if strings.HasPrefix(design, towel) {
			next := design[len(towel):]
			if len(next) == 0 {
				count++
			} else {
				count += p.build(next)
			}
		}
	}

	p.mem[design] = count
	return count
}

func (p *p) Solve1() any {
	count := 0

	for _, design := range p.designs {
		if p.build(design) > 0 {
			count++
		}
	}

	return count
}

func (p *p) Solve2() any {
	count := 0

	for _, design := range p.designs {
		count += p.build(design)
	}

	return count
}

func init() {
	internal.Register(19, &p{})
}
