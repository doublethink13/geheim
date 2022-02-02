package config

//nolint
import (
	"fmt"
	"treuzedev/geheim/packages/shared"

	"gopkg.in/yaml.v2"
)

type Reader func(filename string) ([]byte, error)

type Config struct {
	Check     string
	SecretKey string   `yaml:"secretkey"`
	Files     []string `yaml:"files"`
	Encrypt   bool
	Decrypt   bool
}

func (c *Config) Parse(data []byte) error {
	if err := yaml.Unmarshal(data, c); err != nil {
		return fmt.Errorf("error parsing yaml: %w", err)
	}

	return nil
}

func CompareConfigs(fileA Config, fileB Config) bool {
	switch {
	case fileA.Check != fileB.Check:
		return false
	case fileA.SecretKey != fileB.SecretKey:
		return false
	case !shared.CompareStringSlices(fileA.Files, fileB.Files):
		return false
	case fileA.Encrypt != fileB.Encrypt:
		return false
	case fileA.Decrypt != fileB.Decrypt:
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
