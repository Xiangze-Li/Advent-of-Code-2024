package puzzle23

import (
	"bytes"
	"maps"
	"slices"
	"strings"

	"github.com/Xiangze-Li/advent-2024/internal"
)

type p struct {
	adj   map[string][]string
	names []string
}

func (p *p) Init(data []byte) {
	lines := bytes.Split(data, []byte("\n"))
	p.adj = make(map[string][]string)
	names := make(map[string]bool, 2*len(lines))
	for _, line := range lines {
		sp := strings.SplitN(string(line), "-", 2)
		p.adj[sp[0]] = append(p.adj[sp[0]], sp[1])
		p.adj[sp[1]] = append(p.adj[sp[1]], sp[0])
		names[sp[0]] = true
		names[sp[1]] = true
	}
	p.names = slices.Collect(maps.Keys(names))
}

func (p *p) Solve1() any {
	count := 0

	for _, name := range p.names {
		for _, adj0 := range p.adj[name] {
			for _, adj1 := range p.adj[name] {
				if adj0 == adj1 {
					continue
				}
				if name[0] != 't' && adj0[0] != 't' && adj1[0] != 't' {
					continue
				}

				if slices.Contains(p.adj[adj0], adj1) {
					count++
				}
			}
		}
	}

	return count / 6
}

func intersect(a, b []string) []string {
	m := make(map[string]bool)
	for _, x := range a {
		m[x] = true
	}
	var c []string
	for _, x := range b {
		if m[x] {
			c = append(c, x)
		}
	}
	return c
}

func (p *p) Solve2() any {
	maxiClique := []string{}
	var bronKerbosch func(R, P, X []string)

	bronKerbosch = func(R, P, X []string) {
		if len(P) == 0 && len(X) == 0 {
			if len(R) > len(maxiClique) {
				maxiClique = slices.Clone(R)
			}
			return
		}
		for i := len(P) - 1; i >= 0; i-- {
			v := P[i]
			rr := append(slices.Clone(R), v)
			pp := intersect(P, p.adj[v])
			xx := intersect(X, p.adj[v])
			bronKerbosch(rr, pp, xx)
			P = P[:i]
			X = append(X, v)
		}
	}

	bronKerbosch(nil, slices.Clone(p.names), nil)

	slices.Sort(maxiClique)
	return strings.Join(maxiClique, ",")
}

func init() {
	internal.Register(23, &p{})
}
