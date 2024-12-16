package puzzle16

import (
	"bytes"
	"container/heap"

	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {
	grids    [][]byte
	x, y     int
	ori, dst [2]int
}

func (p *p) Init(data []byte) {
	p.grids = bytes.Split(data, []byte{'\n'})
	p.x, p.y = len(p.grids), len(p.grids[0])
	p.ori = [2]int{p.x - 2, 1}
	p.dst = [2]int{1, p.y - 2}
}

type state struct {
	x, y int
	dir  util.Direction
}

func (p *p) dijkstra(ori state, dst [2]int) (map[state]int, state) {
	dist := map[state]int{}

	h := &heapIn{{ori, 0}}
	heap.Init(h)

	for h.Len() > 0 {
		cur := heap.Pop(h).(heapNode)
		dist[cur.state] = cur.dist
		if cur.state.x == dst[0] && cur.state.y == dst[1] {
			return dist, cur.state
		}
		for dir, delta := range util.Delta4 {
			nextX, nextY := cur.x+delta[0], cur.y+delta[1]
			if p.grids[nextX][nextY] == '#' {
				continue
			}
			var next state
			var nextCost int
			if dir == cur.dir {
				next = state{nextX, nextY, dir}
				nextCost = cur.dist + 1
			} else {
				next = state{cur.x, cur.y, dir}
				if dir == util.Opposite[cur.dir] {
					nextCost = cur.dist + 2000
				} else {
					nextCost = cur.dist + 1000
				}
			}
			if _, ok := dist[next]; !ok {
				h.Upsert(next, nextCost)
			}
		}
	}

	return nil, state{}
}

func (p *p) Solve1() any {
	dist, final := p.dijkstra(state{p.ori[0], p.ori[1], util.E}, p.dst)

	return dist[final]
}

func (p *p) Solve2() any {
	distOri, final := p.dijkstra(state{p.ori[0], p.ori[1], util.E}, p.dst)
	distDst, _ := p.dijkstra(state{p.dst[0], p.dst[1], util.Opposite[final.dir]}, p.ori)
	finalCost := distOri[final]

	cords := map[[2]int]bool{}
	for s, cost := range distOri {
		for dir := range util.Delta4 {
			if cost+distDst[state{s.x, s.y, dir}] == finalCost {
				cords[[2]int{s.x, s.y}] = true
			}
		}
	}
	return len(cords)
}

func init() {
	internal.Register(16, &p{})
}

type (
	heapNode struct {
		state
		dist int
	}
	heapIn []heapNode
)

func (h *heapIn) Len() int           { return len(*h) }
func (h *heapIn) Less(i, j int) bool { return (*h)[i].dist < (*h)[j].dist }
func (h *heapIn) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *heapIn) Push(x any)         { *h = append(*h, x.(heapNode)) }
func (h *heapIn) Pop() any           { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }
func (h *heapIn) Upsert(s state, d int) {
	for i, n := range *h {
		if n.state == s {
			if n.dist > d {
				(*h)[i].dist = d
				heap.Fix(h, i)
				return
			}
		}
	}
	heap.Push(h, heapNode{s, d})
}
