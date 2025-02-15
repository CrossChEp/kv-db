package entity

import (
	"errors"
	"time"
)

const (
	TypeInMemory = "in_memory"

	Byte     = 1
	Kilobyte = Byte * 1024
	Megabyte = Kilobyte * 1024
	Gigabyte = Megabyte * 1024
)

var (
	EngineTypes            = []string{TypeInMemory}
	SuffixTimeFormatToTime = map[string]time.Duration{
		"ms": time.Millisecond,
		"s":  time.Second,
		"m":  time.Minute,
		"h":  time.Hour,
	}
	SuffixSizeFormatToSize = map[string]int{
		"b":  Byte,
		"kb": Kilobyte,
		"mb": Megabyte,
		"gb": Gigabyte,
	}

	ErrInvalidEngineType = errors.New("invalid engine type")

	ErrInvalidTimeoutFormat    = errors.New("invalid timeout format")
	ErrInvalidTimeSuffixFormat = errors.New("invalid time suffix format")

	ErrInvalidBufferSizeFormat = errors.New("invalid buffer size format")
	ErrInvalidBufferSizeSuffix = errors.New("invalid buffer size suffix")
)
