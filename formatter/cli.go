package formatter

import (
	"bytes"

	"github.com/hueristiq/hqgologger/levels"
	"github.com/logrusorgru/aurora/v3"
)

type CLIFomartter struct{}

type CLIFomartterOptions struct {
	Colorize bool
}

var (
	_ Formatter = &CLIFomartter{}

	au aurora.Aurora = aurora.NewAurora(false)
)

func NewCLIFomartter(options *CLIFomartterOptions) *CLIFomartter {
	au = aurora.NewAurora(options.Colorize)

	return &CLIFomartter{}
}

func (c *CLIFomartter) Format(event *Log) ([]byte, error) {
	c.colorizeLabel(event)

	buffer := &bytes.Buffer{}
	buffer.Grow(len(event.Message))

	label, ok := event.Metadata["label"]
	if label != "" && ok {
		buffer.WriteRune('[')
		buffer.WriteString(label)
		buffer.WriteRune(']')
		buffer.WriteRune(' ')
		delete(event.Metadata, "label")
	}
	buffer.WriteString(event.Message)

	for k, v := range event.Metadata {
		buffer.WriteRune(' ')
		buffer.WriteString(c.colorizeKey(k))
		buffer.WriteRune('=')
		buffer.WriteString(v)
	}
	data := buffer.Bytes()
	return data, nil
}

func (c *CLIFomartter) colorizeKey(key string) string {
	return au.Bold(key).String()
}

func (c *CLIFomartter) colorizeLabel(event *Log) {
	label := event.Metadata["label"]

	if label == "" {
		return
	}

	switch event.Level {
	case levels.Levels[levels.LevelSilent]:
		return
	case levels.Levels[levels.LevelFatal]:
		event.Metadata["label"] = au.BrightRed(label).Bold().String()
	case levels.Levels[levels.LevelError]:
		event.Metadata["label"] = au.BrightRed(label).Bold().String()
	case levels.Levels[levels.LevelWarning]:
		event.Metadata["label"] = au.BrightYellow(label).Bold().String()
	case levels.Levels[levels.LevelInfo]:
		event.Metadata["label"] = au.BrightBlue(label).Bold().String()
	case levels.Levels[levels.LevelDebug]:
		event.Metadata["label"] = au.BrightMagenta(label).Bold().String()
	}
}
