package hqgologger

import (
	"github.com/hueristiq/hqgologger/formatter"
	"github.com/hueristiq/hqgologger/levels"
	"github.com/hueristiq/hqgologger/writer"
)

var (
	DefaultLogger *Logger
)

func init() {
	DefaultLogger = &Logger{}
	DefaultLogger.SetMaxLevel(levels.LevelDebug)
	DefaultLogger.SetFormatter(formatter.NewCLI(&formatter.CLIOptions{
		Colorize: true,
	}))
	DefaultLogger.SetWriter(writer.NewCLI())
}

// Print prints a string on screen without any extra labels.
func Print() (event *Event) {
	event = &Event{
		logger:   DefaultLogger,
		level:    levels.LevelInt(-1),
		metadata: make(map[string]string),
	}

	return event
}

// Debug writes an error message on the screen with the default label
func Debug() (event *Event) {
	level := levels.Levels[levels.LevelDebug]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}

// Info writes a info message on the screen with the default label
func Info() (event *Event) {
	level := levels.Levels[levels.LevelInfo]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}

// Warning writes a warning message on the screen with the default label
func Warning() (event *Event) {
	level := levels.Levels[levels.LevelWarning]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}

// Error writes a error message on the screen with the default label
func Error() (event *Event) {
	level := levels.Levels[levels.LevelError]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}

// Fatal exits the program if we encounter a fatal error
func Fatal() (event *Event) {
	level := levels.Levels[levels.LevelFatal]

	event = &Event{
		logger:   DefaultLogger,
		level:    level,
		metadata: make(map[string]string),
	}

	return
}
