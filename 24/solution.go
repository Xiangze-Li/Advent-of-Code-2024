package puzzle24

import (
	"bytes"
	"fmt"
	"io"
	"maps"

	"github.com/Xiangze-Li/advent-2024/internal"
)

const (
	and = "AND"
	or  = "OR"
	xor = "XOR"
)

type gate struct {
	src1, src2 string
	op         string
}

type p struct {
	value map[string]bool
	gates map[string]gate
}

func (p *p) Init(data []byte) {
	sp := bytes.SplitN(data, []byte("\n\n"), 2)
	p.value = make(map[string]bool)
	p.gates = make(map[string]gate)
	for _, line := range bytes.Split(sp[0], []byte("\n")) {
		p.value[string(line[0:3])] = line[5] == '1'
	}
	for _, line := range bytes.Split(sp[1], []byte("\n")) {
		spp := bytes.SplitN(line, []byte(" "), 5)
		src1, src2 := string(spp[0]), string(spp[2])
		if src1 > src2 {
			src1, src2 = src2, src1
		}
		p.gates[string(spp[4])] = gate{src1, src2, string(spp[1])}
	}
}

func (p *p) getVal(tgt string) bool {
	if v, ok := p.value[tgt]; ok {
		return v
	}
	g := p.gates[tgt]
	v1, v2 := p.getVal(g.src1), p.getVal(g.src2)
	switch g.op {
	case and:
		p.value[tgt] = v1 && v2
	case or:
		p.value[tgt] = v1 || v2
	case xor:
		p.value[tgt] = v1 != v2
	}
	return p.value[tgt]
}

func (p *p) Solve1() any {
	var z uint64

	for i := 63; i >= 0; i-- {
		z <<= 1
		tgt := fmt.Sprintf("z%02d", i)
		if _, ok := p.gates[tgt]; !ok {
			continue
		}
		if p.getVal(tgt) {
			z |= 1
		}
	}

	return z
}

/*
CREDIT: https://github.com/ypisetsky/advent-of-code/blob/main/yr2024/day24.py

Code below is a rework of the original Python code to Go.
*/

var (
	blockee = map[string]map[string]bool{}
	blocker = map[string]int{}
	visited = map[string]bool{}
)

func diff(l, r map[string]bool) map[string]bool {
	res := map[string]bool{}
	for k := range l {
		if _, ok := r[k]; !ok {
			res[k] = true
		}
	}
	return res
}

func (p *p) step(w io.Writer, q []string) {
	for len(q) > 0 {
		tgt := q[0]
		q = q[1:]

		visited[tgt] = true
		if gate, ok := p.gates[tgt]; ok {
			fmt.Fprintf(w, "    OP %s %s %s => %s\n", gate.src1, gate.op, gate.src2, tgt)
		} else {
			fmt.Fprintf(w, "    IN %s\n", tgt)
		}

		for k := range blockee[tgt] {
			blocker[k]--
			if blocker[k] == 0 {
				q = append(q, k)
			}
		}
	}
}

func (p *p) oprandOf(oprand, target string) bool {
	gate, ok := p.gates[target]
	if !ok {
		return false
	}
	return gate.src1 == oprand || gate.src2 == oprand
}

