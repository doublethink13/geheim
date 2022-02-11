package config_test

//nolint
import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"treuzedev/geheim/packages/config"
	"treuzedev/geheim/packages/testhelpers"
)

//nolint:paralleltest
func TestCheckFlag(t *testing.T) {
	testCases := testhelpers.GetCheckFlagTestCases()
	for _, test := range testCases {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			testhelpers.SetupCheckFlagEnv(test)

			testConfig := config.Get("", ioutil.ReadFile)
			if !config.CompareConfigs(testConfig, test.Expected) {
				t.Logf("Expected: %v, got: %v", test.Expected, testConfig)
				t.Fail()
			}
		})
	}
}

//nolint:paralleltest
func TestDecryptFlag(t *testing.T) {
	for _, test := range testhelpers.GetDecryptFlagTestCases() {
		test := test

		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		args := []string{
			test.Flags.SecretKey[0],
			test.Flags.SecretKey[1],
			test.Flags.Decrypt[0],
			test.Flags.Decrypt[1],
		}
		os.Args = append([]string{os.Args[0]}, args...)

		t.Run(test.Name, func(t *testing.T) {
			testConfig := config.Get("", ioutil.ReadFile)
			if !config.CompareConfigs(testConfig, test.Expected) {
				t.Logf("Expected: %v, got: %v", test.Expected, testConfig)
				t.Fail()
			}
		})
	}
}

//nolint:paralleltest
func TestEncryptFlag(t *testing.T) {
	for _, test := range testhelpers.GetEncryptFlagTestCases() {
		test := test

		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		args := []string{
			test.Flags.SecretKey[0],
			test.Flags.SecretKey[1],
			test.Flags.Encrypt[0],
			test.Flags.Encrypt[1],
		}
		os.Args = append([]string{os.Args[0]}, args...)

		t.Run(test.Name, func(t *testing.T) {
			testConfig := config.Get("", ioutil.ReadFile)
			if !config.CompareConfigs(testConfig, test.Expected) {
				t.Logf("Expected: %v, got: %v", test.Expected, testConfig)
				t.Fail()
			}
		})
	}
}

//nolint:paralleltest
func TestFileConfig(t *testing.T) {
	testCases := testhelpers.GetFileConfigTestCases()
	for _, test := range testCases {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			testhelpers.FileConfigCliFlagsSetup()

			mockReader := func(filename string) ([]byte, error) {
				return test.Data, test.Err
			}

			var got config.Config

			defer testhelpers.CheckConfig(t, &got, test.Expected)

			got = config.Get("mock", mockReader)
		})
	}
}

//nolint:paralleltest
func TestSecretKeyFlag(t *testing.T) {
	for _, test := range testhelpers.GetSecretKeyFlagTestCases() {
		test := test

		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		args := []string{
			test.Flags.SecretKey[0],
			test.Flags.SecretKey[1],
		}
		os.Args = append([]string{os.Args[0]}, args...)

		t.Run(test.Name, func(t *testing.T) {
			var got config.Config

			defer testhelpers.CheckConfig(t, &got, test.Expected)

			got = config.Get("", ioutil.ReadFile)
		})
	}
}

//nolint:paralleltest
func TestCompareConfigs(t *testing.T) {
	tests := testhelpers.GetCompareConfigsTestCases()

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			if r := config.CompareConfigs(test.A, test.B); test.Expected != r {
				t.Logf(
					"compareConfigs(%v, %v) == %t, expected %t",
					test.A,
					test.B,
					r,
					test.Expected,
				)
				t.Fail()
			}
		})
	}
}
