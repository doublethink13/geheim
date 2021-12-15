package config

import "gopkg.in/yaml.v2"

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
	case len(a.Files) != len(b.Files):
		return false
	}
	for i := 0; i < len(a.Files); i++ {
		if a.Files[i] != b.Files[i] {
			return false
		}
	}
	return true
}
