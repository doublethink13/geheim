package main_test

//nolint
import (
	"fmt"
	"testing"

	testmain "treuzedev/geheim"
	"treuzedev/geheim/packages/testhelpers"
)

//nolint:paralleltest
func TestGeheimEncryptionDecryption(t *testing.T) {
	tests := testhelpers.GetEncryptionDecryptionTestCases()
	for i, test := range tests {
		filepath := fmt.Sprintf("testfile.test.%v", i)
		testhelpers.GenerateTestFiles(test.Testfile, filepath)

		test.Config.Files = []string{filepath}

		t.Run(test.Name, func(t *testing.T) {
			testmain.Geheim(test.Config)

			if !testhelpers.CheckTestfileResult(filepath, test.Expected) {
				t.Fail()
			}
		})

		testhelpers.RemoveTestFile(filepath)
	}
}

//nolint:paralleltest
func TestGeheimCheck(t *testing.T) {
	tests := testhelpers.GetCheckTestCases()
	for i, test := range tests {
		filepath := fmt.Sprintf("testfile.test.%v", i)
		testhelpers.GenerateTestFiles(test.Testfile, filepath)

		test.Config.Files = []string{filepath}

		t.Run(test.Name, func(t *testing.T) {
			defer testhelpers.CheckPanic(t, test.Expected)
			testmain.Geheim(test.Config)
		})

		testhelpers.RemoveTestFile(filepath)
	}
}
