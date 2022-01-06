package config

import (
	"treuzedev/geheim/packages/shared"

	"gopkg.in/yaml.v2"
)

// TODO: why the string?
type Config struct {
	Check     string
	SecretKey string   `yaml:"secretkey"`
	Files     []string `yaml:"files"`
	Encrypt   bool
	Decrypt   bool
}

func (c *Config) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func compareConfigs(a Config, b Config) bool {
	switch {
	case a.Check != b.Check:
		return false
	case a.SecretKey != b.SecretKey:
		return false
	case !shared.CompareStringSlices(a.Files, b.Files):
		return false
	case a.Encrypt != b.Encrypt:
		return false
	case a.Decrypt != b.Decrypt:
		return false
	default:
		return true
	}
}

type CliFlags struct {
	Check     string
	SecretKey string
	Encrypt   bool
	Decrypt   bool
}
