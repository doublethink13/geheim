package main

import (
	"testing"
	"treuzedev/geheim/packages/config"
	"treuzedev/geheim/testhelpers"
)

func TestGeheim(t *testing.T) {
	tests := testhelpers.TestCases
	for _, test := range tests {
		generateTestFiles(test.Config)
		t.Run(test.Name, func(t *testing.T) {
			geheim(test.Config)
			checkTestfileResult(test.Config)
		})
		removeTestFiles(test.Config)
	}
}

func generateTestFiles(config config.Config) {
	return
}

func checkTestfileResult(config config.Config) {
	return
}

func removeTestFiles(config config.Config) {
	return
}
