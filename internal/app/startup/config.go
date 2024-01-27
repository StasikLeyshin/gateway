package startup

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"mir-service/internal/server/grpc-gateway"
	"os"
)

type Config struct {
	Store StoreConfig `yaml:"store"`
	Http  grpc.Config `yaml:"http"`
}

func NewConfig(configFile string) (*Config, error) {
	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("reading file error: %w", err)
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(rawYAML, cfg); err != nil {
		return nil, fmt.Errorf("yaml parsing error: %w", err)
	}

	return cfg, nil
}
