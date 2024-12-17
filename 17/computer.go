//nolint:unused // preserved for record
package puzzle17

type computer struct {
	prog    []uint64
	ip      uint64
	a, b, c uint64
}

func (c *computer) getVal(oprand uint64) uint64 {
	switch oprand {
	case 0, 1, 2, 3:
		return oprand
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	default:
		panic("invalid oprand")
	}
}

func (c *computer) Run() []uint64 {
	out := []uint64{}
	l := uint64(len(c.prog))
	for c.ip < l {
		opcode := c.prog[c.ip]
		oprand := c.prog[c.ip+1]
		switch opcode {
		case 0:
			c.a >>= c.getVal(oprand)
		case 1:
			c.b ^= oprand
		case 2:
			c.b = c.getVal(oprand) & 0b111
		case 3:
			if c.a != 0 {
				c.ip = oprand - 2
			}
		case 4:
			c.b ^= c.c
		case 5:
			out = append(out, c.getVal(oprand)&0b111)
		case 6:
			c.b = c.a >> c.getVal(oprand)
		case 7:
			c.c = c.a >> c.getVal(oprand)
		}
		c.ip += 2
	}
	return out
}

func compiledRun(a uint64) (uint64, int) {
	out := uint64(0)
	count := 0
	var b, c uint64
	for ; a != 0; a >>= 3 {
		b = a & 0b111
		b ^= 0b001
		c = a >> b
		out <<= 3
		out |= ((b ^ c ^ 0b110) & 0b111)
		count++
	}
	return out, count
}

func fromOutput(a, out uint64) []uint64 {
	guesses := []uint64{}
	for guess := range uint64(8) {
		guessA := a<<3 | guess
		b := guess ^ 0b001
		c := guessA >> b
		o := (b ^ c ^ 0b110) & 0b111
		if o == out {
			guesses = append(guesses, guessA)
		}
	}
	return guesses
}
