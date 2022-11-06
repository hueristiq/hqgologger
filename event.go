package hqgologger

import (
	"fmt"

	"github.com/hueristiq/hqgologger/levels"
)

// Event is a log event to be written with data
type Event struct {
	level    levels.LevelInt
	message  string
	metadata map[string]string

	logger *Logger
}

// Label applies a custom label on the log event
func (event *Event) Label(label string) *Event {
	event.metadata["label"] = label

	return event
}

// Rest applies a custom label on the log event
func (event *Event) Rest(character string) *Event {
	event.metadata["rest"] = character

	return event
}

// // Str adds a string metadata item to the log
// func (event *Event) Str(key, value string) *Event {
// 	event.metadata[key] = value

// 	return event
// }

// Msg logs a message to the logger
func (event *Event) Msg(message string) {
	event.message = message

	event.logger.Log(event)
}

// Msgf logs a printf style message to the logger
func (event *Event) Msgf(format string, args ...interface{}) {
	event.message = fmt.Sprintf(format, args...)

	event.logger.Log(event)
}
