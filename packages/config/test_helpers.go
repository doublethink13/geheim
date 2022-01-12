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

var config1 = `---
secretkey: 'test1'
`

var config2 = `---
secretkey: 2
`

var config3 = `---
secretkey: ''
`

var config4 = `---
files: []
`

var config5 = `---
files:
  - testfile1
  - testfile2
`

var config6 = `---
files: 'thisiswrong'
`

var config7 = `---
secretkey: 'hello'
files:
  - testfile3
  - testfile4
`
