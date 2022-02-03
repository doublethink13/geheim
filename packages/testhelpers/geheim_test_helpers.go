package testhelpers

//nolint
import (
	"bytes"
	"io/fs"
	"os"
	"testing"

	"treuzedev/geheim/packages/shared"
)

func GenerateTestFiles(testfile, filepath string) {
	data := []byte(testfile)
	filePermissions := 0o644
	err := os.WriteFile(filepath, data, fs.FileMode(filePermissions))
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
	t.Helper()

	if r := recover(); r != nil {
		if expected != "panic" {
			t.Fail()
		}
	} else {
		if expected != "" {
			t.Fail()
		}
	}
}
