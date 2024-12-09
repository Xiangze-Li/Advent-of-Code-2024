package puzzle09

import (
	"slices"

	"github.com/Xiangze-Li/advent-2024/internal"
)

type p struct {
	grids  []int
	spaces [][2]int
	files  [][3]int
}

func (p *p) Init(data []byte) {
	fileIdx := 0
	for i, c := range data {
		if c == '0' {
			continue
		}
		size := int(c - '0')
		if i%2 == 0 {
			p.files = append(p.files, [3]int{fileIdx, len(p.grids), size})
			p.grids = append(p.grids, slices.Repeat([]int{fileIdx}, size)...)
			fileIdx++
		} else {
			p.spaces = append(p.spaces, [2]int{len(p.grids), size})
			p.grids = append(p.grids, slices.Repeat([]int{-1}, size)...)
		}
	}
}

func (p *p) Solve1() any {
	l, r := 0, len(p.grids)-1
	grids := slices.Clone(p.grids)

	for l < r {
		for grids[l] != -1 {
			l++
		}
		for grids[r] == -1 {
			r--
		}
		if l >= r {
			break
		}
		grids[l], grids[r] = grids[r], grids[l]
		l++
		r--
	}

	var sum uint64
	for i, v := range grids {
		if v != -1 {
			sum += uint64(i) * uint64(v)
		}
	}
	return sum
}

func (p *p) Solve2() any {
	files := slices.Clone(p.files)
	spaces := slices.Clone(p.spaces)

	for i := len(files) - 1; i >= 0; i-- {
		size := files[i][2]
		for j, space := range spaces {
			if space[0] > files[i][1] {
				break
			}
			if space[1] > size {
				files[i][1] = space[0]
				spaces[j][0] += size
				spaces[j][1] -= size
				break
			}
			if space[1] == size {
				files[i][1] = space[0]
				spaces = append(spaces[:j], spaces[j+1:]...)
				break
			}
		}
	}

	var sum uint64
	for _, file := range files {
		for j := range file[2] {
			sum += uint64(file[0]) * uint64(file[1]+j)
		}
	}
	return sum
}

func init() {
	internal.Register(9, &p{})
}
