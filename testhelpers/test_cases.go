package testhelpers

import "treuzedev/geheim/packages/config"

var TestCasesEncryptionDecryption = []struct {
	Name     string
	Config   config.Config
	Testfile string
	Expected string
}{
	{
		Name: "encrypt decrypted file",
		Config: config.Config{
			Check:     "",
			SecretKey: "imsosecret",
			Encrypt:   true,
			Decrypt:   false,
		},
		Testfile: Testfile1D,
		Expected: Testfile1E,
	},
	{
		Name: "decrypt encrypted file",
		Config: config.Config{
			Check:     "",
			SecretKey: "imsosecret",
			Encrypt:   false,
			Decrypt:   true,
		},
		Testfile: Testfile2E,
		Expected: Testfile2D,
	},
	{
		Name: "encrypt encrypted file",
		Config: config.Config{
			Check:     "",
			SecretKey: "imsosecret",
			Encrypt:   true,
			Decrypt:   false,
		},
		Testfile: Testfile1E,
		Expected: Testfile1E,
	},
	{
		Name: "decrypt decrypted file",
		Config: config.Config{
			Check:     "",
			SecretKey: "imsosecret",
			Encrypt:   false,
			Decrypt:   true,
		},
		Testfile: Testfile2D,
		Expected: Testfile2D,
	},
}

var TestCasesCheck = []struct {
	Name     string
	Config   config.Config
	Testfile string
	Expected string
}{

	{
		Name: "confirm that encrypted file is encrypted",
		Config: config.Config{
			Check:     "encrypted",
			SecretKey: "imsosecret",
			Encrypt:   true,
			Decrypt:   false,
		},
		Testfile: Testfile0E,
		Expected: "",
	},
	{
		Name: "confirm that decrypted file is decrypted",
		Config: config.Config{
			Check:     "decrypted",
			SecretKey: "imsosecret",
			Encrypt:   true,
			Decrypt:   false,
		},
		Testfile: Testfile0D,
		Expected: ""},
	{
		Name: "panics when encrypted file is decrypted",
		Config: config.Config{
			Check:     "e",
			SecretKey: "imsosecret",
			Encrypt:   true,
			Decrypt:   false,
		},
		Testfile: Testfile0D,
		Expected: "panic",
	},
	{
		Name: "panics when decrypted file is encrypted",
		Config: config.Config{
			Check:     "d",
			SecretKey: "imsosecret",
			Encrypt:   true,
			Decrypt:   false,
		},
		Testfile: Testfile0E,
		Expected: "panic",
	},
	{
		Name: "panics when check flag has wrong value",
		Config: config.Config{
			Check:     "wrongvalue",
			SecretKey: "imsosecret",
			Encrypt:   true,
			Decrypt:   false,
		},
		Testfile: Testfile0E,
		Expected: "panic",
	},
}

var TestCasesCheckState = []struct {
	Name     string
	State    string
	Expected string
}{
	{
		Name:     "run default case",
		State:    "imweird",
		Expected: "panic",
	},
}
