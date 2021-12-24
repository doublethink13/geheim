package config

import (
	"treuzedev/geheim/shared"

	"gopkg.in/yaml.v2"
)

// TODO: why the string?
type Config struct {
	SecretKey string   `yaml:"secretKey"`
	Files     []string `yaml:"files"`
}

func (c *Config) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func compareConfigs(a Config, b Config) bool {
	switch {
	case a.SecretKey != b.SecretKey:
		return false
	case !shared.CompareStringSlices(a.Files, b.Files):
		return false
	default:
		return true
	}
}

type CliFlags struct {
	SecretKey string
}
