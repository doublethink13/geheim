package config

import (
	"flag"
	"os"
	"testing"
	"treuzedev/geheim/packages/shared"
)

func TestCheckFlag(t *testing.T) {
	testCases := getCheckFlagTestCases()
	for _, test := range testCases {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			setupCheckFlagEnv(test)

			config := Get()
			if !compareConfigs(config, test.expected) {
				t.Logf("Expected: %v, got: %v", test.expected, config)
				t.Fail()
			}
		})
	}
}

func setupCheckFlagEnv(test CheckFlagTestCase) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	args := make([]string, 0)

	if !shared.CompareStringSlices(test.flags.check, []string{}) {
		args = append(args, test.flags.check...)
	}

	if !shared.CompareStringSlices(test.flags.secretKey, []string{}) {
		args = append(args, test.flags.secretKey...)
	}

	os.Args = append([]string{os.Args[0]}, args...)
}

func getCheckFlagTestCases() []CheckFlagTestCase { // nolint
	return []CheckFlagTestCase{
		{
			name: "short flag + short option (encrypted file)",
			flags: flags{
				check:     []string{"-c", "e"},
				secretKey: []string{},
				encrypt:   []string{},
				decrypt:   []string{},
			},
			expected: Config{
				Check:     "e",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name: "short flag + long option (encrypted file)",
			flags: flags{
				check:     []string{"-c", "encrypted"},
				secretKey: []string{},
				encrypt:   []string{},
				decrypt:   []string{},
			},
			expected: Config{
				Check:     "encrypted",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name: "long flag + short option (encrypted file)",
			flags: flags{
				check:     []string{"--check", "e"},
				secretKey: []string{},
				encrypt:   []string{},
				decrypt:   []string{},
			},
			expected: Config{
				Check:     "e",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name: "long flag + long option (encrypted file)",
			flags: flags{
				check:     []string{"--check", "encrypted"},
				secretKey: []string{},
				encrypt:   []string{},
				decrypt:   []string{},
			},
			expected: Config{
				Check:     "encrypted",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name: "short flag + short option (decrypted file)",
			flags: flags{
				check:     []string{"-c", "d"},
				secretKey: []string{},
				encrypt:   []string{},
				decrypt:   []string{},
			},
			expected: Config{
				Check:     "d",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name: "short flag + long option (decrypted file)",
			flags: flags{
				check:     []string{"-c", "decrypted"},
				secretKey: []string{},
				encrypt:   []string{},
				decrypt:   []string{},
			},
			expected: Config{
				Check:     "decrypted",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name: "long flag + short option (decrypted file)",
			flags: flags{
				check:     []string{"--check", "d"},
				secretKey: []string{},
				encrypt:   []string{},
				decrypt:   []string{},
			},
			expected: Config{
				Check:     "d",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name: "long flag + long option (decrypted file)",
			flags: flags{
				check:     []string{"--check", "decrypted"},
				secretKey: []string{},
				encrypt:   []string{},
				decrypt:   []string{},
			},
			expected: Config{
				Check:     "decrypted",
				SecretKey: "",
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
}
