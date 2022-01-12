package config

import "testing"

type flags struct {
	check     []string
	secretKey []string
	encrypt   []string
	decrypt   []string
}

func checkConfig(t *testing.T, got *Config, expected Config) {
	r := recover()
	if r != nil {
		if !compareConfigs(Config{}, expected) {
			t.Logf("Expected: %v, got: panic", expected)
			t.Fail()
		}
	} else {
		if !compareConfigs(Config{}, expected) {
			if !compareConfigs(*got, expected) {
				t.Logf("Expected: %v, got: %v", expected, *got)
				t.Fail()
			}
		}
	}
}
