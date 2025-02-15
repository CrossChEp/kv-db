package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/CrossChEp/kv-db/internal/fields"
)

func (s *Server) Run(ctx context.Context, hf handleFunc) {
	fmt.Println("Server started!")

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			conn, err := s.listener.Accept()
			if err != nil {
				s.log.Error(s.log.WithFields(ctx,
					fields.Join(
						fields.WithAddress(s.cfg.Address),
						fields.WithError(err),
					)), "Could not accept connections")

				continue
			}

			go s.handleConn(ctx, conn, hf)
		}
	}()

	wg.Wait()
}

func (s *Server) handleConn(ctx context.Context, conn net.Conn, hf handleFunc) {
	defer func() {
		if p := recover(); p != nil {
			s.log.Error(
				s.log.WithFields(ctx, fields.WithPanic(p)),
				"handleConn",
			)
		}

		if err := conn.Close(); err != nil {
			s.log.Error(
				s.log.WithFields(ctx, fields.WithError(err)),
				"could not close connection",
			)
		}
	}()

	if err := conn.SetReadDeadline(time.Now().Add(s.idleTimeout)); err != nil {
		s.log.Error(s.log.WithFields(ctx,
			fields.Join(
				fields.WithError(err),
			)),
			"Could not set deadline to connections")

		return
	}

	msg := make([]byte, s.maxMessageSize)

	for {
		count, err := conn.Read(msg)
		if err != nil {
			if errors.Is(err, io.EOF) {
				s.log.Error(s.log.WithFields(
					ctx,
					fields.WithError(err),
				), "could not read query")
			}

			return
		}

		if count == s.maxMessageSize {
			s.log.Error(
				s.log.WithFields(ctx,
					fields.Join(
						fields.WithReadCount(count),
						fields.WithBufferSize(s.maxMessageSize),
					),
				),
				"buffer is too small",
			)

			return
		}

		query := msg[:count]

		resp := hf(ctx, string(query))

		if err = conn.SetWriteDeadline(time.Now().Add(s.idleTimeout)); err != nil {
			s.log.Error(s.log.WithFields(ctx,
				fields.Join(
					fields.WithError(err),
				)),
				"Could not set deadline to connections")

			return
		}

		if _, err = conn.Write([]byte(resp)); err != nil {
			s.log.Error(s.log.WithFields(
				ctx,
				fields.WithError(err),
			), "could not write response")

			return
		}
	}
}