//nolint:gocognit,funlen,gocritic,gocyclo,cyclop // manual analysis
func (p *p) analyze(w io.Writer, newVisits map[string]bool, i int, carryIn string) string {
	x := fmt.Sprintf("x%02d", i)
	y := fmt.Sprintf("y%02d", i)

	if i == 0 {
		var resultXOR, carryAND string
		for tgt := range newVisits {
			if tgt[0] == 'x' && tgt != x || tgt[0] == 'y' && tgt != y {
				fmt.Fprintln(w, "E: unexpected constant", tgt)
				continue
			}

			gate := p.gates[tgt]
			switch gate.op {
			case xor:
				if tgt != "z00" || resultXOR != "" {
					fmt.Fprintln(w, "E: unexpected XORs", tgt, resultXOR)
				}
				resultXOR = tgt
			case and:
				if carryAND != "" {
					fmt.Fprintln(w, "E: unexpected ANDs", tgt, carryAND)
				}
				carryAND = tgt
			}
		}
		fmt.Fprintf(w, "resultXOR: %q carryAND: %q\n", resultXOR, carryAND)
		return carryAND
	}

	if !newVisits[x] {
		fmt.Fprintln(w, "E: missing x", x)
	}
	if !newVisits[y] {
		fmt.Fprintln(w, "E: missing y", y)
	}
	if len(newVisits) != 7 {
		fmt.Fprintln(w, "E: len(newVisits) != 7")
	}

	var localXOR, localAND, resultXOR, carryAND, carryOR string

	for tgt := range newVisits {
		if tgt[0] == 'x' && tgt != x || tgt[0] == 'y' && tgt != y {
			fmt.Fprintln(w, "E: unexpected constant", tgt)
			continue
		}
		gate := p.gates[tgt]
		switch gate.op {
		case xor:
			if p.oprandOf(x, tgt) && p.oprandOf(y, tgt) {
				localXOR = tgt
			} else if resultXOR == "" {
				resultXOR = tgt
			} else {
				fmt.Fprintln(w, "E: unexpected XORs", tgt, resultXOR)
			}
		case and:
			if p.oprandOf(x, tgt) && p.oprandOf(y, tgt) {
				localAND = tgt
			} else if carryAND == "" {
				carryAND = tgt
			} else {
				fmt.Fprintln(w, "E: unexpected ANDs", tgt, carryAND)
			}
		case or:
			if carryOR == "" {
				carryOR = tgt
			} else {
				fmt.Fprintln(w, "E: unexpected ORs", tgt, carryOR)
			}
		}
	}

	if !p.oprandOf(carryIn, resultXOR) {
		fmt.Fprintf(w, "E: carryIn %s is not an operand of resultXOR %s\n", carryIn, resultXOR)
	}
	if !p.oprandOf(carryIn, carryAND) {
		fmt.Fprintf(w, "E: carryIn %s is not an operand of carryAND %s\n", carryIn, carryAND)
	}
	if !p.oprandOf(localXOR, resultXOR) {
		fmt.Fprintf(w, "E: localXOR %s is not an operand of resultXOR %s\n", localXOR, resultXOR)
	}
	if !p.oprandOf(localXOR, carryAND) {
		fmt.Fprintf(w, "E: localXOR %s is not an operand of carryAND %s\n", localXOR, carryAND)
	}
	if !p.oprandOf(localAND, carryOR) {
		fmt.Fprintf(w, "E: localAND %s is not an operand of carryOR %s\n", localAND, carryOR)
	}
	if !p.oprandOf(carryAND, carryOR) {
		fmt.Fprintf(w, "E: carryAND %s is not an operand of carryOR %s\n", carryAND, carryOR)
	}

	fmt.Fprintf(w, "carryIn: %q localXOR: %q localAND: %q resultXOR: %q carryAND: %q carryOR: %q\n",
		carryIn, localXOR, localAND, resultXOR, carryAND, carryOR)

	return carryOR
}

func (p *p) Solve2() any {
	prev := map[string]bool{}
	var carry string

	for tgt, gate := range p.gates {
		if blockee[gate.src1] == nil {
			blockee[gate.src1] = map[string]bool{}
		}
		if blockee[gate.src2] == nil {
			blockee[gate.src2] = map[string]bool{}
		}
		blockee[gate.src1][tgt] = true
		blockee[gate.src2][tgt] = true
		blocker[tgt] = 2
	}

	w := new(bytes.Buffer)

	for i := range 45 {
		fmt.Fprintf(w, "=== %02d ===\n", i)
		q := []string{fmt.Sprintf("x%02d", i), fmt.Sprintf("y%02d", i)}
		p.step(w, q)
		carry = p.analyze(w, diff(visited, prev), i, carry)
		prev = maps.Clone(visited)
	}

	return w.String() + "\nMANUAL ANALYSIS\n" +
		"CREDIT: https://github.com/ypisetsky/advent-of-code/blob/main/yr2024/day24.py\n" +
		"dqr,dtk,pfw,shh,vgs,z21,z33,z39"
}

func init() {
	internal.Register(24, &p{})
}
