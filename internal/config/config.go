package config

import (
	"log"
	"os"

	"sigs.k8s.io/yaml"
)

const (
	CONFIG_PATH = "./internal/config/config.yml"
)

type Config struct {
	IsDefault bool `json:"-"`
	Log       struct {
		Level string `json:"level"`
	} `json:"log"`

	File struct {
		Format string `json:"format"`
		Path   string `json:"path"`
	} `json:"file"`
}

func (cfg *Config) Save() {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatalf("failed to marshal config to YAML: %v", err)
	}

	if err := os.WriteFile(CONFIG_PATH, data, 0600); err != nil {
		return
	}
}

func MustLoad() *Config {
	var cfg Config

	data, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		cfg = Config{
			IsDefault: true,
			Log: struct {
				Level string `json:"level"`
			}{
				Level: "info",
			},
			File: struct {
				Format string `json:"format"`
				Path   string `json:"path"`
			}{
				Format: "yaml",
				Path:   "./data",
			},
		}
		return &cfg
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to parse config YAML: %v. Please check the file format", err)
	}

	return &cfg
}
