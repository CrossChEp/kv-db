package config

import (
	"fmt"
	"github.com/CrossChEp/kv-db/internal/entity"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const (
	tmpCfgPath = "config.yml"
	customCfg  = `engine:
  type: "in_memory"
network:
  address: "localhost:3000"
  max_connections: 12
  max_message_size: "4KB"
  idle_timeout: "5m"
logging:
  level: "info"
  output: "../../logs/logs.log"`
)

func TestNew(t *testing.T) {
	tt := []struct {
		name       string
		cfgPath    string
		out        *Config
		wantErr    error
		withConfig bool
	}{
		{
			name: "success_without_flags_without_config",
			out: &Config{
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
			},
			wantErr: nil,
			cfgPath: defaultCfgPath,
		},
		{
			name: "success_without_flags_with_custom_cfg",
			out: &Config{
				Engine: Engine{
					Type: defaultEngineType,
				},
				Network: Network{
					Address:        "localhost:3000",
					MaxConnections: 12,
					MaxMessageSize: defaultMaxMessageSize,
					IdleTimeout:    defaultIdleTimeout,
				},
				Logging: Logging{
					Level:  entity.Info,
					Output: "../../logs/logs.log",
				},
			},
			wantErr:    nil,
			cfgPath:    tmpCfgPath,
			withConfig: true,
		},
	}

	t.Parallel()

	for _, tc := range tt {
		fmt.Println(os.Getwd())

		t.Run(tc.name, func(t *testing.T) {
			if tc.withConfig {
				file := setupTmpFile(t)
				defer func() {
					if file == nil {
						return
					}

					os.Remove(file.Name())
				}()
			}

			os.Setenv(configPathKey, tc.cfgPath)

			res, err := New()

			require.ErrorIs(t, err, tc.wantErr)
			require.Equal(t, tc.out, res)
		})
	}
}

func setupTmpFile(t *testing.T) *os.File {
	tmpFile, err := os.Create(tmpCfgPath)
	require.NoError(t, err)

	fmt.Println(os.Getwd())

	_, err = tmpFile.Write([]byte(customCfg))
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	return tmpFile
}
