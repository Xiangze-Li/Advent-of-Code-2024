#!/usr/bin/env python3

import os
import sys

content_solve = """package puzzle{day_string}

import (
	"github.com/Xiangze-Li/advent-2024/internal"
	"github.com/Xiangze-Li/advent-2024/util"
)

type p struct {{
}}

func (p *p) Init(data []byte) {{
}}

func (p *p) Solve1() any {{
	return 0
}}

func (p *p) Solve2() any {{
	return 0
}}

func init() {{
	internal.Register({day}, &p{{}})
}}
"""

content_reg = """package registry

import _ "github.com/Xiangze-Li/advent-2024/{day_string}"
"""

def prepare(day : int) -> None :
  assert (1 <= day <= 25), "Day must be between 1 and 25"
  day_string: str = f"{day:02d}"
  os.makedirs(day_string, mode=0o755, exist_ok=True)
  os.system(f"aoc -y=2024 -I -i={day_string}/input.txt -d={day} -q download")
  with open(f"{day_string}/solution.go", "w") as f:
    f.write(content_solve.format(day_string=day_string, day=day))
  with open(f"internal/registry/{day_string}.go", "w") as f:
    f.write(content_reg.format(day_string=day_string))


if __name__ == "__main__":
  if len(sys.argv) != 2:
    print("Usage: python prepare.py <day>")
    sys.exit(1)
  day = int(sys.argv[1])
  prepare(day)
  print(f"Day {day} prepared.")
