package day2

import (
	"adventofcode/dayhandler"
	"fmt"
)

func init() {
	dayhandler.DayHandlers["day2"] = newDayHandler()
}

type dayHandler struct {
}

// AdvancedHandler implements dayhandler.dayHandler.
func (d dayHandler) AdvancedHandler(string) error {
	return fmt.Errorf("unimplemented")}

// BasicHandler implements dayhandler.dayHandler.
func (d dayHandler) BasicHandler(string) error {
	return fmt.Errorf("unimplemented")
}

func newDayHandler() dayHandler {
	return dayHandler{}
}