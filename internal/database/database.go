package database

import "github.com/CrossChEp/kv-db/internal"

type Database struct {
	log    internal.Logger
	parser Parser
	engine Engine
}

func New(log internal.Logger, parser Parser, engine Engine) *Database {
	return &Database{
		log:    log,
		parser: parser,
		engine: engine,
	}
}
