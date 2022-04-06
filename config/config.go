package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Hub       Service       `yaml:"hub"`
	Keylogger LoggerService `yaml:"keylogger"`
}

// A service represents a specification for a specific service
type Service struct {
	Port int    `yaml:"port"`
	URL  string `yaml:"url"`
}

// LoggerService is a logger config
type LoggerService struct {
	Service    `yaml:"service,inline"`
	DeviceName string `yaml:"deviceName"`
}

// Load loads the config from the specified YAML config file
func Load(filePath string) (*Config, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(bytes, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
