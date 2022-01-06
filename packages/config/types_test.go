package config

import "testing"

type CompareConfigsCases struct {
	A        Config
	B        Config
	Expected bool
}

func TestCompareConfigs(t *testing.T) {
	tests := []CompareConfigsCases{
		{Config{"encrypted", "imsosecret", []string{".geheim/config.yaml"}, true, true}, Config{"encrypted", "imsosecret", []string{".geheim/config.yaml"}, true, true}, true},
		{Config{"encrypted", "imsosecret", []string{".geheim/config.yaml"}, true, true}, Config{"encrypted", "imnot", []string{".geheim/config.yaml"}, true, true}, false},
		{Config{"encrypted", "imsosecret", []string{".geheim/config.yaml"}, true, true}, Config{"encrypted", "imsosecret", []string{"~/geheim/config.yaml"}, true, true}, false},
		{Config{"encrypted", "", []string{".geheim/config.yaml"}, true, true}, Config{"encrypted", "imsosecret", []string{".geheim/config.yaml"}, true, true}, false},
		{Config{"encrypted", "imsosecret", []string{}, true, true}, Config{"encrypted", "imsosecret", []string{".geheim/config.yaml"}, true, true}, false},
		{Config{"encrypted", "imsosecret", []string{".geheim/config.yaml", "~/geheim/config.yaml"}, true, true}, Config{"encrypted", "imsosecret", []string{".geheim/config.yaml"}, true, true}, false},
	}
	for _, test := range tests {
		r := compareConfigs(test.A, test.B)
		if test.Expected != r {
			t.Errorf("compareConfigs(%v, %v) == %t, expected %t", test.A, test.B, r, test.Expected)
		}
	}
}
