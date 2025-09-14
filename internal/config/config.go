package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const defaultLogPath = "logs/logs.log"

type Config struct {
	LogPath string `envconfig:"LOG_PATH"`
}

func New() *Config {
	godotenv.Load()

	var cfg Config
	_ = envconfig.Process("", &cfg)

	if cfg.LogPath == "" {
		cfg.LogPath = defaultLogPath
	}

	return &cfg
}
