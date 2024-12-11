package day1

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"adventofcode/dayhandler"
)

func init() {
	dayhandler.DayHandlers["day1"] = newDayHandler()
}

type dayHandler struct {
}

// AdvancedHandler implements dayhandler.dayHandler.
func (d dayHandler) AdvancedHandler(s string) error {
	leftPairs, rightPairs, err := inputHandler(s)
	if err != nil {
		return err
	}
	slices.Sort(leftPairs)
	slices.Sort(rightPairs)

	similarity := 0
	for i := 0; i < len(leftPairs); i++ {
		count := 0
		for j := 0; j < len(rightPairs); j++ {
			if rightPairs[j] == leftPairs[i] {
				count++
				continue
			}
		}
		similarity += count * leftPairs[i]
	}
	fmt.Printf("Similarity: %d", similarity)
	return nil
}

// BasicHandler implements dayhandler.dayHandler.
func (d dayHandler) BasicHandler(s string) error {
	leftPairs, rightPairs, err := inputHandler(s)
	if err != nil {
		return err
	}
	slices.Sort(leftPairs)
	slices.Sort(rightPairs)

	totalDistance := 0
	for i, l := range leftPairs {
		totalDistance += int(math.Abs(float64(l - rightPairs[i])))
	}
	fmt.Printf("TotalDistance: %d", totalDistance)
	return nil
}

func inputHandler(s string) (leftPairs, rightPairs []int, err error) {
	pairs := strings.Split(s, "\n")
	for _, p := range pairs {
		pair := strings.Split(p, "   ")
		left, err := strconv.ParseInt(strings.TrimSpace(pair[0]), 10, 0)
		if err != nil {
			return nil, nil, err
		}
		right, err := strconv.ParseInt(strings.TrimSpace(pair[1]), 10, 0)
		if err != nil {
			return nil, nil, err
		}
		leftPairs = append(leftPairs, int(left))
		rightPairs = append(rightPairs, int(right))
	}
	return
}

func newDayHandler() dayHandler {
	return dayHandler{}
}
