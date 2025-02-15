package main

import (
	"context"

	"github.com/CrossChEp/kv-db/internal/common"
	"github.com/CrossChEp/kv-db/internal/compute"
	"github.com/CrossChEp/kv-db/internal/config"
	"github.com/CrossChEp/kv-db/internal/database"
	"github.com/CrossChEp/kv-db/internal/logger"
	"github.com/CrossChEp/kv-db/internal/network/server"
	"github.com/CrossChEp/kv-db/internal/storage/engine"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	logOutput, file, err := logger.WithFileOutput(cfg.Logging.Output)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log := logger.New(
		logger.WithLogFormat(logger.TextFormat),
		logger.WithOutput(logOutput),
	)

	engine := engine.New()
	parser := compute.New(log)

	db := database.New(log, parser, engine)

	idleTimout, err := common.ParseIdleTimeout(cfg.IdleTimeout)
	if err != nil {
		panic(err)
	}

	messageSize, err := common.ParseBufferSize(cfg.MaxMessageSize)
	if err != nil {
		panic(err)
	}

	s, err := server.New(
		cfg,
		log,
		server.WithIdleTimeout(idleTimout),
		server.WithBufferSize(messageSize),
	)
	if err != nil {
		panic(err)
	}

	s.Run(ctx, db.Handle)
}
