package hqgologger

import (
	"os"
	"strings"

	"github.com/hueristiq/hqgologger/formatter"
	"github.com/hueristiq/hqgologger/levels"
	"github.com/hueristiq/hqgologger/writer"
)

// Logger is the logger for logging structured data.
type Logger struct {
	formatter formatter.Formatter
	writer    writer.Writer
	maxLevel  levels.LevelInt
}

// SetMaxLevel sets the Logger's max logging level
func (logger *Logger) SetMaxLevel(level levels.LevelStr) {
	logger.maxLevel = levels.Levels[level]
}

// SetFormatter sets the Logger's formatter
func (logger *Logger) SetFormatter(formatter formatter.Formatter) {
	logger.formatter = formatter
}

// SetWriter sets the Logger's writer
func (logger *Logger) SetWriter(writer writer.Writer) {
	logger.writer = writer
}

// Log logs an Event
func (logger *Logger) Log(event *Event) {
	if event.level > event.logger.maxLevel {
		return
	}

	if label, ok := event.metadata["label"]; !ok {
		labels := map[levels.LevelInt]string{
			levels.Levels[levels.LevelFatal]:   "FTL",
			levels.Levels[levels.LevelError]:   "ERR",
			levels.Levels[levels.LevelWarning]: "WRN",
			levels.Levels[levels.LevelInfo]:    "INF",
			levels.Levels[levels.LevelDebug]:   "DBG",
		}

		if label, ok = labels[event.level]; ok {
			event.metadata["label"] = label
		}
	}

	event.message = strings.TrimSuffix(event.message, "\n")

	data, err := logger.formatter.Format(&formatter.Log{
		Message:  event.message,
		Level:    event.level,
		Metadata: event.metadata,
	})
	if err != nil {
		return
	}

	if character, ok := event.metadata["rest"]; ok {
		data = appendRest(data, character)
	}

	logger.writer.Write(data, event.level)

	if event.level == levels.Levels[levels.LevelFatal] {
		os.Exit(1)
	}
}

// Print prints a string on screen without any extra labels.
func (logger *Logger) Print() *Event {
	event := &Event{
		logger:   logger,
		level:    levels.LevelInt(-1),
		metadata: make(map[string]string),
	}

	return event
}

// Debug writes an error message on the screen with the default label
func (logger *Logger) Debug() *Event {
	level := levels.Levels[levels.LevelDebug]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

// Info writes a info message on the screen with the default label
func (logger *Logger) Info() *Event {
	level := levels.Levels[levels.LevelInfo]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

// Warning writes a warning message on the screen with the default label
func (logger *Logger) Warning() *Event {
	level := levels.Levels[levels.LevelWarning]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

// Error writes a error message on the screen with the default label
func (logger *Logger) Error() *Event {
	level := levels.Levels[levels.LevelError]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

// Fatal exits the program if we encounter a fatal error
func (logger *Logger) Fatal() *Event {
	level := levels.Levels[levels.LevelFatal]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}
