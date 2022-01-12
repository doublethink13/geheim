package config

import (
	"flag"
	"os"
	"testing"
)

func TestDecryptFlag(t *testing.T) {
	for _, test := range testDecryptFlagCases {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		args := []string{
			test.flags.secretKey[0],
			test.flags.secretKey[1],
			test.flags.decrypt[0],
			test.flags.decrypt[1],
		}
		os.Args = append([]string{os.Args[0]}, args...)
		t.Run(test.name, func(t *testing.T) {
			config := Get()
			if !compareConfigs(config, test.expected) {
				t.Logf("Expected: %v, got: %v", test.expected, config)
				t.Fail()
			}
		})
	}
}

var testDecryptFlagCases = []struct {
	name     string
	flags    flags
	expected Config
}{
	{
		name: "short flag + no option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{"-d", ""},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test",
			Encrypt:   false,
			Decrypt:   true,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "short flag + bool option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{"-d", "true"},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test",
			Encrypt:   false,
			Decrypt:   true,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "short flag + not bool option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{"-d", "imnotabool"},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test",
			Encrypt:   false,
			Decrypt:   true,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "long flag + no option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{"--decrypt", ""},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test",
			Encrypt:   false,
			Decrypt:   true,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "short flag + bool option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{"--decrypt", ""},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test",
			Encrypt:   false,
			Decrypt:   true,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "long flag + not bool option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{"--decrypt", "imnotabool"},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test",
			Encrypt:   false,
			Decrypt:   true,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "short flag + false",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{"-d=false", ""},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "long flag + false",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{"--decrypt=false", ""},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "not set",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{"", ""},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
}
