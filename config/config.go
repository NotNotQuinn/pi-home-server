package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Hub       Service `yaml:"hub"`
	Keylogger Service `yaml:"keylogger"`
}

// A service represents a specification for a specific service
type Service struct {
	Port int    `yaml:"port"`
	URL  string `yaml:"url"`
}

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
