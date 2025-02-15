package server

import (
	"time"
)

type Option func(opt *options)

type options struct {
	bufferSize  int
	idleTimeout time.Duration
}

func WithBufferSize(bs int) Option {
	return func(opt *options) {
		opt.bufferSize = bs
	}
}

func WithIdleTimeout(it time.Duration) Option {
	return func(opt *options) {
		opt.idleTimeout = it
	}
}
