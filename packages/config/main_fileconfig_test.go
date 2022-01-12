package config

import (
	"fmt"
	"testing"
)

func init() {
	reader = func(filename string) ([]byte, error) {
		fmt.Println(filename)
		testString := []byte("hello")
		return testString, nil
	}
}

func TestFileConfig(t *testing.T) {
	reader("test")
}
