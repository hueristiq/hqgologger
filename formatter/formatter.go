package formatter

import "github.com/hueristiq/hqgologger/levels"

type Log struct {
	Message  string
	Level    levels.LevelInt
	Metadata map[string]string
}

type Formatter interface {
	Format(log *Log) ([]byte, error)
}
