package testhelpers

//nolint
import (
	"flag"
	"os"
	"testing"

	"treuzedev/geheim/packages/config"
	"treuzedev/geheim/packages/shared"
)

type Flags struct {
	Check     []string
	SecretKey []string
	Encrypt   []string
	Decrypt   []string
}

type FileConfigTestCase struct {
	Name     string
	Data     []byte
	Err      error
	Expected config.Config
}

type FlagsTestCase struct {
	Name     string
	Flags    Flags
	Expected config.Config
}

type CompareConfigsCase struct {
	Name     string
	A        config.Config
	B        config.Config
	Expected bool
}

func CheckConfig(t *testing.T, got *config.Config, expected config.Config) {
	isExpectedEmptyConfig := config.CompareConfigs(
		config.Config{
			Check:     "",
			SecretKey: "",
			Files:     []string{},
			Encrypt:   false,
			Decrypt:   false,
		},
		expected,
	)

	isGotEqualToExpected := config.CompareConfigs(*got, expected)

	if r := recover(); r != nil && !isExpectedEmptyConfig {
		t.Logf("Expected: %v, got: panic", expected)
		t.Logf("%v", r)
		t.Fail()
	} else if !isExpectedEmptyConfig && !isGotEqualToExpected {
		t.Logf("Expected: %v, got: %v", expected, *got)
		t.Fail()
	}
}

func SetupCheckFlagEnv(test FlagsTestCase) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	args := make([]string, 0)

	if !shared.CompareStringSlices(test.Flags.Check, []string{}) {
		args = append(args, test.Flags.Check...)
	}

	if !shared.CompareStringSlices(test.Flags.SecretKey, []string{}) {
		args = append(args, test.Flags.SecretKey...)
	}

	os.Args = append([]string{os.Args[0]}, args...)
}

func FileConfigCliFlagsSetup() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	args := []string{
		"-k",
		"",
	}
	os.Args = append([]string{os.Args[0]}, args...)
}
