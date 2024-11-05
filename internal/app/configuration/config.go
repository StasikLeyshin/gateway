package configuration

import (
	"fmt"
	"gateway/internal/repository/transfer/connector"
	"gateway/internal/server/grpc"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Grpc        grpc.Config      `yaml:"grpc"`
	GrpcGateway grpc.Config      `yaml:"grpc_gateway"`
	Connector   connector.Config `yaml:"route"`
}

func NewConfig(configPath string) (*Config, error) {
	rawYAML, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("reading file error: %w", err)
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(rawYAML, cfg); err != nil {
		return nil, fmt.Errorf("yaml parsing error: %w", err)
	}

	return cfg, nil
}

func (c *Config) ReConfigure(configPath string) (*Config, error) {
	config, err := NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	c = config

	return c, nil
}

func (c *Config) GetGrpcConfig() grpc.Config {
	return c.Grpc
}

func (c *Config) GetGrpcGatewayConfig() grpc.Config {
	return c.GrpcGateway
}

func (c *Config) GetConnectorConfig() *connector.Config {
	return &c.Connector
}
