package levels

type Level int

const (
	LevelSilent Level = iota
	LevelFatal
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

func (level Level) String() string {
	return [...]string{"silent", "fatal", "error", "warning", "info", "debug"}[level]
}
