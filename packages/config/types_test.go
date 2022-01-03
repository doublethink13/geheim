package config

import "testing"

type CompareConfigsCases struct {
	A        Config
	B        Config
	Expected bool
}

func TestCompareConfigs(t *testing.T) {
	tests := []CompareConfigsCases{
		{Config{"imsosecret", []string{".geheim/config.yaml"}, true, true}, Config{"imsosecret", []string{".geheim/config.yaml"}, true, true}, true},
		{Config{"imsosecret", []string{".geheim/config.yaml"}, true, true}, Config{"imnot", []string{".geheim/config.yaml"}, true, true}, false},
		{Config{"imsosecret", []string{".geheim/config.yaml"}, true, true}, Config{"imsosecret", []string{"~/geheim/config.yaml"}, true, true}, false},
		{Config{"", []string{".geheim/config.yaml"}, true, true}, Config{"imsosecret", []string{".geheim/config.yaml"}, true, true}, false},
		{Config{"imsosecret", []string{}, true, true}, Config{"imsosecret", []string{".geheim/config.yaml"}, true, true}, false},
		{Config{"imsosecret", []string{".geheim/config.yaml", "~/geheim/config.yaml"}, true, true}, Config{"imsosecret", []string{".geheim/config.yaml"}, true, true}, false},
	}
	for _, test := range tests {
		r := compareConfigs(test.A, test.B)
		if test.Expected != r {
			t.Errorf("compareConfigs(%v, %v) == %t, expected %t", test.A, test.B, r, test.Expected)
		}
	}
}
