package logger

import (
	"io"
	"os"

	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrLevel
)

const (
	JSONFormat = iota
	TextFormat
)

var (
	levelToZapLevel = map[int]zapcore.Level{
		DebugLevel: zapcore.DebugLevel,
		InfoLevel:  zapcore.InfoLevel,
		WarnLevel:  zapcore.WarnLevel,
		ErrLevel:   zapcore.ErrorLevel,
	}
)

type options struct {
	level   int
	format  int
	out     io.Writer
	adapter adapter
}

type Option func(opt *options)

func WithLogLevel(level int) Option {
	return func(opt *options) {
		opt.level = level
	}
}

func WithLogFormat(format int) Option {
	return func(opt *options) {
		opt.format = format
	}
}

func WithFileOutput(filePath string) (io.Writer, *os.File, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}

	return file, file, nil
}

func WithOutput(out io.Writer) Option {
	return func(opt *options) {
		opt.out = out
	}
}

func WithDriver(d adapter) Option {
	return func(opt *options) {
		opt.adapter = d
	}
}
