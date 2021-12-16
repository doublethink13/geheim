package main

import (
	"fmt"
	"treuzedev/geheim/config"
	"treuzedev/geheim/encrypt"
)

func main() {
	config := config.Get()
	fmt.Println(config)
	for _, filePath := range config.Files {
		encrypt.EncryptFile(filePath, config.SecretKey)
	}
}
