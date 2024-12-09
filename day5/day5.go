package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"adventofcode/dayhandler"
)

func init() {
	dayhandler.DayHandlers["day5"] = newDayHandler()
}

type dayHandler struct {
}

func newDayHandler() dayHandler {
	return dayHandler{}
}

func (d dayHandler) BasicHandler(str string) error {
	strs := strings.Split(str, "\n\n")
	rules := map[string][]string{}
	for _, r := range strings.Split(strs[0], "\n") {
		pages := strings.Split(r, "|")
		// fmt.Printf("Pages: %v\n", pages)
		rules[pages[1]] = append(rules[pages[1]], pages[0])
		// fmt.Printf("%v\n", rules[pages[1]])
	}
	sum := 0
	for _, r := range strings.Split(strs[1], "\n") {
		correct := true
		printed := []string{}
		update := strings.Split(r, ",")
		validRules := filterUnneededVals(rules, update)
		for _, p := range update {
			val, ok := validRules[p]
			if !ok || len(val) == 0 {

				printed = append(printed, p)
				continue
			}
			if !containsAll(printed, val) {
				correct = false
				break
			}
			printed = append(printed, p)
		}
		if correct {
			val, err := strconv.ParseInt(printed[len(printed)/2], 10, 0)
			if err != nil {
				return err
			}
			sum += int(val)
		}
	}
	fmt.Println(sum)
	return nil
}

func containsAll(s, c []string) bool {
	for _, v := range c {
		if !slices.Contains(s, v) {
			return false
		}
	}
	return true
}

func filterUnneededVals(vals map[string][]string, needed []string) map[string][]string {
	fmt.Printf("Rules: %v\n", vals)
	neededVals := map[string][]string{}
	for _, v := range needed {
		slice, ok := vals[v]
		if !ok {
			continue
		}
		output := []string{}
		for _, n := range slice {
			if slices.Contains(needed, n) {
				output = append(output, n)
			}
		}
		neededVals[v] = output
	}
	return neededVals
}

func (d dayHandler) AdvancedHandler(str string) error {
	return fmt.Errorf("Unimplemented")
}
