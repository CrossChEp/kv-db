package client

import (
	"fmt"
	"net"
)

type Client struct {
	conn       net.Conn
	bufferSize int
}

func New(addr string, opts ...Option) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("net.Dial: %w", err)
	}

	var ops options
	for _, opt := range opts {
		opt(&ops)
	}

	cl := &Client{conn: conn}

	if ops.bufferSize != 0 {
		cl.bufferSize = ops.bufferSize
	}

	return cl, nil
}
