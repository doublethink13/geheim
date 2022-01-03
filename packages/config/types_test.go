package config

import "testing"

type CompareConfigsCases struct {
	A        Config
	B        Config
	Expected bool
}

func TestCompareConfigs(t *testing.T) {
	tests := []CompareConfigsCases{
		{Config{"imsecret", []string{".geheim/config.yaml"}, true, true}, Config{"imsecret", []string{".geheim/config.yaml"}, true, true}, true},
		{Config{"imsecret", []string{".geheim/config.yaml"}, true, true}, Config{"imnot", []string{".geheim/config.yaml"}, true, true}, false},
		{Config{"imsecret", []string{".geheim/config.yaml"}, true, true}, Config{"imsecret", []string{"~/geheim/config.yaml"}, true, true}, false},
		{Config{"", []string{".geheim/config.yaml"}, true, true}, Config{"imsecret", []string{".geheim/config.yaml"}, true, true}, true},
		{Config{"imsecret", []string{}, true, true}, Config{"imsecret", []string{".geheim/config.yaml"}, true, true}, true},
		{Config{"imsecret", []string{".geheim/config.yaml", "~/geheim/config.yaml"}, true, true}, Config{"imsecret", []string{".geheim/config.yaml"}, true, true}, false},
	}
	for _, test := range tests {
		r := compareConfigs(test.A, test.B)
		if test.Expected != r {
			t.Errorf("compareConfigs(%s, %s) == %t, expected %t", test.A, test.B, r, test.Expected)
		}
	}
}
