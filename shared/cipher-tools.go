package shared

import (
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
	"time"
)

func GetCipherResources(keyAsString string) (cipher.AEAD, []byte) {
	keyAsBytes := generateEncryptionKey(keyAsString)
	block, err := aes.NewCipher(keyAsBytes)
	CheckError(err)
	gcm, err := cipher.NewGCM(block)
	CheckError(err)
	nonce := make([]byte, gcm.NonceSize())
	CheckError(err)
	return gcm, nonce
}

func RandomizeNonce(nonce []byte) []byte {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < len(nonce); i++ {
		nonce[i] = (nonce[i] + 1) * byte(r.Intn(255))
	}
	return nonce
}

func generateEncryptionKey(keyAsString string) []byte {
	bytesToAppend := GetBytesToAppend()
	keyAsBytes := []byte(keyAsString)
	encryptionKeySize := GetEncryptionKeySize()
	if len(keyAsBytes) > encryptionKeySize {
		keyAsBytes = keyAsBytes[:encryptionKeySize]
	} else {
		for i := 0; len(keyAsBytes) < encryptionKeySize; i++ {
			keyAsBytes = append(keyAsBytes, bytesToAppend[i])
		}
	}
	return keyAsBytes
}
