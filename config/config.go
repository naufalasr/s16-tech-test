package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		APIKey string `yaml:"api_key"`
	} `yaml:"server"`
	Omdb struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"omdb"`
}

var ConfigInUse Config

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&ConfigInUse); err != nil {
		return nil, err
	}

	return &ConfigInUse, nil
}
