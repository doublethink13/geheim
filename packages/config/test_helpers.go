package config

import "testing"

type flags struct {
	check     []string
	secretKey []string
	encrypt   []string
	decrypt   []string
}

type FileConfigTestCase struct {
	name     string
	data     []byte
	err      error
	expected Config
}

type CheckFlagTestCase struct {
	name     string
	flags    flags
	expected Config
}

func checkConfig(t *testing.T, got *Config, expected Config) {
	if r := recover(); r != nil && !compareConfigs(Config{}, expected) {
		t.Logf("Expected: %v, got: panic", expected)
		t.Fail()
	} else {
		if !compareConfigs(Config{}, expected) && !compareConfigs(*got, expected) {
			t.Logf("Expected: %v, got: %v", expected, *got)
			t.Fail()
		}
	}
}

func getConfig1() string {
	return `---
secretkey: 'test1'
`
}

func getConfig2() string {
	return `---
secretkey: 123456789
`
}

func getConfig3() string {
	return `---
secretkey: ''
`
}

func getConfig4() string {
	return `---
secretkey: 'test'
files: []
`
}

func getConfig5() string {
	return `---
secretkey: 'test'
files:
  - testfile1
  - testfile2
`
}

func getConfig6() string {
	return `---
secretkey: 'test'
files: 'thisiswrong'
`
}
