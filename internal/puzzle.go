package internal

import (
	"log"
)

type Puzzle interface {
	Init(data []byte)
	Solve1() any
	Solve2() any
}

var puzzles = make(map[int]Puzzle)

func Register(day int, puzzle Puzzle) {
	if _, exist := puzzles[day]; exist {
		log.Fatalf("duplicate registration for puzzle %d\n", day)
	}
	puzzles[day] = puzzle
}

func Get(day int) Puzzle {
	puzzle, exist := puzzles[day]
	if !exist {
		log.Fatalf("no puzzle registered for day %d\n", day)
	}
	return puzzle
}
