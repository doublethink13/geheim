package config

import "testing"

type CompareConfigsCases struct {
	A        Config
	B        Config
	Expected bool
}

func TestCompareConfigs(t *testing.T) {
	tests := []CompareConfigsCases{
		{Config{"imsecret", []string{".geheim/config.yaml"}}, Config{"imsecret", []string{".geheim/config.yaml"}}, true},
		{Config{"imsecret", []string{".geheim/config.yaml"}}, Config{"imnot", []string{".geheim/config.yaml"}}, false},
		{Config{"imsecret", []string{".geheim/config.yaml"}}, Config{"imsecret", []string{"~/geheim/config.yaml"}}, false},
		{Config{"", []string{".geheim/config.yaml"}}, Config{"imsecret", []string{".geheim/config.yaml"}}, false},
		{Config{"imsecret", []string{}}, Config{"imsecret", []string{".geheim/config.yaml"}}, false},
		{Config{"imsecret", []string{".geheim/config.yaml", "~/geheim/config.yaml"}}, Config{"imsecret", []string{".geheim/config.yaml"}}, false},
	}
	for _, test := range tests {
		r := compareConfigs(test.A, test.B)
		if test.Expected != r {
			t.Errorf("compareConfigs(%s, %s) == %t, expected %t", test.A, test.B, r, test.Expected)
		}
	}
}
