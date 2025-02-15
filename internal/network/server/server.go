package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/CrossChEp/kv-db/internal"
	"github.com/CrossChEp/kv-db/internal/config"
)

type handleFunc func(ctx context.Context, query string) string

type Server struct {
	cfg            *config.Config
	listener       net.Listener
	log            internal.Logger
	idleTimeout    time.Duration
	maxMessageSize int
}

func New(cfg *config.Config, log internal.Logger, opts ...Option) (*Server, error) {
	listener, err := net.Listen("tcp", cfg.Network.Address)
	if err != nil {
		return nil, fmt.Errorf("net.Listen: %w", err)
	}

	var ops options
	for _, opt := range opts {
		opt(&ops)
	}

	if ops.idleTimeout == 0 {
		return nil, fmt.Errorf("could not get idle timeout")
	}

	if ops.bufferSize == 0 {
		return nil, fmt.Errorf("could not get buff size")
	}

	return &Server{
		cfg:            cfg,
		listener:       listener,
		log:            log,
		idleTimeout:    ops.idleTimeout,
		maxMessageSize: ops.bufferSize,
	}, nil
}
