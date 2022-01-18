package config

import (
	"flag"
	"os"
	"testing"
)

func TestEncryptFlag(t *testing.T) {
	for _, test := range testEncryptFlagCases {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		args := []string{
			test.flags.secretKey[0],
			test.flags.secretKey[1],
			test.flags.encrypt[0],
			test.flags.encrypt[1],
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

var testEncryptFlagCases = []struct {
	name     string
	flags    flags
	expected Config
}{
	{
		name: "short flag + no option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{"-e", ""},
			decrypt:   []string{},
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
		name: "short flag + bool option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{"-e", "true"},
			decrypt:   []string{},
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
		name: "short flag + not bool option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{"-e", "imnotabool"},
			decrypt:   []string{},
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
		name: "long flag + no option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{"--encrypt", ""},
			decrypt:   []string{},
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
		name: "short flag + bool option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{"--encrypt", "true"},
			decrypt:   []string{},
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
		name: "long flag + not bool option",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{"--encrypt", "imnotabool"},
			decrypt:   []string{},
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
		name: "short flag + false",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{"-e=false", ""},
			decrypt:   []string{},
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
			encrypt:   []string{"--encrypt=false", ""},
			decrypt:   []string{},
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
			encrypt:   []string{"", ""},
			decrypt:   []string{},
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
