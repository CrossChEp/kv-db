package config

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"slices"

	"gopkg.in/yaml.v3"

	"github.com/CrossChEp/kv-db/internal/entity"
	"github.com/CrossChEp/kv-db/internal/utils/pointer"
)

const (
	configPathKey = "CONFIG_PATH"

	defaultCfgPath        = "config/config.yaml"
	defaultLogOutput      = "logs/logs.log"
	defaultEngineType     = entity.TypeInMemory
	defaultAddress        = "localhost:5000"
	defaultMaxConnections = 100
	defaultMaxMessageSize = "4KB"
	defaultIdleTimeout    = "5m"
	defaultLogLevel       = entity.Debug
)

var (
	flagEngineType = flag.String(
		"engine-type",
		"",
		fmt.Sprintf("Engine type of database: %+v", entity.EngineTypes),
	)
	flagAddress        = flag.String("address", "", "host address of database")
	flagMaxConnections = flag.Int64("max-conns", 0, "max numbers of connections to database")
	flagMaxMessageSize = flag.String("max-message-size", "", "max size of message")
	flagIdleTimeout    = flag.String("idle-timeout", "", "idle timout")
	flagLogLevel       = flag.String("log-level", "", "level of db logs")
	flagOutput         = flag.String("output", "", "output directory of logs")
)

type Config struct {
	Engine  `yaml:"engine"`
	Network `yaml:"network"`
	Logging `yaml:"logging"`
}

type Engine struct {
	Type string `yaml:"type"`
}

type Network struct {
	Address        string `yaml:"address"`
	MaxConnections int64  `yaml:"max_connections"`
	MaxMessageSize string `yaml:"max_message_size"`
	IdleTimeout    string `yaml:"idle_timeout"`
}

type Logging struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

func New() (*Config, error) {
	godotenv.Load()

	cfg := getDefaultCfg()

	cfgPath := os.Getenv(configPathKey)
	if cfgPath == "" {
		cfgPath = defaultCfgPath
	}

	if err := fillFieldsByCfg(cfg, cfgPath); err != nil {
		return nil, fmt.Errorf("Config.FillWithConfig: %w", err)
	}

	if err := fillFieldsByFlags(cfg); err != nil {
		return nil, fmt.Errorf("Config.FillWithFlags: %w", err)
	}

	return cfg, nil
}

func getDefaultCfg() *Config {
	return &Config{
		Engine: Engine{
			Type: defaultEngineType,
		},
		Network: Network{
			Address:        defaultAddress,
			MaxConnections: defaultMaxConnections,
			MaxMessageSize: defaultMaxMessageSize,
			IdleTimeout:    defaultIdleTimeout,
		},
		Logging: Logging{
			Level:  defaultLogLevel,
			Output: defaultLogOutput,
		},
	}
}

func fillFieldsByCfg(cfg *Config, cfgPath string) error {
	cfgFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil
	}

	var tmpCfg Config

	if err = yaml.Unmarshal(cfgFile, &tmpCfg); err != nil {
		return fmt.Errorf("Config.Unmarshal: %w", err)
	}

	if tmpCfg.Address != "" {
		cfg.Address = tmpCfg.Address
	}

	if tmpCfg.Engine.Type != "" {
		et := tmpCfg.Engine.Type

		if !slices.Contains(entity.EngineTypes, et) {
			return fmt.Errorf("%w: allowed types: %+v", entity.ErrInvalidEngineType, entity.EngineTypes)
		}

		cfg.Engine.Type = et
	}

	if tmpCfg.MaxConnections != 0 {
		cfg.MaxConnections = tmpCfg.MaxConnections
	}

	if tmpCfg.MaxMessageSize != "" {
		cfg.MaxMessageSize = tmpCfg.MaxMessageSize
	}

	if tmpCfg.IdleTimeout != "" {
		cfg.IdleTimeout = tmpCfg.IdleTimeout
	}

	if tmpCfg.Logging.Level != "" {
		logLevel := tmpCfg.Logging.Level

		if _, ok := entity.StringLevelToLoggerLevel[logLevel]; !ok {
			return fmt.Errorf("%w: allowed levels: %+v", entity.ErrInvalidLogLevel, entity.LogLevels)
		}

		cfg.Logging.Level = logLevel
	}

	if tmpCfg.Logging.Output != "" {
		o := tmpCfg.Logging.Output

		if _, err = os.Stat(o); err != nil {
			return err
		}

		cfg.Logging.Output = o
	}

	return nil
}

func fillFieldsByFlags(cfg *Config) error {
	flag.Parse()

	if pointer.Deref(flagAddress) != "" {
		cfg.Address = *flagAddress
	}

	if pointer.Deref(flagEngineType) != "" {
		et := *flagEngineType

		if slices.Contains(entity.EngineTypes, et) {
			return fmt.Errorf("%w: allowed types: %+v", entity.ErrInvalidEngineType, entity.EngineTypes)
		}

		cfg.Engine.Type = et
	}

	if pointer.Deref(flagMaxConnections) != 0 {
		cfg.MaxConnections = *flagMaxConnections
	}

	if pointer.Deref(flagMaxMessageSize) != "" {
		ms := *flagMaxMessageSize

		if len(ms) != 2 {
			return fmt.Errorf("invalid max message size")
		}

		cfg.MaxMessageSize = ms
	}

	if pointer.Deref(flagIdleTimeout) != "" {
		cfg.IdleTimeout = *flagIdleTimeout
	}

	if pointer.Deref(flagLogLevel) != "" {
		logLevel := *flagLogLevel

		if _, ok := entity.StringLevelToLoggerLevel[logLevel]; !ok {
			return fmt.Errorf("%w: allowed levels: %+v", entity.ErrInvalidLogLevel, entity.LogLevels)
		}

		cfg.Logging.Level = logLevel
	}

	if pointer.Deref(flagOutput) != "" {
		o := *flagOutput

		if _, err := os.Stat(o); err != nil {
			return err
		}

		cfg.Logging.Output = o
	}

	return nil
}
