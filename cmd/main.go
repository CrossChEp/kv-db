package main

import (
	"context"

	"github.com/CrossChEp/kv-db/internal/compute"
	"github.com/CrossChEp/kv-db/internal/config"
	"github.com/CrossChEp/kv-db/internal/database"
	"github.com/CrossChEp/kv-db/internal/logger"
	"github.com/CrossChEp/kv-db/internal/storage/engine"
)

func main() {
	cfg := config.New()
	ctx := context.Background()

	logOutput, file, err := logger.WithFileOutput(cfg.LogPath)
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

	db.Start(ctx)
}
