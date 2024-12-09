package dayhandler

var (
	DayHandlers map[string]dayHandler
)

type dayHandler interface {
	BasicHandler(string) error
	AdvancedHandler(string) error
}

func init() {
	DayHandlers = map[string]dayHandler{}
}
