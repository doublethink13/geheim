package testhelpers

import (
	"bytes"
	"os"
	"testing"
	"treuzedev/geheim/packages/shared"
)

func GenerateTestFiles(testfile, filepath string) {
	data := []byte(testfile)
	err := os.WriteFile(filepath, data, 0644)
	shared.CheckError(err, nil)
}

func CheckTestfileResult(filepath, expected string) (areEqual bool) {
	r, err := os.ReadFile(filepath)
	shared.CheckError(err, nil)
	expectedBytes := []byte(expected)
	return bytes.Equal(r, expectedBytes)
}

func RemoveTestFile(filepath string) {
	err := os.Remove(filepath)
	shared.CheckError(err, nil)
}

func CheckPanic(t *testing.T, expected string) {
	r := recover()
	if r != nil {
		if expected != "panic" {
			t.Fail()
		}
	} else {
		if expected != "" {
			t.Fail()
		}
	}
}
