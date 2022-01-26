package testhelpers

import "treuzedev/geheim/packages/config"

type EncryptionDecryptionTestCase struct {
	Name     string
	Config   config.Config
	Testfile string
	Expected string
}

func GetEncryptionDecryptionTestCases() (testCases []EncryptionDecryptionTestCase) {
	return []EncryptionDecryptionTestCase{
		{
			Name: "encrypt decrypted file",
			Config: config.Config{
				Check:     "",
				SecretKey: "imsosecret",
				Encrypt:   true,
				Decrypt:   false,
			},
			Testfile: GetTestfile1D(),
			Expected: GetTestfile1E(),
		},
		{
			Name: "decrypt encrypted file",
			Config: config.Config{
				Check:     "",
				SecretKey: "imsosecret",
				Encrypt:   false,
				Decrypt:   true,
			},
			Testfile: GetTestfile2E(),
			Expected: GetTestfile2D(),
		},
		{
			Name: "encrypt encrypted file",
			Config: config.Config{
				Check:     "",
				SecretKey: "imsosecret",
				Encrypt:   true,
				Decrypt:   false,
			},
			Testfile: GetTestfile1E(),
			Expected: GetTestfile1E(),
		},
		{
			Name: "decrypt decrypted file",
			Config: config.Config{
				Check:     "",
				SecretKey: "imsosecret",
				Encrypt:   false,
				Decrypt:   true,
			},
			Testfile: GetTestfile2D(),
			Expected: GetTestfile2D(),
		},
	}
}

type CheckTestCase struct {
	Name     string
	Config   config.Config
	Testfile string
	Expected string
}

func GetCheckTestCases() (testCases []CheckTestCase) {
	return []CheckTestCase{

		{
			Name: "confirm that encrypted file is encrypted",
			Config: config.Config{
				Check:     "encrypted",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
			},
			Testfile: GetTestfile0E(),
			Expected: "",
		},
		{
			Name: "confirm that decrypted file is decrypted",
			Config: config.Config{
				Check:     "decrypted",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
			},
			Testfile: GetTestfile0D(),
			Expected: "",
		},
		{
			Name: "panics when encrypted file is decrypted",
			Config: config.Config{
				Check:     "e",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
			},
			Testfile: GetTestfile0D(),
			Expected: "panic",
		},
		{
			Name: "panics when decrypted file is encrypted",
			Config: config.Config{
				Check:     "d",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
			},
			Testfile: GetTestfile0E(),
			Expected: "panic",
		},
		{
			Name: "panics when check flag has wrong value",
			Config: config.Config{
				Check:     "wrongvalue",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
			},
			Testfile: GetTestfile0E(),
			Expected: "panic",
		},
	}
}

type CheckStateTestCase struct {
	Name     string
	State    string
	Expected string
}

func GetCheckStateTestCases() (testCases []CheckStateTestCase) {
	return []CheckStateTestCase{
		{
			Name:     "run default case",
			State:    "imweird",
			Expected: "panic",
		},
	}
}
