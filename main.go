package main

import (
	"adventofcode/day5"
	"os"
)

func main() {
	args := os.Args
	data, err := os.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	input := string(data)
	// switch args[2] {
	// case "basic":
	// 	day3.MulBasicHandler(input)
	// case "advanced":
	// 	day3.MulAdvancedHandler(input)
	// }
	err = day5.BasicHandler(input)
	if err != nil {
		panic(err)
	}
}
