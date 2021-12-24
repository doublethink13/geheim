package main

import (
	"fmt"
	"treuzedev/geheim/config"
	"treuzedev/geheim/decrypt"
	"treuzedev/geheim/encrypt"
	"treuzedev/geheim/shared"
)

func main() {
	config := config.Get()
	fmt.Println(config)
	for _, filePath := range config.Files {
		if !shared.CheckIfEncrypted(filePath) {
			encrypt.EncryptFile(filePath, config.SecretKey)
		} else {
			decrypt.Decrypt(filePath, config.SecretKey)
		}
	}
}
