package entity

import (
	"errors"

	"github.com/CrossChEp/kv-db/internal/logger"
)

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
)

var (
	StringLevelToLoggerLevel = map[string]int{
		Debug: logger.DebugLevel,
		Info:  logger.InfoLevel,
		Warn:  logger.WarnLevel,
		Error: logger.ErrLevel,
	}
	LogLevels = []string{Debug, Info, Warn, Error}

	ErrInvalidLogLevel = errors.New("invalid log level provided")
)
