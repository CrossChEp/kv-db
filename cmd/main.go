package main

import (
	"os"

	"github.com/CrossChEp/kv-db/internal/compute"
	"github.com/CrossChEp/kv-db/internal/config"
	"github.com/CrossChEp/kv-db/internal/database"
	"github.com/CrossChEp/kv-db/internal/logger"
	"github.com/CrossChEp/kv-db/internal/storage/engine"
)

func main() {
	cfg := config.New()

	file, err := os.OpenFile(cfg.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	log := logger.New(
		logger.WithLogFormat(logger.TextFormat),
		logger.WithOutput(file),
	)

	engine := engine.New()
	parser := compute.New(log)

	db := database.New(log, parser, engine)

	db.Start()
}
