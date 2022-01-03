package config

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestLocalConfig(t *testing.T) {
	tests := []Config{{"imsosecret", []string{"secrets.geheim.yaml"}, true, true}}
	for _, test := range tests {
		config := Get()
		if !compareConfigs(config, test) {
			t.Errorf("Get() == %s, expected %s", config, test)
		}
	}
}
