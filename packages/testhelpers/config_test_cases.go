package testhelpers

import "treuzedev/geheim/packages/config"

//nolint
func GetCheckFlagTestCases() []FlagsTestCase {
	return []FlagsTestCase{
		{
			Name: "short flag + short option (encrypted file)",
			Flags: Flags{
				Check:     []string{"-c", "e"},
				SecretKey: []string{},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "e",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + long option (encrypted file)",
			Flags: Flags{
				Check:     []string{"-c", "encrypted"},
				SecretKey: []string{},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "encrypted",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + short option (encrypted file)",
			Flags: Flags{
				Check:     []string{"--check", "e"},
				SecretKey: []string{},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "e",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + long option (encrypted file)",
			Flags: Flags{
				Check:     []string{"--check", "encrypted"},
				SecretKey: []string{},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "encrypted",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + short option (decrypted file)",
			Flags: Flags{
				Check:     []string{"-c", "d"},
				SecretKey: []string{},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "d",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + long option (decrypted file)",
			Flags: Flags{
				Check:     []string{"-c", "decrypted"},
				SecretKey: []string{},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "decrypted",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + short option (decrypted file)",
			Flags: Flags{
				Check:     []string{"--check", "d"},
				SecretKey: []string{},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "d",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + long option (decrypted file)",
			Flags: Flags{
				Check:     []string{"--check", "decrypted"},
				SecretKey: []string{},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "decrypted",
				SecretKey: "",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "not set",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
	}
}

//nolint
func GetDecryptFlagTestCases() (testCases []FlagsTestCase) {
	return []FlagsTestCase{
		{
			Name: "short flag + no option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{"-d", ""},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   false,
				Decrypt:   true,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + bool option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{"-d", "true"},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   false,
				Decrypt:   true,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + not bool option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{"-d", "imnotabool"},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   false,
				Decrypt:   true,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + no option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{"--decrypt", ""},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   false,
				Decrypt:   true,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + bool option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{"--decrypt", ""},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   false,
				Decrypt:   true,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + not bool option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{"--decrypt", "imnotabool"},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   false,
				Decrypt:   true,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + false",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{"-d=false", ""},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + false",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{"--decrypt=false", ""},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "not set",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{},
				Decrypt:   []string{"", ""},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
	}
}

//nolint
func GetEncryptFlagTestCases() (testCases []FlagsTestCase) {
	return []FlagsTestCase{
		{
			Name: "short flag + no option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{"-e", ""},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + bool option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{"-e", "true"},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + not bool option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{"-e", "imnotabool"},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + no option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{"--encrypt", ""},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + bool option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{"--encrypt", "true"},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + not bool option",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{"--encrypt", "imnotabool"},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "short flag + false",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{"-e=false", ""},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag + false",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{"--encrypt=false", ""},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "not set",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test"},
				Encrypt:   []string{"", ""},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
	}
}

//nolint
func GetFileConfigTestCases() []FileConfigTestCase {
	return []FileConfigTestCase{
		{
			Name: "correct SecretKey (string from string), files not set",
			Data: []byte(GetConfig1()),
			Err:  nil,
			Expected: config.Config{
				Check:     "",
				SecretKey: "test1",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "correct SecretKey (string from int), files not set",
			Data: []byte(GetConfig2()),
			Err:  nil,
			Expected: config.Config{
				Check:     "",
				SecretKey: "123456789",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name:     "SecretKey key is set, but value is empty string (panics because cli flag is also empty), files not set",
			Data:     []byte(GetConfig3()),
			Err:      nil,
			Expected: config.Config{},
		},
		{
			Name: "correct SecretKey, files set to empty array",
			Data: []byte(GetConfig4()),
			Err:  nil,
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "correct SecretKey, files set to array with two files",
			Data: []byte(GetConfig5()),
			Err:  nil,
			Expected: config.Config{
				Check:     "",
				SecretKey: "test",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"testfile1", "testfile2"},
			},
		},
		{
			Name:     "correct SecretKey, files key is set to a string (panics)",
			Data:     []byte(GetConfig6()),
			Err:      nil,
			Expected: config.Config{},
		},
	}
}

func GetSecretKeyFlagTestCases() (testCases []FlagsTestCase) {
	return []FlagsTestCase{
		{
			Name: "short flag",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"-k", "test1"},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test1",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "long flag",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"--secretkey", "test2"},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{
				Check:     "",
				SecretKey: "test2",
				Encrypt:   true,
				Decrypt:   false,
				Files:     []string{"secrets.geheim.yaml"},
			},
		},
		{
			Name: "not set",
			Flags: Flags{
				Check:     []string{},
				SecretKey: []string{"", ""},
				Encrypt:   []string{},
				Decrypt:   []string{},
			},
			Expected: config.Config{},
		},
	}
}

//nolint
func GetCompareConfigsTestCases() (testCases []CompareConfigsCase) {
	return []CompareConfigsCase{
		{
			Name: "equal configs",
			A: config.Config{
				Check:     "encrypted",
				SecretKey: "imsosecret",
				Files:     []string{".geheim/config.yaml"},
				Encrypt:   true,
				Decrypt:   true,
			},
			B: config.Config{
				Check:     "encrypted",
				SecretKey: "imsosecret",
				Files:     []string{".geheim/config.yaml"},
				Encrypt:   true,
				Decrypt:   true,
			},
			Expected: true,
		},
		{
			Name: "secretkey not equal",
			A: config.Config{
				Check:     "encrypted",
				SecretKey: "imsosecret",
				Files:     []string{".geheim/config.yaml"},
				Encrypt:   true,
				Decrypt:   true,
			},
			B: config.Config{
				Check:     "encrypted",
				SecretKey: "imnot",
				Files:     []string{".geheim/config.yaml"},
				Encrypt:   true,
				Decrypt:   true,
			},
			Expected: false,
		},
		{
			Name: "files not equal",
			A: config.Config{
				Check:     "encrypted",
				SecretKey: "imsosecret",
				Files:     []string{".geheim/config.yaml"},
				Encrypt:   true,
				Decrypt:   true,
			},
			B: config.Config{
				Check:     "encrypted",
				SecretKey: "imsosecret",
				Files:     []string{"~/geheim/config.yaml"},
				Encrypt:   true,
				Decrypt:   true,
			},
			Expected: false,
		},
	}
}
