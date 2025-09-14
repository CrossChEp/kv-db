package compute

import "github.com/CrossChEp/kv-db/internal"

type Parser struct {
	log internal.Logger
}

func New(log internal.Logger) *Parser {
	return &Parser{log: log}
}
