package writer

import "github.com/hueristiq/hqgologger/levels"

type Writer interface {
	Write(data []byte, level levels.LevelInt)
}
