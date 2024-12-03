package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Xiangze-Li/advent-2024/internal"
	_ "github.com/Xiangze-Li/advent-2024/internal/registry"
	"github.com/Xiangze-Li/advent-2024/util"
)

func main() {
	var day, part int
	if len(os.Args) <= 1 {
		day = time.Now().Day()
	} else {
		day = util.Must(strconv.Atoi(os.Args[1]))
	}
	if len(os.Args) > 2 {
		part = util.Must(strconv.Atoi(os.Args[2]))
	} else {
		part = 3
	}

	puzzle := internal.Get(day)
	input := bytes.TrimSpace(input(day))

	output := strings.Builder{}

	{
		beginInit := time.Now()
		puzzle.Init(input)
		durInit := time.Since(beginInit)
		output.WriteString(fmt.Sprintf("\tDay\t%02d\n\tInit\t%v\n", day, durInit))
	}

	if part&1 != 0 {
		begin1 := time.Now()
		res1 := puzzle.Solve1()
		dur1 := time.Since(begin1)
		output.WriteString(fmt.Sprintf("\n\tPart\t1 (%v)\n%v\n", dur1, res1))
	}

	if part&2 != 0 {
		begin2 := time.Now()
		res2 := puzzle.Solve2()
		dur2 := time.Since(begin2)
		output.WriteString(fmt.Sprintf("\n\tPart\t2 (%v)\n%v\n", dur2, res2))
	}

	fmt.Fprint(os.Stdout, output.String())
}

func input(day int) []byte {
	return util.Must(os.ReadFile(fmt.Sprintf("%02d/input.txt", day)))
}
