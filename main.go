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
		if config.Encrypt {
			if !shared.CheckIfEncrypted(filePath) {
				encrypt.EncryptFile(filePath, config.SecretKey)
			}
		}
		if config.Decrypt {
			if shared.CheckIfEncrypted(filePath) {
				decrypt.Decrypt(filePath, config.SecretKey)
			}
		}
	}
}
