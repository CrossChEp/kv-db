package client

import (
	"errors"
	"fmt"
	"io"
)

func (c *Client) Send(query string) (string, error) {
	if _, err := c.conn.Write([]byte(query)); err != nil {
		return "", fmt.Errorf("conn.Write: %w", err)
	}

	resp := make([]byte, c.bufferSize)
	count, err := c.conn.Read(resp)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return "", nil
		}

		return "", fmt.Errorf("conn.Read: %w", err)
	}
	if count == c.bufferSize {
		return "", fmt.Errorf("buffer is too small")
	}

	return string(resp[:count]), nil
}
