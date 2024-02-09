package startup

import (
	"fmt"
	"gateway/internal/server/grpc"

	//"github.com/StasikLeyshin/grpc-kafka-services/internal/server/grpc"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Grpc        grpc.Config `yaml:"grpc"`
	GrpcGateway grpc.Config `yaml:"grpc_gateway"`
	//GrpcConfig GrpcConfig `yaml:"grpc"`
	//KafkaConfig KafkaConfig    `yaml:"kafka"`
	//Database    DatabaseConfig `yaml:"database"`
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

//func (c *Config) Reconfiguration(configPath string) (*Config, error) {
//	config, err := NewConfig(configPath)
//	if err != nil {
//		return nil, err
//	}
//
//	c = config
//}
