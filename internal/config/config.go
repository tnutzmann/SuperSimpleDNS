package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfigFromFile(filepath string) (cfg *Config, err error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("ERROR: unable to read file '%v': %w", filepath, err)
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("ERROR: unable to parse YAML: %w", err)
	}

	return cfg, nil
}
