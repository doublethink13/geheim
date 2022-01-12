package config

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestFileConfig(t *testing.T) {
	setup()
	for _, test := range testFileConfigCases {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		args := []string{
			"-k",
			"",
		}
		os.Args = append([]string{os.Args[0]}, args...)
		reader = func(filename string) ([]byte, error) {
			return test.data, test.err
		}
		config := Get()
		fmt.Println(config)
	}
}

func setup() {
	getConfigLocation = func() (location string) { return "test" }
}

var testFileConfigCases = []struct {
	name     string
	data     []byte
	err      error
	expected Config
}{
	{
		name: "correct secretkey, files not set",
		data: []byte(config1),
		err:  nil,
		expected: Config{
			Check:     "",
			SecretKey: "test1",
			Encrypt:   true,
			Decrypt:   false,
			Files:     []string{"secrets.geheim.yaml"},
		},
	},
}
