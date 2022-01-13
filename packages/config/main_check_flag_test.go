package config

import (
	"flag"
	"os"
	"testing"
)

func TestCheckFlag(t *testing.T) {
	for _, test := range testCheckFlagCases {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		args := []string{
			test.flags.secretKey[0],
			test.flags.secretKey[1],
			test.flags.check[0],
			test.flags.check[1],
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

var testCheckFlagCases = []struct {
	name     string
	flags    flags
	expected Config
}{
	{
		name: "short flag + short option (encrypted file)",
		flags: flags{
			check:     []string{"-c", "e"},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "e",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "short flag + long option (encrypted file)",
		flags: flags{
			check:     []string{"-c", "encrypted"},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "encrypted",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "long flag + short option (encrypted file)",
		flags: flags{
			check:     []string{"--check", "e"},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "e",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "long flag + long option (encrypted file)",
		flags: flags{
			check:     []string{"--check", "encrypted"},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "encrypted",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "short flag + short option (decrypted file)",
		flags: flags{
			check:     []string{"-c", "d"},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "d",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "short flag + long option (decrypted file)",
		flags: flags{
			check:     []string{"-c", "decrypted"},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "decrypted",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "long flag + short option (decrypted file)",
		flags: flags{
			check:     []string{"--check", "d"},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "d",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "long flag + long option (decrypted file)",
		flags: flags{
			check:     []string{"--check", "decrypted"},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "decrypted",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "not set",
		flags: flags{
			check:     []string{"", ""},
			secretKey: []string{"-k", "test"},
			encrypt:   []string{},
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
