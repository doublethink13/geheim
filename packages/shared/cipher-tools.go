package shared

import (
	"crypto/aes"
	"crypto/cipher"
)

func GetCipher(keyAsString string) cipher.Block {
	keyAsBytes := generateEncryptionKey(keyAsString)
	cipher, err := aes.NewCipher(keyAsBytes)
	CheckError(err)
	return cipher
}

func generateEncryptionKey(keyAsString string) []byte {
	keyAsBytes := []byte(keyAsString)
	encryptionKeySize := GetEncryptionKeySize()
	if len(keyAsBytes) > encryptionKeySize {
		keyAsBytes = keyAsBytes[:encryptionKeySize]
	} else {
		bytesToAppend := GetBytesToAppend()
		for i := 0; len(keyAsBytes) < encryptionKeySize; i++ {
			keyAsBytes = append(keyAsBytes, bytesToAppend[i])
		}
	}
	return keyAsBytes
}
