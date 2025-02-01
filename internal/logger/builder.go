package logger

import (
	"io"

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
