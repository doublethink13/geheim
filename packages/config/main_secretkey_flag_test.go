package config

import (
	"flag"
	"os"
	"testing"
)

func TestSecretKeyFlag(t *testing.T) {
	for _, test := range testSecretKeyFlagCases {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		args := []string{
			test.flags.secretKey[0],
			test.flags.secretKey[1],
		}
		os.Args = append([]string{os.Args[0]}, args...)
		t.Run(test.name, func(t *testing.T) {
			var got Config
			defer checkConfig(t, &got, test.expected)
			// why is Get() panicking when secretkey is not set but recover is not catching it?
			got = Get()
		})
	}
}

var testSecretKeyFlagCases = []struct {
	name     string
	flags    flags
	expected Config
}{
	{
		name: "short flag",
		flags: flags{
			check:     []string{},
			secretKey: []string{"-k", "test1"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test1",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "long flag",
		flags: flags{
			check:     []string{},
			secretKey: []string{"--secretkey", "test2"},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{
			Check:     "",
			SecretKey: "test2",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
	{
		name: "not set",
		flags: flags{
			check:     []string{},
			secretKey: []string{"", ""},
			encrypt:   []string{},
			decrypt:   []string{},
		},
		expected: Config{},
	},
}
