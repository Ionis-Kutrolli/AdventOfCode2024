package main

import (
	"adventofcode/dayhandler"
	"os"

	_ "adventofcode/day1"
	_ "adventofcode/day2"
	_ "adventofcode/day3"
	_ "adventofcode/day5"
	// _ "adventofcode/day6"
	// _ "adventofcode/day7"
	// _ "adventofcode/day8"
	// _ "adventofcode/day9"
	// _ "adventofcode/day10"
	// _ "adventofcode/day11"
	// _ "adventofcode/day12"
	// _ "adventofcode/day13"
	// _ "adventofcode/day14"
	// _ "adventofcode/day15"
	// _ "adventofcode/day16"
	// _ "adventofcode/day17"
	// _ "adventofcode/day18"
	// _ "adventofcode/day19"
	// _ "adventofcode/day20"
	// _ "adventofcode/day21"
	// _ "adventofcode/day22"
	// _ "adventofcode/day23"
	// _ "adventofcode/day24"
	// _ "adventofcode/day25"
)

func main() {
	args := os.Args
	data, err := os.ReadFile(args[3])
	if err != nil {
		panic(err)
	}

	input := string(data)
	switch args[2] {
	case "basic":
		err = dayhandler.DayHandlers[args[1]].BasicHandler(input)
	case "advanced":
		err = dayhandler.DayHandlers[args[1]].AdvancedHandler(input)
	}
	if err != nil {
		panic(err)
	}
}
