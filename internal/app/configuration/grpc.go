package configuration

type GrpcConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
