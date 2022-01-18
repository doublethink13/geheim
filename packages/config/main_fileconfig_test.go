package config

import (
	"flag"
	"os"
	"testing"
)

func TestFileConfig(t *testing.T) {
	setupCleanup(t)
	generalFileConfigSetup()

	testCases := getFileConfigTestCases()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			setupFileConfigCliFlags()
			setupFileConfigReader(test.data, test.err)

			var got Config
			defer checkConfig(t, &got, test.expected)
			got = Get()
		})
	}
}

func setupCleanup(t *testing.T) {
	originalReader := reader
	originalGetFileConfigLocation := getConfigLocation
	t.Cleanup(func() {
		reader = originalReader
		getConfigLocation = originalGetFileConfigLocation
	})
}

func generalFileConfigSetup() {
	getConfigLocation = func() (location string) { return "test" }
}

func setupFileConfigCliFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	args := []string{
		"-k",
		"",
	}
	os.Args = append([]string{os.Args[0]}, args...)
}

func setupFileConfigReader(data []byte, err error) {
	reader = func(filename string) ([]byte, error) {
		return data, err
	}
}

func getFileConfigTestCases() []FileConfigTestCase { //nolint
	return []FileConfigTestCase{
		{
			name: "correct secretkey (string from string), files not set",
			data: []byte(getConfig1()),
			err:  nil,
			expected: Config{
				Check:     "",
				SecretKey: "test1",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name: "correct secretkey (string from int), files not set",
			data: []byte(getConfig2()),
			err:  nil,
			expected: Config{
				Check:     "",
				SecretKey: "123456789",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name:     "secretkey key is set, but value is empty string (panics because cli flag is also empty), files not set",
			data:     []byte(getConfig3()),
			err:      nil,
			expected: Config{},
		},
		{
			name: "correct secretkey, files set to empty array",
			data: []byte(getConfig4()),
			err:  nil,
			expected: Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			name: "correct secretkey, files set to array with two files",
			data: []byte(getConfig5()),
			err:  nil,
			expected: Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"testfile1", "testfile2"},
			},
		},
		{
			name:     "correct secretkey, files key is set to a string (panics)",
			data:     []byte(getConfig6()),
			err:      nil,
			expected: Config{},
		},
	}
}
