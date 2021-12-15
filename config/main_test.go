package config

import (
	"testing"
)

func TestLocalConfig(t *testing.T) {
	tests := []Config{{"imsosecret", []string{"secrets.geheim.yaml"}}}
	for _, test := range tests {
		config := Get()
		if !compareConfigs(config, test) {
			t.Errorf("Get() == %s, expected %s", config, test)
		}
	}
}
