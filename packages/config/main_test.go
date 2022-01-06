package config

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestLocalConfig(t *testing.T) {
	tests := []Config{{"", "imsosecret", []string{
		"testfiles/config.json",
		"testfiles/coverage_testfile.xml",
		"testfiles/helpers.sh",
		"testfiles/id_rsa",
		"testfiles/id_rsa.pub",
		"testfiles/known_hosts",
		"testfiles/secrets_test.geheim.yaml",
		"testfiles/simple.txt",
		"testfiles/supervisor_env",
	}, true, false}}
	for _, test := range tests {
		config := Get()
		if !compareConfigs(config, test) {
			t.Errorf("Get() == %v, expected %v", config, test)
		}
	}
}
