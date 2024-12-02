package puzzle02

import (
	"bytes"
	"cmp"
	"slices"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	reports [][]int64
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	p.reports = make([][]int64, 0, len(lines))
	for _, line := range lines {
		sp := bytes.Fields(line)
		p.reports = append(p.reports, util.ArrayBytesToInt64(sp))
	}
}

func check(report []int64) bool {
	if len(report) <= 1 {
		return true
	}

	asc := cmp.Compare(report[0], report[1])
	if asc == 0 {
		return false
	}
	for i := 1; i < len(report); i++ {
		if asc != cmp.Compare(report[i-1], report[i]) {
			return false
		}
		diff := report[i] - report[i-1]
		if !(-3 <= diff && diff <= 3) {
			return false
		}
	}

	return true
}

func (p *p) Solve1() any {
	count := 0

	for _, report := range p.reports {
		if check(report) {
			count++
		}
	}

	return count
}

func (p *p) Solve2() any {
	count := 0

	for _, report := range p.reports {
		if check(report) {
			count++
			continue
		}

		for removePos := range report {
			if check(slices.Delete(slices.Clone(report), removePos, removePos+1)) {
				count++
				break
			}
		}
	}

	return count
}

func init() {
	internal.Register(2, &p{})
}
