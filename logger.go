package hqgologger

import (
	"fmt"
	"os"
	"strings"

	"github.com/hueristiq/hqgologger/formatter"
	"github.com/hueristiq/hqgologger/levels"
	"github.com/hueristiq/hqgologger/writer"
)

var (
	DefaultLogger *Logger

	labels = map[levels.Level]string{
		levels.LevelFatal:   "FTL",
		levels.LevelError:   "ERR",
		levels.LevelWarning: "WRN",
		levels.LevelInfo:    "INF",
		levels.LevelDebug:   "DBG",
	}
)

func init() {
	DefaultLogger = &Logger{}
	DefaultLogger.SetMaxLevel(levels.LevelInfo)
	cli := formatter.NewCLIFomartter(&formatter.CLIFomartterOptions{
		Colorize: true,
	})
	DefaultLogger.SetFormatter(cli)
	DefaultLogger.SetWriter(writer.NewCLIWriter())
}

// Logger is the logger for logging structured data.
type Logger struct {
	maxLevel  levels.Level
	formatter formatter.Formatter
	writer    writer.Writer
}

// SetMaxLevel sets the max logging level for logger
func (logger *Logger) SetMaxLevel(level levels.Level) {
	logger.maxLevel = level
}

// SetFormatter sets the formatter instance for a logger
func (logger *Logger) SetFormatter(formatter formatter.Formatter) {
	logger.formatter = formatter
}

// SetWriter sets the writer instance for a logger
func (logger *Logger) SetWriter(writer writer.Writer) {
	logger.writer = writer
}

func (logger *Logger) Log(event *Event) {
	if !isCurrentLevelEnabled(event) {
		return
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
	logger.writer.Write(data, event.level)

	if event.level == levels.LevelFatal {
		os.Exit(1)
	}
}

// Print prints a string on screen without any extra labels.
func (logger *Logger) Print() *Event {
	level := levels.LevelDebug
	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

// Debug writes an error message on the screen with the default label
func (logger *Logger) Debug() *Event {
	level := levels.LevelDebug
	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return event
}

// Info writes a info message on the screen with the default label
func (logger *Logger) Info() *Event {
	level := levels.LevelInfo
	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return event
}

// Warning writes a warning message on the screen with the default label
func (logger *Logger) Warning() *Event {
	level := levels.LevelWarning
	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return event
}

// Error writes a error message on the screen with the default label
func (logger *Logger) Error() *Event {
	level := levels.LevelError
	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return event
}

// Fatal exits the program if we encounter a fatal error
func (logger *Logger) Fatal() *Event {
	level := levels.LevelFatal
	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return event
}

// Event is a log event to be written with data
type Event struct {
	logger   *Logger
	level    levels.Level
	message  string
	metadata map[string]string
}

// Label applies a custom label on the log event
func (event *Event) Label(label string) *Event {
	event.metadata["label"] = label
	return event
}

// Str adds a string metadata item to the log
func (event *Event) Str(key, value string) *Event {
	event.metadata[key] = value
	return event
}

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

// Print prints a string on screen without any extra labels.
func Print() (event *Event) {
	level := levels.LevelDebug
	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

// Debug writes an error message on the screen with the default label
func Debug() (event *Event) {
	level := levels.LevelDebug
	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return
}

// Info writes a info message on the screen with the default label
func Info() (event *Event) {
	level := levels.LevelInfo
	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return
}

// Warning writes a warning message on the screen with the default label
func Warning() (event *Event) {
	level := levels.LevelWarning
	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return
}

// Error writes a error message on the screen with the default label
func Error() (event *Event) {
	level := levels.LevelError
	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return
}

// Fatal exits the program if we encounter a fatal error
func Fatal() (event *Event) {
	level := levels.LevelFatal
	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}
	event.metadata["label"] = labels[level]

	return
}

func isCurrentLevelEnabled(event *Event) bool {
	return event.level <= event.logger.maxLevel
}
