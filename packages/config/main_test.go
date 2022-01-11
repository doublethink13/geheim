package config

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

type flags struct {
	check     []string
	secretKey []string
	encrypted string
	decrypted string
}

var testCheckEncryptedCases = []struct {
	name     string
	flags    flags
	expected Config
}{
	{
		name: "short flag + short option",
		flags: flags{
			check:     []string{"-c", "e"},
			secretKey: []string{"-k", "test"},
			encrypted: "",
			decrypted: "",
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
		name: "short flag + long option",
		flags: flags{
			check:     []string{"-c", "encrypted"},
			secretKey: []string{"-k", "test"},
			encrypted: "",
			decrypted: "",
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
		name: "long flag + short option",
		flags: flags{
			check:     []string{"--check", "e"},
			secretKey: []string{"-k", "test"},
			encrypted: "",
			decrypted: "",
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
		name: "long flag + long option",
		flags: flags{
			check:     []string{"--check", "encrypted"},
			secretKey: []string{"-k", "test"},
			encrypted: "",
			decrypted: "",
		},
		expected: Config{
			Check:     "encrypted",
			SecretKey: "test",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
}

func TestCheckEncrypted(t *testing.T) {
	for _, test := range testCheckEncryptedCases {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		args := []string{
			test.flags.check[0],
			test.flags.check[1],
			test.flags.secretKey[0],
			test.flags.secretKey[1],
		}
		os.Args = append([]string{os.Args[0]}, args...)
		t.Run(test.name, func(t *testing.T) {
			config := Get()
			fmt.Println(config)
			if !compareConfigs(config, test.expected) {
				t.Logf("Expected: %v, got: %v", test.expected, config)
				t.Fail()
			}
		})
	}
}

// func init() {
// 	_, filename, _, _ := runtime.Caller(0)
// 	dir := path.Join(path.Dir(filename), "../..")
// 	err := os.Chdir(dir)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func TestLocalConfig(t *testing.T) {
// 	tests := []Config{{"", "imsosecret", []string{
// 		"testfiles/config.json",
// 		"testfiles/coverage_testfile.xml",
// 		"testfiles/helpers.sh",
// 		"testfiles/id_rsa",
// 		"testfiles/id_rsa.pub",
// 		"testfiles/known_hosts",
// 		"testfiles/secrets_test.geheim.yaml",
// 		"testfiles/simple.txt",
// 		"testfiles/supervisor_env",
// 	}, true, false}}
// 	for _, test := range tests {
// 		config := Get()
// 		if !compareConfigs(config, test) {
// 			t.Errorf("Get() == %v, expected %v", config, test)
// 		}
// 	}
// }
