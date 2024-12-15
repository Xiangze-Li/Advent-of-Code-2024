package puzzle15

import (
	"github.com/Xiangze-Li/advent-2024/util"
)

type warehouse1 struct {
	gg   [][]byte
	x, y int
}

func (w *warehouse1) Init(grid [][]byte) (int, int) {
	w.gg = util.Clone(grid)
	w.x = len(grid)
	w.y = len(grid[0])
	for i := range w.x {
		for j := range w.y {
			if w.gg[i][j] == '@' {
				w.gg[i][j] = '.'
				return i, j
			}
		}
	}
	return -1, -1
}

func (w *warehouse1) Push(i, j int, dir util.Direction) bool {
	switch w.gg[i][j] {
	case '.':
		return true
	case '#':
		return false
	case 'O':
		delta := util.Delta4[dir]
		if !w.Push(i+delta[0], j+delta[1], dir) {
			return false
		}
		w.gg[i][j] = '.'
		w.gg[i+delta[0]][j+delta[1]] = 'O'
	}
	return true
}

type warehouse2 struct {
	gg   [][]byte
	x, y int
}

func (w *warehouse2) Init(grid [][]byte) (int, int) {
	w.x = len(grid)
	w.y = 2 * len(grid[0])
	w.gg = util.SliceND[byte](w.x, w.y).([][]byte)
	var ii, jj int
	for i, row := range grid {
		for j, cell := range row {
			switch cell {
			case '#':
				w.gg[i][j*2] = '#'
				w.gg[i][j*2+1] = '#'
			case 'O':
				w.gg[i][j*2] = '['
				w.gg[i][j*2+1] = ']'
			case '@':
				ii = i
				jj = j * 2
				fallthrough
			case '.':
				w.gg[i][j*2] = '.'
				w.gg[i][j*2+1] = '.'
			}
		}
	}
	return ii, jj
}

func (w *warehouse2) Push(dir util.Direction, i, j int) bool {
	switch dir {
	case util.N, util.S:
		return w.pushX(dir, i, j)
	case util.W, util.E:
		return w.pushY(dir, i, j)
	default:
		panic("invalid direction")
	}
}

func (w *warehouse2) pushX(dir util.Direction, cords ...int) bool {
	delta := util.Delta4[dir]
	boxes := []int{}
	for k := 0; k < len(cords); k += 2 {
		switch w.gg[cords[k]][cords[k+1]] {
		case '[':
			boxes = append(boxes, cords[k]+delta[0], cords[k+1], cords[k]+delta[0], cords[k+1]+1)
		case ']':
			boxes = append(boxes, cords[k]+delta[0], cords[k+1]-1, cords[k]+delta[0], cords[k+1])
		case '#':
			return false
		case '.': // no-op
		}
	}
	if len(boxes) == 0 {
		return true
	}
	if !w.pushX(dir, boxes...) {
		return false
	}
	for k := 0; k < len(boxes); k += 4 {
		w.gg[boxes[k]][boxes[k+1]] = '['
		w.gg[boxes[k]-delta[0]][boxes[k+1]] = '.'
		w.gg[boxes[k+2]][boxes[k+3]] = ']'
		w.gg[boxes[k+2]-delta[0]][boxes[k+3]] = '.'
	}
	return true
}

func (w *warehouse2) pushY(dir util.Direction, i, j int) bool {
	switch w.gg[i][j] {
	case '.':
		return true
	case '#':
		return false
	case '[', ']':
		delta := util.Delta4[dir]
		if !w.pushY(dir, i, j+delta[1]) {
			return false
		}
		w.gg[i][j+delta[1]] = w.gg[i][j]
		w.gg[i][j] = '.'
	}
	return true
}
