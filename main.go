package main

import (
	"treuzedev/geheim/packages/config"
	"treuzedev/geheim/packages/decrypt"
	"treuzedev/geheim/packages/encrypt"
	"treuzedev/geheim/packages/shared"
)

func main() {
	config := config.Get()
	for _, filePath := range config.Files {
		if !shared.CheckIfEncrypted(filePath) {
			encrypt.EncryptFile(filePath, config.SecretKey)
		} else {
			decrypt.Decrypt(filePath, config.SecretKey)
		}
	}
}
