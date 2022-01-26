package geheim_test

//nolint
import (
	"fmt"
	"testing"

	"treuzedev/geheim"
	"treuzedev/geheim/packages/testhelpers"
)

func TestGeheimEncryptionDecryption(t *testing.T) {
	t.Parallel()

	tests := testhelpers.GetEncryptionDecryptionTestCases()
	for i, test := range tests {
		test := test

		filepath := fmt.Sprintf("testfile.test.%v", i)
		testhelpers.GenerateTestFiles(test.Testfile, filepath)

		test.Config.Files = []string{filepath}

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			geheim.Geheim(test.Config)

			if !testhelpers.CheckTestfileResult(filepath, test.Expected) {
				t.Fail()
			}
		})

		testhelpers.RemoveTestFile(filepath)
	}
}

func TestGeheimCheck(t *testing.T) {
	t.Parallel()

	tests := testhelpers.GetCheckTestCases()
	for i, test := range tests {
		test := test

		filepath := fmt.Sprintf("testfile.test.%v", i)
		testhelpers.GenerateTestFiles(test.Testfile, filepath)

		test.Config.Files = []string{filepath}

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			defer testhelpers.CheckPanic(t, test.Expected)

			geheim.Geheim(test.Config)
		})

		testhelpers.RemoveTestFile(filepath)
	}
}
