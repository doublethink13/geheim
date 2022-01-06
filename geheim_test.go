package main

import (
	"testing"
	"treuzedev/geheim/testhelpers"
)

func TestGeheim(t *testing.T) {
	tests := testhelpers.TestCases
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			geheim(test.Config)
		})
	}
}
