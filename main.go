package main

import (
	"os"

	"adventofcode/dayhandler"
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
	err = dayhandler.DayHandlers[args[1]].BasicHandler(input)
	if err != nil {
		panic(err)
	}
}
