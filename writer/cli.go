package writer

import (
	"os"
	"sync"

	"github.com/hueristiq/hqgologger/levels"
)

type CLIWriter struct{}

var (
	_ Writer = &CLIWriter{}

	mutex *sync.Mutex
)

func NewCLIWriter() *CLIWriter {
	mutex = &sync.Mutex{}

	return &CLIWriter{}
}

func (w *CLIWriter) Write(data []byte, level levels.Level) {
	mutex.Lock()
	defer mutex.Unlock()

	switch level {
	case levels.LevelError, levels.LevelFatal:
		os.Stderr.Write(data)
		os.Stderr.Write([]byte("\n"))
	default:
		os.Stdout.Write(data)
		os.Stdout.Write([]byte("\n"))
	}
}
