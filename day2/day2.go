package day2

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"adventofcode/dayhandler"
)

func init() {
	dayhandler.DayHandlers["day2"] = newDayHandler()
}

type dayHandler struct {
}

func newDayHandler() dayHandler {
	return dayHandler{}
}

// AdvancedHandler implements dayhandler.dayHandler.
func (d dayHandler) AdvancedHandler(s string) error {
	lines, err := inputHandler(s)
	if err != nil {
		return err
	}
	numSafe := 0
	for _, line := range lines {
		for i := -1; i < len(line); i++ {
			clone := line
			if i != -1 {
				clone = slices.Clone(line)
				clone = slices.Delete(clone, i, i+1)
			}
			increasing := 0.0
			safe := slices.IsSortedFunc(clone, func(a, b int) int {
				diff := float64(a - b)
				absDiff := math.Abs(diff)
				if absDiff < 1 || absDiff > 3 {
					return -1
				}
				if increasing == 0 {
					increasing = math.Copysign(1, diff)
					return 1
				} else {
					return int(math.Copysign(1, diff*increasing))
				}
			})
			if safe {
				numSafe++
				break
			}
		}
	}
	fmt.Printf("Number Safe: %d", numSafe)
	return nil
}

// BasicHandler implements dayhandler.dayHandler.
func (d dayHandler) BasicHandler(s string) error {
	lines, err := inputHandler(s)
	if err != nil {
		return err
	}
	numSafe := 0
	for _, line := range lines {
		increasing := 0.0
		safe := slices.IsSortedFunc(line, func(a, b int) int {
			diff := float64(a - b)
			absDiff := math.Abs(diff)
			if absDiff < 1 || absDiff > 3 {
				return -1
			}
			if increasing == 0 {
				increasing = math.Copysign(1, diff)
				return 1
			} else {
				return int(math.Copysign(1, diff*increasing))
			}
		})
		if safe {
			numSafe++
		}
	}
	fmt.Printf("Number Safe: %d", numSafe)
	return nil
}

func inputHandler(s string) ([][]int, error) {
	lines := strings.Split(s, "\n")
	parsedLines := [][]int{}
	for i, line := range lines {
		parsedLines = append(parsedLines, []int{})
		for _, num := range strings.Split(line, " ") {
			val, err := strconv.ParseInt(strings.TrimSpace(num), 10, 0)
			if err != nil {
				return nil, fmt.Errorf("failed to parse int %s: %w", num, err)
			}
			parsedLines[i] = append(parsedLines[i], int(val))
		}
	}
	return parsedLines, nil
}
