package testhelpers

import "treuzedev/geheim/packages/config"

var TestCases = []struct {
	Name     string
	Config   config.Config
	Expected string
}{
	{Name: "encrypt decrypted file", Config: config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   true,
		Decrypt:   false,
		Files:     []string{"testfile1.test"},
	}, Expected: ""},
	{Name: "decrypt encrypted file", Config: config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   false,
		Decrypt:   true,
		Files:     []string{"testfile2.test"},
	}, Expected: ""},
	{Name: "encrypt encrypted file", Config: config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   true,
		Decrypt:   false,
		Files:     []string{"testfile1.test"},
	}, Expected: ""},
	{Name: "decrypt decrypted file", Config: config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   false,
		Decrypt:   true,
		Files:     []string{"testfile2.test"},
	}, Expected: ""},
	{Name: "check encrypted file that is encrypted", Config: config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   true,
		Decrypt:   false,
		Files:     []string{"testfile3.test"},
	}, Expected: ""},
	{Name: "check decrypted file that is decrypted", Config: config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   true,
		Decrypt:   false,
		Files:     []string{"testfile4.test"},
	}, Expected: ""},
	{Name: "check encrypted file that is decrypted", Config: config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   true,
		Decrypt:   false,
		Files:     []string{"testfile5.test"},
	}, Expected: ""},
	{Name: "check decrypted file that is encrypted", Config: config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   true,
		Decrypt:   false,
		Files:     []string{"testfile6.test"},
	}, Expected: ""},
}
