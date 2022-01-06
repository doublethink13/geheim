package main

import (
	"fmt"
	"testing"
	"treuzedev/geheim/packages/config"
	"treuzedev/geheim/testhelpers"
)

func TestGeheimEncryptionDecryption(t *testing.T) {
	tests := testhelpers.TestCasesEncryptionDecryption
	for i, test := range tests {
		filepath := fmt.Sprintf("testfile.test.%v", i)
		testhelpers.GenerateTestFiles(test.Testfile, filepath)
		test.Config.Files = []string{filepath}
		t.Run(test.Name, func(t *testing.T) {
			geheim(test.Config)
			if !testhelpers.CheckTestfileResult(filepath, test.Expected) {
				t.Fail()
			}
		})
		testhelpers.RemoveTestFile(filepath)
	}
}
func TestGeheimCheck(t *testing.T) {
	tests := testhelpers.TestCasesCheck
	for i, test := range tests {
		filepath := fmt.Sprintf("testfile.test.%v", i)
		testhelpers.GenerateTestFiles(test.Testfile, filepath)
		test.Config.Files = []string{filepath}
		t.Run(test.Name, func(t *testing.T) {
			defer testhelpers.CheckPanic(t, test.Expected)
			geheim(test.Config)
		})
		testhelpers.RemoveTestFile(filepath)
	}
}
func TestCheckState(t *testing.T) {
	tests := testhelpers.TestCasesCheckState
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			defer testhelpers.CheckPanic(t, test.Expected)
			config := config.Config{}
			config.Files = []string{""}
			checkState(test.State, config)
		})
	}
}
