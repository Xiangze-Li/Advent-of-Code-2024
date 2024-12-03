package puzzle03

import (
	"regexp"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	data []byte
}

func (p *p) Init(data []byte) {
	p.data = data
}

func (p *p) Solve1() any {
	matches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllSubmatch(p.data, -1)
	acc := int64(0)
	for _, match := range matches {
		cov := util.ArrayBytesToInt64(match[1:])
		acc += cov[0] * cov[1]
	}
	return acc
}

func (p *p) Solve2() any {
	matches := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`).
		FindAllSubmatch(p.data, -1)
	acc := int64(0)
	enabled := true

	for _, match := range matches {
		switch string(match[0]) {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if !enabled {
				continue
			}
			cov := util.ArrayBytesToInt64(match[1:])
			acc += cov[0] * cov[1]
		}
	}

	return acc
}

func init() {
	internal.Register(3, &p{})
}
